package controllers

import (
	"gin/models"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func Login(ginServer *gin.Engine) {

	ginServer.POST("/toLogin", func(context *gin.Context) {

		//1.获取登录用户信息
		username := context.PostForm("username")
		passeord := context.PostForm("password")

		//2.创建用户结构体
		user := models.User{
			Username: username,
			Password: passeord,
			ID:       "100",
		}

		//3.把结构体纳入到session中
		session := sessions.Default(context)
		session.Set("user", user) //存储数据
		session.Save()            //生成user映射结构体

		//4.返回信息
		context.JSON(200, gin.H{"user": user, "status": 200})
	})
}
