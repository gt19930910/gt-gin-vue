package api

import (
	myjwt "gt-gin-vue/middleware/jwt"
	"gt-gin-vue/model"
	"log"
	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 注册信息
type RegistInfo struct {
	UserName string `json:"username"`
	Pwd string `json:"pwd"`
}

// Register 注册用户
func RegisterUser(c *gin.Context) {
	var registerInfo RegistInfo
	if c.BindJSON(&registerInfo) == nil {
		err := model.Register(registerInfo.UserName, registerInfo.Pwd)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "注册成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "注册失败:" + err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "解析数据失败！",
		})
	}
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	model.User
}

// Login 登录
func Login(c *gin.Context) {
	var loginReq model.LoginReq
	if c.BindJSON(&loginReq) == nil {
		isPass, user, err := model.LoginCheck(loginReq)
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败," + err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
	}
}

// 生成令牌
func generateToken(c *gin.Context, user model.User) {
	j := &myjwt.JWT{
		[]byte("GaoTaoLearn"),
	}
	claims := myjwt.CustomClaims{
		user.Id,
		user.Username,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "GaoTaoLearn",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

//下面的函数若要有权限问题需要再加中间件

func TestToken(c *gin.Context) {
	//若要加admin权限，需要解析claims
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

func DeleteUser(c *gin.Context) {
	//若要加admin权限，需要解析claims
	claims := c.MustGet("claims").(*myjwt.CustomClaims)

	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "未传递用户名",
		})
		return
	}
	err := model.Delete(name, claims.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "删除用户失败 " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "成功删除用户 " + name,
	})
}

func InfoUsers(c *gin.Context) {
	//若要加admin权限，需要解析claims
	//claims := c.MustGet("claims").(*myjwt.CustomClaims)
}

func UpdateUser(c *gin.Context) {
	//若要加admin权限，需要解析claims
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	name := c.PostForm("name")
	pwd  := c.PostForm("pwd")
	if name == "" && pwd == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "无任何修改参数",
		})
	} 
	user := model.User {
		Id:claims.Id,
		Password: pwd,
		Username: name,
	}
	retUser, err := model.Update(user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "更新用户数据失败 " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "更新用户数据成功 ",
		"data": retUser,
	})
}

func ResetPassword(c *gin.Context) {
	//若要加admin权限，需要解析claims
	claims := c.MustGet("claims").(*myjwt.CustomClaims)

	err := model.ResetPwd(claims.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "重置密码失败 " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "重置密码为123",
	})
}
