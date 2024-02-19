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

		//后台管理员
		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)
		adminRouters.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)
		adminRouters.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)

		//轮播图
		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

		//角色管理
		adminRouters.GET("/role", admin.RoleController{}.Index)
		adminRouters.GET("/role/add", admin.RoleController{}.Add)
		//角色添加操作
		adminRouters.POST("/role/doAdd", admin.RoleController{}.DoAdd)
		//角色编辑
		adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
		//角色编辑操作
		adminRouters.POST("/role/doEdit", admin.RoleController{}.DoEdit)
		//角色删除
		adminRouters.GET("/role/delete", admin.RoleController{}.Delete)
		//角色权限
		adminRouters.GET("/role/auth", admin.RoleController{}.Auth)
		//角色权限操作
		adminRouters.POST("/role/doAuth", admin.RoleController{}.DoAuth)

		//权限管理
		adminRouters.GET("/access", admin.AccessController{}.Index)
		adminRouters.GET("/access/add", admin.AccessController{}.Add)
		adminRouters.POST("/access/doAdd", admin.AccessController{}.DoAdd)
		adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
		adminRouters.POST("/access/doEdit", admin.AccessController{}.DoEdit)
		adminRouters.GET("/access/delete", admin.AccessController{}.Delete)

	}

}
