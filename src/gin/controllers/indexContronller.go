package controllers

import (
	"gin/filters"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ginServer *gin.Engine) {
	ginServer.GET("/hello", filters.MyHandler(), func(context *gin.Context) { //单业务指定拦截
		context.JSON(http.StatusOK, gin.H{"message": "hello gin"})
	})

	ginServer.GET("/index", filters.MyHandler(), func(context *gin.Context) {
		session := sessions.Default(context)
		user := session.Get("user") //取出已经存在的session

		//查看是否取到
		context.HTML(200, "index.html", gin.H{
			"message": "gin测试静态网页", "name": "波波", "user": user})
	})

}
