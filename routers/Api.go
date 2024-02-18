package routers

import "github.com/gin-gonic/gin"

// ApiRoutersInit 是一个初始化Api路由的函数
// 参数c是一个gin.Engine实例，用于设置路由
func ApiRoutersInit(c *gin.Engine) {
	apiRouters := c.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是一个api接口")
		})
		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(200, "我是一个api接口-Userlist")
		})
		apiRouters.GET("/plist", func(c *gin.Context) {
			c.String(200, "我是一个api接口-Plist")
		})
	}

}
