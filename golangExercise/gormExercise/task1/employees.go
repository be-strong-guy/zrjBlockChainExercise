package task1

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

/*
*假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，
包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，
并将结果映射到一个 Employee 结构体中。
*/
type Employee struct {
	gorm.Model
	ID         uint `gorm:"primarykey;auto_increment"`
	Name       string
	Department string
	Salary     int
}

func Run2(db *gorm.DB) {
	sqlDB, err := db.DB()
	xdb := sqlx.NewDb(sqlDB, "mysql")
	employee := Employee{
		Name:       "张三丰",
		Department: "技术部",
		Salary:     10000,
	}
	db.AutoMigrate(employee)
	db.Create(&employee)
	// 查询
	var employees []Employee
	query := `SELECT id, name, department, salary FROM employees WHERE department = ?`
	err = xdb.Select(&employees, query, "技术部")
	if err != nil {
		panic(err)
	}
	fmt.Println("技术部员工有:", employees)
	var topEmployee Employee
	queryMax := `SELECT id, name, department, salary 
                 FROM employees 
                 ORDER BY salary DESC 
                 LIMIT 1`
	err = xdb.Get(&topEmployee, queryMax)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("工资最高的员工是：", topEmployee)
}
