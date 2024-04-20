package api

import (
	"gin1/models"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	//空的结构体
	emptyMap := make(map[string]interface{})
	models.JsonResponse(c, 200, "我是一个api接口", emptyMap)
}
