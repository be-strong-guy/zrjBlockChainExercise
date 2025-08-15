package task1

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

/*
*
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
并将结果映射到 Book 结构体切片中，确保类型安全。
*/
type Book struct {
	gorm.Model
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Title  string
	Author string
	Price  float64
}

func RunBook(db *gorm.DB) {
	sqlDB, err := db.DB()
	xdb := sqlx.NewDb(sqlDB, "mysql")
	book := Book{
		Title:  "九阳神功",
		Author: "张三丰",
		Price:  1000.09,
	}
	db.AutoMigrate(book)
	db.Create(&book)

	//查询
	var books []Book
	query := `SELECT id,title,author,price FROM books where price > ?`
	err = xdb.Select(&books, query, 50)
	if err != nil {
		panic(err)
	}
	for _, b := range books {
		fmt.Println("查到的结果是:", books)
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n",
			b.ID, b.Title, b.Author, b.Price)
	}

}
