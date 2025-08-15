package task1

import (
	"fmt"

	"gorm.io/gorm"
)

/*
*
假设有一个名为 students 的表，包含字段 id （主键，自增）、
name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、
grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
type Student struct {
	gorm.Model
	ID    uint `gorm:"primarykey"`
	Name  string
	Age   int
	Grade string
}

func Run1(db *gorm.DB) {
	student := Student{
		Name:  "张三",
		Age:   10,
		Grade: "三年级",
	}
	err := db.AutoMigrate(&student)
	if err != nil {
		panic(err)
	}
	//插入
	db.Create(&student)
	results := []Student{}
	//查询
	db.Where("age > 18").Find(&results)
	fmt.Println("所有大于18岁学生是：", results)
	//更新
	db.Model(&Student{}).Where("name = ? ", "张三").Update("grade", "四年级")
	db.Find(&student)
	fmt.Println("更新后的结果是：", student)
	//删除
	db.Unscoped().Where("age < ?", 15).Delete(&Student{})
}
