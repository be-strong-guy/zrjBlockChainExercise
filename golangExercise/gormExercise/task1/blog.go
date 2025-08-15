package task1

import (
	"fmt"

	"gorm.io/gorm"
)

/*
*
假设你要开发一个博客系统，有以下几个实体：
User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，
其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	PostCount int
	Posts     []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	Title         string
	Content       string
	CommentStatus string    `gorm:"type:varchar(100)"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
}

// 钩子：在文章创建后更新用户的 PostCount
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 增加用户文章数
	err = tx.Model(&User{}).
		Debug().
		Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
	return
}

type Comment struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	PostID  uint
	Content string
}

// 钩子：在评论删除时检查文章的评论数量，更新post表的comment_status
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if err := tx.Model(&Comment{}).
		Debug().
		Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

func CreateBlog(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
	// 初始化几条数据
	user := User{
		Name: "张三丰",
		Posts: []Post{
			Post{
				Title:   "九阳神功",
				Content: "九阳神功免费，但是需要配合九阴真经一起练哦！不然没有效果呢~",
				Comments: []Comment{
					Comment{
						Content: "坑人的商家，九阳神功免费，它九阴真经卖10个亿！！",
					},
				},
			},
		},
	}
	db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&user)
}

/*
*
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func SelectPostAndComments(db *gorm.DB) {
	// 1、查询某个用户发布的所有文章及其对应的评论信息
	var user User
	db.Debug().Preload("Posts.Comments").Where("id = ? ", 3).Find(&user)
	for _, p := range user.Posts {
		fmt.Printf("%s发布的文章标题：%s \n", user.Name, p.Title)
		fmt.Printf("%s张三丰发布的文章内容：%s\n", user.Name, p.Content)
		for _, comment := range p.Comments {
			fmt.Println("文章评价：", comment.Content)
		}
	}

	// 查询评论数量最多的文章
	type PostWithCount struct {
		ID           uint
		Title        string
		CommentCount int
	}

	var result PostWithCount

	err := db.Debug().Table("posts").
		Select("posts.id, posts.title, COUNT(comments.id) AS comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("评论最多的文章是:", result.Title)

}

/*
*为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
/**
这文章数量字段，上面没有提到，那就加一个，应该加在user里面，这个人有多少文章嘛
同样的评论状态字段也没有，加在文章post表里面
*/
func HookBlog(db *gorm.DB) {
	post := Post{
		Title:         "九阴真经",
		Content:       "九阴真经10个亿哦，你真买得起吗？",
		UserID:        3,
		CommentStatus: "有评论",
	}
	// 会触发 AfterCreate
	db.Create(&post)

	comment := Comment{Content: "坑爹！", PostID: post.ID}
	db.Create(&comment)
	// 会触发 AfterDelete,这里没有真删除，是软删除，
	// 可以看到数据还在，真删除的话db.Unscoped().Delete(&comment)
	db.Delete(&comment)

}
