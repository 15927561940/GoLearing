package filters

import (
	"fmt"
	"gin/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MyHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		session := sessions.Default(context)
		sessionUser := session.Get("user")
		user := sessionUser.(models.User) //将session转换成user才能被读取

		//可以通过两个方法来决定你的请求是继续Next还是终止Abort

		//放入全局数据信息
		context.Set("userId", 1200000)

		fmt.Println("中间件开始")

		//判断是否能进行下一步
		if user.Username != "bobo" {
			//判断逻辑，数据存在则往下走
			context.Abort()
			context.JSON(200, gin.H{"message": "无权限"})
		} else {
			context.Next()
		}

		fmt.Println("中间件结束")
	}
}

func RegFilter(ginServer *gin.Engine) {

	//ginServer.Use(myHandler()) //注册中间件---全局注册---统一鉴权或者拦截处理

}
