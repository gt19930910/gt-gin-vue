package model

import (

	// "fmt"
	// "log"
)

const (
	dbName     = "myBlog.db"
	userBucket = "user"
)

// User 用户类
type User struct {
	Id         string `json:"userId"`
	Name       string `json:"userName"`
	Gender     string `json:"gender"`
	Phone      string `json:"userMobile"`
	Pwd        string `json:"pwd"`
	Permission string `json:"permission"`
}

// LoginReq 登录请求参数类
type LoginReq struct {
	Phone string `json:"mobile"`
	Pwd   string `json:"pwd"`
}

// Register 插入用户，先检查是否存在用户，如果没有则存入
func Register(phone string, pwd string) error {

	return nil
}

// CheckUser 检查用户是否存在
func CheckUser(phone string) bool {

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
