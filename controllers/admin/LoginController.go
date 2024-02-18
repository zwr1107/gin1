package admin

import (
	"encoding/json"
	"fmt"
	"gin1/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})

}
func (con LoginController) DoLogin(c *gin.Context) {
	// 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	verifyValue := c.PostForm("verifyValue")
	captchaId := c.PostForm("captchaId")
	//验证验证码
	if !models.VerifyCaptcha(captchaId, verifyValue) {
		con.Error(c, "验证码错误", "/admin/login")
		return
	}

	// 验证用户名和密码
	var manager []models.Manager
	models.DB.Where("username=?", username).Find(&manager)
	if manager == nil {
		con.Error(c, "用户名不存在", "/admin/login")
		return
	}
	password = models.Md5(password)
	if password != manager[0].Password {
		con.Error(c, "密码错误", "/admin/login")
		return
	}
	// 登录成功
	// 设置session
	session := sessions.Default(c)
	//注意：session.Set没法直接保存结构体对应的切片 把结构体转换成json字符串
	userinfoSlice, _ := json.Marshal(manager)
	//fmt.Println(string(userinfoSlice))
	session.Set("userinfo", string(userinfoSlice))
	err := session.Save()
	if err != nil {
		fmt.Println("session保存失败", err)
		con.Error(c, "登录失败", "/admin/login")
		return
	}

	// 跳转
	con.Success(c, "登录成功", "/admin")

}

// GetCaptcha 获取验证码
func (con LoginController) GetCaptcha(c *gin.Context) {
	// 生成验证码,返回验证码id,base64图片字符串
	id, base64, err := models.GetCaptcha()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "获取验证码失败",
		})
		return
	}
	//定义一个map，用来存储验证码id和base64图片字符串
	data := make(map[string]interface{})
	data["id"] = id
	data["image"] = base64
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// LoginOut 退出登录
func (con LoginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.Success(c, "退出登录成功", "/admin/login")
}
