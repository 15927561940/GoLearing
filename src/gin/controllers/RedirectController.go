package controllers

import "github.com/gin-gonic/gin"

func Redirect(ginServer *gin.Engine) {

	//重定向
	ginServer.GET("/testredirect", func(context *gin.Context) {
		context.Redirect(301, "https://www.kuangstudy.com/course/cplay/1498123423423423466")
	})

	//站内重定向
	ginServer.GET("/testinweb", func(context *gin.Context) {
		context.Redirect(301, "/hello")
	})
}
