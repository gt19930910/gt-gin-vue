package main

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

type User struct {
	Id           uint `gorm:"primary_key"`
	Username     string
	Password     string
}

func main() {

	db, err := gorm.Open("mysql", "gt:abc123@tcp(127.0.0.1:3306)/gt_gin_vue?charset=utf8")
	if err != nil {
		fmt.Println("err")
		return
	}
	defer db.Close()
	db.CreateTable(&User{})
	user := User {
		Id:1,
		Username:"usme",
		Password:"pwd",
	}
	fmt.Println(user);
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("插入失败", err)
	}
}
