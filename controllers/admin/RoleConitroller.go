package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleController struct {
	BaseController
}

// 角色列表页
func (con RoleController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{})
}

func (con RoleController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
