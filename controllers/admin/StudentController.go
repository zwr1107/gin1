package admin

import (
	"gin1/models"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	BaseController
}

func (con StudentController) GetList(c *gin.Context) {
	resData := []models.Student{}
	models.DB.Preload("Lesson").Find(&resData)
	con.Success(c, resData)
}

func (con StudentController) Add(c *gin.Context) {
	con.Success(c, "我是一个学生接口-添加")
}

func (con StudentController) Edit(c *gin.Context) {
	con.Success(c, "我是一个学生接口-编辑")
}

func (con StudentController) Del(c *gin.Context) {
	con.Success(c, "我是一个学生接口-删除")
}
