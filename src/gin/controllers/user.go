package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(ginServer *gin.Engine) {
	ginServer.GET("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "user   xxxxxxxxxxxx gin"})
	})
	//3.启动，要给个端口
	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "post   xxxxxxxxxxxx gin"})
	})
	ginServer.DELETE("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "delete   xxxxxxxxxxxx gin"})
	})

	ginServer.PUT("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "put   xxxxxxxxxxxx gin"})
	})

}
