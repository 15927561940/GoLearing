package controllers

import "github.com/gin-gonic/gin"

func Error(ginServer *gin.Engine) {

	//没有路由则捕捉404错误,然后跳转到自定义界面
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(200, "404.html", gin.H{"error": "404页面找不到"})
	})
}
