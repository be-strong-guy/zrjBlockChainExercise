package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zrjBlockChainExercise/golangExercise/gormExercise/mysqlTest"
)

func main() {
	dsn := "root:Pass123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	mysqlTest.Run(db)
}
