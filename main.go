package main

import (

	"github.com/gin-gonic/gin"
	"gt-gin-vue/middleware/jwt"
	"log"

	"gt-gin-vue/api"
	"gt-gin-vue/model"
)

func main() {

	db, err := model.OpenMysql()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	r := gin.Default()

	a := r.Group("/auth")
	{
		a.POST("/login", api.Login)
		a.POST("/register", api.RegisterUser)
	}
		
	u := r.Group("/user")
	u.Use(jwt.JWTAuth())
	{
		u.GET("/test", api.TestToken)
		u.GET("/resetpwd", api.ResetPassword)
		
		u.POST("/delete", api.DeleteUser)
		//u.GET("/info", api.InfoUsers)
		u.POST("/update", api.UpdateUser)
	}
	

	r.Run(":8080")
}
