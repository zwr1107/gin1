package home

import (
	"fmt"
	"gin1/models"
	"github.com/gin-gonic/gin"
	"time"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	fmt.Println(models.TimeToDate(int(time.Now().Unix())))
	c.HTML(200, "default/index.html", gin.H{
		"title": "首页",
	})
}
