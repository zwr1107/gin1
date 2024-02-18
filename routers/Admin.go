package routers

import (
	"gin1/controllers/admin"
	"gin1/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.AdminAuthMiddleware())
	{
		//后台首页
		adminRouters.GET("/", admin.MainController{}.Index)
		//欢迎页
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)
		//登录页
		adminRouters.GET("/login", admin.LoginController{}.Index)
		//登录操作
		adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
		//获取验证码
		adminRouters.GET("/getCaptcha", admin.LoginController{}.GetCaptcha)
		//退出登录
		adminRouters.GET("/loginOut", admin.LoginController{}.LoginOut)

		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

	}
}
vController{}.Edit)
		navRouters.Any("/del", admin.NavController{}.Del)
		navRouters.Any("/getinfo", admin.NavController{}.GetInfo)
	}

	//学生路由
	studentRouters := r.Group("/admin/student")
	{
		studentRouters.GET("/", admin.StudentController{}.GetList)
		studentRouters.GET("/add", admin.StudentController{}.Add)
		studentRouters.GET("/edit", admin.StudentController{}.Edit)
		studentRouters.GET("/del", admin.StudentController{}.Del)
	}
}
