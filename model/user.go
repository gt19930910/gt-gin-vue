package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"errors"
	"fmt"
)

// User 用户类
type User struct {
	Id           string
	Username     string
	//Gender     string `json:"gender"`
	//Phone      string `json:"userMobile"`
	Password     string
	//Permission string `json:"permission"`
}

// LoginReq 登录请求参数类
type LoginReq struct {
	Phone string `json:"mobile"`
	Pwd   string `json:"pwd"`
}

var (
	db *gorm.DB
)

func OpenMysql() error {
	db, err := gorm.Open("mysql", "gt:abc123@(127.0.0.1:3306)/gt_gin_vue?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	defer db.Close()
	//db.SingularTable(true)
	db.LogMode(true)//开启sql debug 模式
	return nil
}

func Add() error {
	
	user := &User {
		Username:"username",
		Password:"pwd",
	}
	fmt.Println(*user);
	if err := db.Create(user).Error; err != nil {
		fmt.Println("插入失败", err)
		return err
	}
	return nil
}

// Register 插入用户，先检查是否存在用户，如果没有则存入
func Register(username string, pwd string) error {
	if CheckUser(username) {
		return errors.New("用户已存在！")
	}

	user := User {
		Username:"username",
		Password:"pwd",
	}
	fmt.Println(user);
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("插入失败", err)
		return err
	}
	return nil
}

// CheckUser 检查用户是否存在
func CheckUser(username string) bool {

	return false
}

// LoginCheck 登录验证
func LoginCheck(loginReq LoginReq) (bool, User, error) {
	var a User
	return false, a, nil
}

// EditUserReq 更新用户信息数据类
type EditUserReq struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	UserGender string `json:"gender"`
}

// UpdateUser 更新用户信息
func UpdateUser(editUser EditUserReq) (User, error) {
	
	var a User
	return a, nil
}

//ResetPwd 重置密码
func ResetPwd(mobile string, pwd string) error {

	return nil
}
