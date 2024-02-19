package admin

import (
	"fmt"
	"gin1/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

var roleData []models.Role

func (con ManagerController) Index(c *gin.Context) {
	//查询数据
	var data []models.Manager
	models.DB.Preload("Role").Find(&data)
	//fmt.Println(data)
	//渲染模板
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"data": data,
	})

}
func (con ManagerController) Add(c *gin.Context) {
	//查询角色数据
	models.DB.Find(&roleData)

	//fmt.Println(roleData)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleData": roleData,
	})
}

// DoAdd 添加管理员
func (con ManagerController) DoAdd(c *gin.Context) {
	//获取参数
	username := strings.Trim(c.PostForm("username"), " ")
	password := c.PostForm("password")
	roleId := c.PostForm("role_id")
	email := c.PostForm("email")
	mobile := c.PostForm("mobile")

	//验证参数
	if username == "" {
		con.Error(c, "用户名不能为空", "")
		return
	}
	if password == "" {
		con.Error(c, "密码不能为空", "")
		return
	}
	//密码长度验证
	if len(password) < 6 {
		con.Error(c, "密码长度不能小于6位", "")
		return
	}

	if roleId == "" {
		con.Error(c, "角色不能为空", "")
		return
	}
	//验证用户名是否存在
	var existManager models.Manager
	models.DB.Where("username=?", username).First(&existManager)
	//fmt.Println(existManager.Id)
	if existManager.Id > 0 {
		con.Error(c, "用户名已存在", "/admin/manager/add")
		return
	}

	//添加数据
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		RoleId:   models.StrToInt(roleId),
		AddTime:  int(models.GetTime()),
		IsSuper:  0,
		Status:   1,
		Email:    email,
		Mobile:   models.StrToInt(mobile),
	}
	err := models.DB.Create(&manager).Error
	if err != nil {
		con.Error(c, "添加失败", "")
		return
	}
	con.Success(c, "添加成功", "/admin/manager")
}

// Edit 编辑管理员
func (con ManagerController) Edit(c *gin.Context) {
	//获取参数
	id := c.Query("id")
	fmt.Println(id)
	//查询数据
	var data models.Manager
	models.DB.Where("id=?", id).First(&data)
	if data.Id == 0 {
		con.Error(c, "数据不存在", "/admin/manager")
		return
	}

	models.DB.Find(&roleData)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"data":     data,
		"roleData": roleData,
	})
}

// DoEdit 编辑管理员
func (con ManagerController) DoEdit(c *gin.Context) {
	//获取参数
	id := c.PostForm("id")
	username := strings.Trim(c.PostForm("username"), " ")
	roleId := c.PostForm("role_id")
	email := c.PostForm("email")
	mobile := c.PostForm("mobile")
	isSuper := c.PostForm("is_super")
	status := c.PostForm("status")
	password := c.PostForm("password")

	//验证参数
	if username == "" {
		con.Error(c, "用户名不能为空", "")
		return
	}
	if roleId == "" {
		con.Error(c, "角色不能为空", "")
		return
	}

	var data models.Manager
	if len(password) > 0 {
		//密码长度验证
		if len(password) < 6 {
			con.Error(c, "密码长度不能小于6位", "")
			return
		}
		//更新密码
		password = models.Md5(password)
		data.Password = password

	}

	//查询数据

	models.DB.Where("id=?", id).First(&data)

	//更新数据
	data.Username = username
	data.RoleId = models.StrToInt(roleId)
	data.Email = email
	data.Mobile = models.StrToInt(mobile)
	data.IsSuper = models.StrToInt(isSuper)
	data.Status = models.StrToInt(status)

	err := models.DB.Save(&data).Error
	if err != nil {
		con.Error(c, "编辑失败", "")
		return
	}
	con.Success(c, "编辑成功", "/admin/manager/edit/"+id)
}

// Delete 删除管理员
func (con ManagerController) Delete(c *gin.Context) {
	//获取参数
	id := c.Query("id")

	//删除数据
	err := models.DB.Where("id=?", id).Delete(&models.Manager{}).Error
	if err != nil {
		con.Error(c, "删除失败", "")
		return
	}
	con.Success(c, "删除成功", "/admin/manager/")
}
