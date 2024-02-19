package admin

import (
	"fmt"
	"gin1/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

// Index 角色列表页
func (con RoleController) Index(c *gin.Context) {

	//查询数据
	var data []models.Role
	models.DB.Find(&data)
	fmt.Println(data)
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 0,
	//	"msg":  "获取成功",
	//	"data": data,
	//})
	//渲染模板
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"data": data,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func (con RoleController) DoAdd(c *gin.Context) {
	//获取参数
	title := c.PostForm("title")
	description := c.PostForm("description")
	//验证参数
	if title == "" {
		con.Error(c, "标题不能为空", "")
		return
	}
	if description == "" {
		con.Error(c, "描述不能为空", "")
		return
	}

	//添加数据
	role := models.Role{
		Title:       title,
		Description: description,
		Status:      1,
		AddTime:     int(models.GetTime()),
	}
	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "添加失败", "")
		return
	}
	con.Success(c, "添加成功", "/admin/role")
}

func (con RoleController) Edit(c *gin.Context) {
	//获取参数
	id := c.Query("id")
	//查询数据
	role := models.Role{}
	models.DB.Where("id=?", id).First(&role)
	//渲染模板
	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role": role,
	})
}

func (con RoleController) DoEdit(c *gin.Context) {
	//获取参数
	id := c.PostForm("id")
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	//验证参数
	if title == "" {
		con.Error(c, "标题不能为空", "")
		return
	}
	if description == "" {
		con.Error(c, "描述不能为空", "")
		return
	}
	//更新数据
	role := models.Role{}
	models.DB.Where("id=?", id).First(&role)
	role.Title = title
	role.Description = description
	err := models.DB.Save(&role).Error
	if err != nil {
		con.Error(c, "编辑失败", "")
		return
	}
	con.Success(c, "编辑成功", "/admin/role/edit?id="+id)
}

func (con RoleController) Delete(c *gin.Context) {
	//获取参数
	id := c.Query("id")
	//删除数据
	err := models.DB.Where("id=?", id).Delete(&models.Role{}).Error
	if err != nil {
		con.Error(c, "删除失败", "")
		return
	}
	con.Success(c, "删除成功", "/admin/role")
}

// Auth 授权页面
func (con RoleController) Auth(c *gin.Context) {
	//获取参数
	id := c.Query("id")
	//获取所有权限
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	//获取当前角色的权限
	roleAccessList := []models.RoleAccess{}
	models.DB.Where("role_id=?", id).Find(&roleAccessList)
	//fmt.Println(roleAccessList)
	//使用数组存储当前角色的权限id
	roleAccessArr := make(map[int]int)
	for _, v := range roleAccessList {
		roleAccessArr[v.AccessId] = v.AccessId
	}
	//判断当前权限是否在当前角色的权限中
	for k, v := range accessList {
		for k1, v1 := range v.AccessItem {
			if _, ok := roleAccessArr[v1.Id]; ok {
				accessList[k].AccessItem[k1].Checked = "checked"
			}
		}

		//判断当前模块是否在当前角色的权限中
		if _, ok := roleAccessArr[v.Id]; ok {
			accessList[k].Checked = "checked"
		}

	}

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"accessList": accessList,
		"roleId":     id,
	})

}

// 设置权限
func (con RoleController) DoAuth(c *gin.Context) {
	//获取参数
	roleId := c.PostForm("role_id")
	accessIds := c.PostFormArray("access_node[]")
	//删除当前角色的所有权限
	models.DB.Where("role_id=?", roleId).Delete(&models.RoleAccess{})
	//添加新的权限
	for _, v := range accessIds {
		accessId, _ := strconv.Atoi(v)
		roleAccess := models.RoleAccess{
			RoleId:   models.StrToInt(roleId),
			AccessId: accessId,
		}
		models.DB.Create(&roleAccess)
	}

	con.Success(c, "授权成功", "/admin/role/auth?id="+roleId)

}
