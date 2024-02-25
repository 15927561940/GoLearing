package main

import (
	"encoding/gob"
	"fmt"
	"gin/controllers"
	"gin/filters"
	"gin/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("gin 来了")

	//1创建ginserver
	ginServer := gin.Default()

	//2..创建cookie存储
	store := cookie.NewStore([]byte("secret"))
	ginServer.Use(sessions.Sessions("mysession", store))

	//注册modle,,,,modle结构体要被管理起来
	gob.Register(models.User{})

	//2注册中间件
	filters.RegFilter(ginServer)

	//2.创建服务  fun一打就出来了
	controllers.User(ginServer)
	controllers.Index(ginServer) //把ginServer传进来，和indexControllers的方法对应
	controllers.Json(ginServer)
	controllers.Params(ginServer)
	controllers.Redirect(ginServer)
	controllers.Error(ginServer)
	controllers.Course(ginServer)

	controllers.Login(ginServer) //定义上要用的模块

	ginServer.LoadHTMLGlob("template/*") //加载静态网页模板的位置

	//4.开始配置静态资源地址，作用是告诉gin这个是静态资源    直接下载即可  不用去找动态路由地址
	ginServer.Static("/static", "./static") //访问网址      本地文件夹

	ginServer.Run(":8091")
}
