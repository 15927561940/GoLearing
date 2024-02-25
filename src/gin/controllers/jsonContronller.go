package controllers

import (
	"gin/models"
	"github.com/gin-gonic/gin"
)

func Json(ginServer *gin.Engine) {
	//1.返回普通数据类型
	ginServer.GET("/json1", func(context *gin.Context) {
		context.JSON(200, "siucess")
	})

	//2.返回map,使用gin.H或者自定义map
	ginServer.GET("/jsonMap", func(context *gin.Context) {
		context.JSON(200, gin.H{"code": "200", "message": "成功"})

		//自定义
		m := map[string]any{}
		m["code"] = 200
		m["message"] = "成功"
		context.JSON(200, m)

	})

	//3.返回结构体
	ginServer.GET("/jsonStruct", func(context *gin.Context) {
		user := models.User{
			Username: "bobo",
			Password: "bobo123456",
		}

		context.JSON(200, user)
	})

	//4.返回切片结构体
	ginServer.GET("/jsonSlice", func(context *gin.Context) {

		user2 := make([]models.User, 2)
		context.JSON(200, user2)

	})

}
