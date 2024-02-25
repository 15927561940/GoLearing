package controllers

import (
	"gin/filters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Course(ginServer *gin.Engine) {

	userGroup := ginServer.Group("/course", filters.MyHandler()) //指定拦截
	{
		userGroup.GET("/list", func(context *gin.Context) {
			value, _ := context.Get("userID")

			context.JSON(http.StatusOK, gin.H{"message": "hello gin", "value的值是": value})

		})

		userGroup.GET("/save", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "hello gin"})
		})

		userGroup.GET("/del", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "hello gin"})
		})
	}

}
