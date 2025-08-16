package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	jwtSecret   = []byte("your_secret_key")
	tokenExpire = 24 * time.Hour
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:64;not null" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"` // 存哈希
	Email    string `gorm:"size:255" json:"email"`
	Posts    []Post `json:"-"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"size:255;not null" json:"title"`
	Content  string    `gorm:"type:text;not null" json:"content"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"-"`
	Comments []Comment `json:"comments,omitempty"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	User    User   `json:"-"`
	Post    Post   `json:"-"`
}

/*************** 请求/响应 DTO ***************/
type registerReq struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Password string `json:"password" binding:"required,min=6,max=128"`
	Email    string `json:"email"`
}
type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type createPostReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type updatePostReq struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}
type createCommentReq struct {
	Content string `json:"content" binding:"required"`
}

/*************** JWT 工具/中间件 ***************/
type claims struct {
	UserID   uint   `json:"uid"`
	Username string `json:"uname"`
	jwt.RegisteredClaims
}

// 生成jwt
func genToken(u User) (string, error) {
	c := &claims{
		UserID:   u.ID,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString(jwtSecret)
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}
		tokenStr := strings.TrimPrefix(h, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		cl, ok := token.Claims.(*claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}
		// 存上下文
		c.Set("uid", cl.UserID)
		c.Set("uname", cl.Username)
		c.Next()
	}
}

/*************** 密码工具 ***************/
func hashPassword(pw string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(b), err
}
func checkPassword(hash string, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

/*************** 统一错误响应 ***************/
func abortErr(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
}

/*************** 处理器：认证 ***************/
func register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		abortErr(c, http.StatusBadRequest, err)
		return
	}
	h, err := hashPassword(req.Password)
	if err != nil {
		abortErr(c, http.StatusInternalServerError, err)
		return
	}
	u := User{Username: req.Username, Password: h, Email: req.Email}
	if err := db.Create(&u).Error; err != nil {
		abortErr(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": u.ID, "username": u.Username})
}

func login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		abortErr(c, http.StatusBadRequest, err)
		return
	}
	var u User
	if err := db.Where("username = ?", req.Username).First(&u).Error; err != nil {
		abortErr(c, http.StatusUnauthorized, errors.New("invalid username or password"))
		return
	}
	if err := checkPassword(u.Password, req.Password); err != nil {
		abortErr(c, http.StatusUnauthorized, errors.New("invalid username or password"))
		return
	}
	tok, err := genToken(u)
	if err != nil {
		abortErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tok})
}

/*************** 处理器：文章 ***************/
func createPost(c *gin.Context) {
	var req createPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		abortErr(c, http.StatusBadRequest, err)
		return
	}
	uid := c.GetUint("uid")
	p := Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  uid,
	}
	if err := db.Create(&p).Error; err != nil {
		abortErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, p)
}

func listPosts(c *gin.Context) {
	var posts []Post
	if err := db.Preload("Comments").Order("created_at desc").Find(&posts).Error; err != nil {
		abortErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func getPost(c *gin.Context) {
	var p Post
	if err := db.Preload("Comments").First(&p, c.Param("id")).Error; err != nil {
		abortErr(c, http.StatusNotFound, errors.New("post not found"))
		return
	}
	c.JSON(http.StatusOK, p)
}

func updatePost(c *gin.Context) {
	var p Post
	if err := db.First(&p, c.Param("id")).Error; err != nil {
		abortErr(c, http.StatusNotFound, errors.New("post not found"))
		return
	}
	uid := c.GetUint("uid")
	if p.UserID != uid {
		abortErr(c, http.StatusForbidden, errors.New("not your post"))
		return
	}
	var req updatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		abortErr(c, http.StatusBadRequest, err)
		return
	}
	updates := map[string]interface{}{}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Content != nil {
		updates["content"] = *req.Content
	}
	if len(updates) == 0 {
		c.JSON(http.StatusOK, p)
		return
	}
	if err := db.Model(&p).Updates(updates).Error; err != nil {
		abortErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, p)
}

func deletePost(c *gin.Context) {
	var p Post
	if err := db.First(&p, c.Param("id")).Error; err != nil {
		abortErr(c, http.StatusNotFound, errors.New("post not found"))
		return
	}
	uid := c.GetUint("uid")
	if p.UserID != uid {
		abortErr(c, http.StatusForbidden, errors.New("not your post"))
		return
	}
	if err := db.Delete(&p).Error; err != nil {
		abortErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": p.ID})
}

/*************** 处理器：评论 ***************/
func createComment(c *gin.Context) {
	var p Post
	if err := db.First(&p, c.Param("id")).Error; err != nil {
		abortErr(c, http.StatusNotFound, errors.New("post not found"))
		return
	}
	var req createCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		abortErr(c, http.StatusBadRequest, err)
		return
	}
	uid := c.GetUint("uid")
	cmt := Comment{
		Content: req.Content,
		UserID:  uid,
		PostID:  p.ID,
	}
	if err := db.Create(&cmt).Error; err != nil {
		abortErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, cmt)
}

func listCommentsOfPost(c *gin.Context) {
	var cmts []Comment
	if err := db.Where("post_id = ?", c.Param("id")).Order("created_at asc").Find(&cmts).Error; err != nil {
		abortErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, cmts)
}

/*************** 初始化 ***************/
func connect() *gorm.DB {
	dsn := "root:Pass123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func initDB() {
	var err error

	// 初始化DB
	db = connect()
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		log.Fatalf("migrate error: %v", err)
	}
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}

}

func main() {
	//  JWT 秘钥
	if v := os.Getenv("JWT_SECRET"); v != "" {
		jwtSecret = []byte(v)
	}

	initDB()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// 健康检查
	r.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	// 认证
	r.POST("/v1/register", register)
	r.POST("/v1/login", login)

	// 文章操作（需要登录）
	auth := r.Group("/v1", authRequired())
	{
		// posts
		auth.POST("/posts", createPost)
		auth.GET("/posts", listPosts)
		auth.GET("/posts/:id", getPost)
		auth.PATCH("/posts/:id", updatePost) // 仅作者可改
		auth.DELETE("/posts/:id", deletePost)

		// comments
		auth.POST("/posts/:id/comments", createComment)
		auth.GET("/posts/:id/comments", listCommentsOfPost)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
