package routers

import (
	"gin1/controllers/home"

	"github.com/gin-gonic/gin"
)

func WebRoutersInit(r *gin.Engine) {
	webRouters := r.Group("/")
	{
		webRouters.GET("/", home.IndexController{}.Index)
		webRouters.GET("/userlist", func(c *gin.Context) {
			c.String(200, "我是一个web接口-Userlist")
		})
		webRouters.GET("/plist", func(c *gin.Context) {
			c.String(200, "我是一个web接口-Plist")
		})

		//配置es
		webRouters.GET("/es/index", home.SearchController{}.Index)
		webRouters.GET("/es/add", home.SearchController{}.AddGoods)
	}
}
