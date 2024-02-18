package middlewares

import (
	"encoding/json"
	"gin1/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//进行权限判断 没有登录的用户 不能进入后台管理中心
		//获取访问地址
		//fmt.Println(c.Request.URL.Path)
		pathUrl := strings.Split(c.Request.URL.Path, "?")[0]
		//获取session
		session := sessions.Default(c)
		userinfo := session.Get("userinfo")
		//类型断言 来判断 userinfo是不是一个string
		userinfoStr, ok := userinfo.(string)
		if ok {
			var userinfoStruct []models.Manager
			err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
			if err != nil || len(userinfoStruct) == 0 {
				c.Redirect(302, "/admin/login")
			}

			//已经登录
			if pathUrl == "/admin/login" {
				c.Redirect(302, "/admin")
				c.Abort()
			}
		} else {
			//没有登录 不能访问后台.除了登录页和登录操作
			if pathUrl != "/admin/login" && pathUrl != "/admin/doLogin" && pathUrl != "/admin/getCaptcha" {
				c.Redirect(302, "/admin/login")
				c.Abort()
			}
		}

	}
}
