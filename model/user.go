package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"errors"
	"log"
)

var (
	db *gorm.DB
	db_err error
)

// User 用户类 在数据库中表名默认会加上s，所有的字段名都会自动匹配起来，只是大写都变成了小写
// 可以自己在后面改名称用`gorm:"column:username"`如果列名称已经定了就不能插入了
type User struct {
	Id 	     uint `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Username string `gorm:"column:username"` 
	Password string `gorm:"column:password"`
}

// LoginReq 登录请求参数类
type LoginReq struct {
	Name string `json:"name"`
	Pwd   string `json:"pwd"`
}

func OpenMysql() (*gorm.DB, error) {
	db, db_err = gorm.Open("mysql", "gt:123@(127.0.0.1:3306)/dbtest?charset=utf8&parseTime=True&loc=Local")
	if db_err != nil {
		return nil, db_err
	}
	db.LogMode(true)//开启sql debug 模式

	ret := db.HasTable(&User{})
	if ret == false {
		db.CreateTable(&User{})//初始化一个空的表
	}
	return db, nil
}


func Add(user *User) error {
	
	if err := db.Create(user).Error; err != nil {
		log.Fatal("插入失败:", err)
		return err
	}

	return nil
}

// Register 插入用户，先检查是否存在用户，如果没有则存入
func Register(username string, pwd string) error {
	if CheckUser(username) {
		return errors.New("用户已存在！")
	}
	user := &User {
		Username:username,
		Password:pwd,
	}
	return Add(user);
}

// CheckUser 检查用户是否存在
func CheckUser(username string) bool {

	ret := db.Where("username = ?", username).First(&User{}).RecordNotFound()
	if ret == true {
		return false
	}
	return true
}

func Delete(username string, id uint) error {

	if !CheckUser(username) {
		return errors.New("该用户不存在")
	}
	// 删除表中所有的数据
	// u := User {}
	// db.Delete(&u)
	u := User {
		Username:username, 
		Id:id,
	}
	err := db.Delete(&u).Error
	return err
}

func Update(updateUser User) (User, error) {

	//根据前面的user的主键不变，修改成后面user里面的字段
	err := db.Model(&updateUser).Updates(&updateUser).Error
	return updateUser, err
}

// LoginCheck 登录验证
func LoginCheck(loginReq LoginReq) (bool, User, error) {
	user := User {}
	ret := db.Where("username = ?", loginReq.Name).First(&user).RecordNotFound()
	if ret == true {
		return false, user, errors.New("用户名未注册")
	}
	if user.Password != loginReq.Pwd {
		return false, user, errors.New("密码错误")
	}
	return true, user, nil
}

//ResetPwd 重置密码
func ResetPwd(id uint) error {

	user := User {Id:id}
	err := db.Model(&user).Update("password", "123").Error
	return err
}
