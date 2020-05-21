package main

import (

	"github.com/gin-gonic/gin"
	"gt-gin-vue/middleware/jwt"
	"log"

	"gt-gin-vue/api"
	"gt-gin-vue/model"
)

func main() {

	err := model.OpenMysql()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()
	r.POST("/login", api.Login)
	r.POST("/register", api.RegisterUser)

	taR := r.Group("/data")
	taR.Use(jwt.JWTAuth())

	{
		taR.GET("/dataByTime", api.GetDataByTime)
	}
	r.Run(":8080")
}
