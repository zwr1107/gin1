package admin

import (
	"fmt"
	"gin1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessController struct {
	BaseController
}

// Index 角色列表页
func (con AccessController) Index(c *gin.Context) {

	//查询数据
	var data []models.Access
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&data)

	//fmt.Println(data)
	//渲染模板
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"data": data,
	})
}

func (con AccessController) Add(c *gin.Context) {
	//获取顶级模块
	var data []models.Access
	models.DB.Where("module_id=?", 0).Find(&data)

	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"data": data,
	})
}
func (con AccessController) DoAdd(c *gin.Context) {
	//获取表单数据
	moduleName := c.PostForm("module_name")
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId := c.PostForm("module_id")
	sort := c.PostForm("sort")
	description := c.PostForm("description")
	status := c.PostForm("status")

	if moduleName == "" || actionName == "" || url == "" || moduleId == "" || sort == "" || description == "" || status == "" {
		con.Error(c, "数据不完整", "/admin/access/add")
		return
	}

	//插入数据
	access := models.Access{
		ModuleName:  moduleName,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    models.StrToInt(moduleId),
		Sort:        models.StrToInt(sort),
		Description: description,
		Status:      models.StrToInt(status),
		AddTime:     int(models.GetTime()),
	}
	models.DB.Create(&access)

	//跳转
	con.Success(c, "添加成功", "/admin/access")

}

func (con AccessController) Edit(c *gin.Context) {
	//获取参数
	id := c.Query("id")
	if id == "" {
		con.Error(c, "参数错误", "")
		return
	}

	//查询数据
	var data models.Access
	models.DB.Where("id=?", id).First(&data)

	fmt.Println(data)

	//获取顶级模块
	var moduleData []models.Access
	models.DB.Where("module_id=?", 0).Find(&moduleData)

	//渲染模板
	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"data":       data,
		"moduleData": moduleData,
	})

}

func (con AccessController) DoEdit(c *gin.Context) {
	//获取表单数据
	id := c.PostForm("id")
	moduleName := c.PostForm("module_name")
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId := c.PostForm("module_id")
	sort := c.PostForm("sort")
	description := c.PostForm("description")
	status := c.PostForm("status")

	if moduleName == "" || actionName == "" || url == "" || moduleId == "" || sort == "" || description == "" || status == "" {
		con.Error(c, "数据不完整", "/admin/access/edit?id="+id)
		return
	}

	//更新数据
	models.DB.Model(&models.Access{}).Where("id=?", id).Updates(map[string]interface{}{
		"module_name": moduleName,
		"action_name": actionName,
		"url":         url,
		"module_id":   models.StrToInt(moduleId),
		"sort":        models.StrToInt(sort),
		"description": description,
		"status":      models.StrToInt(status),
	})

	//跳转
	con.Success(c, "修改成功", "/admin/access")
}

func (con AccessController) Delete(c *gin.Context) {
	//获取参数
	id := c.Query("id")
	if id == "" {
		con.Error(c, "参数错误", "")
		return
	}

	//判断是否有子节点
	var count int64
	models.DB.Model(&models.Access{}).Where("module_id=?", id).Count(&count)
	if count > 0 {
		con.Error(c, "有子节点,不能删除", "")
		return
	}

	//删除数据
	models.DB.Where("id=?", id).Delete(&models.Access{})

	//跳转
	con.Success(c, "删除成功", "/admin/access")

}
