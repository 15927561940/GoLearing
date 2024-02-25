package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Params(ginServer *gin.Engine) {

	////1.获取传统的k,v的方式
	//ginServer.GET("/params/info", func(context *gin.Context) {
	//	userid := context.Query("userid")
	//	username := context.Query("username")
	//
	//	//2.返回,,,后续要做啥根据业务逻辑来
	//	context.JSON(200, gin.H{"userid": userid, "username": username})
	//})

	//2.restful风格api获取请求参数
	//ginServer.GET("/params/info/:userid/:username", func(context *gin.Context) {
	//	userid := context.Param("userid")
	//	username := context.Param("username")
	//
	//	//2.返回,,,后续要做啥根据业务逻辑来
	//	context.JSON(200, gin.H{"userid": userid, "username": username})
	//})

	//3  post    form表单操作
	//ginServer.POST("/params/info/save", func(context *gin.Context) {
	//	//form表单操作
	//	username := context.PostForm("username")
	//	password := context.PostForm("password")
	//
	//	fmt.Println(username, password)
	//	context.JSON(200, gin.H{"username": username, "passwd": password})
	//})

	//4.json格式
	ginServer.POST("/params/info/json", func(context *gin.Context) {
		//GetRawData  原理是从c.request.Body中获取数据，返回值是一个[]byte  字节
		//func (c *Context) GetRawData() ([]byte, error) {
		//	return io.ReadAll(c.Request.Body)
		//}
		data, _ := context.GetRawData()

		//开始定义map或者结构体转化
		m := make(map[string]any)
		//把前台传过来的字符串转换成map
		_ = json.Unmarshal(data, &m) //操作地址

		context.JSON(200, m)
	})

}
