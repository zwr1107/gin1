package main

import (
	"gin1/models"
	"gin1/routers"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// MiddleWare 中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request", "中间件")
		c.Next()
	}
}

func main() {
	//gin框架默认是开发模式，设置为发布模式
	gin.SetMode(gin.ReleaseMode)
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"TimeToDate": models.TimeToDate,
	})
	//定义模板
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	//配置全局中间件
	//r.Use(MiddleWare())
	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	routers.ApiRoutersInit(r)
	routers.WebRoutersInit(r)
	routers.AdminRoutersInit(r)

	// 启动HTTP 服务，默认在0.0.0.0:8080 启动服务
	r.Run()
}
