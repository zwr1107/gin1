package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (con BaseController) Success(c *gin.Context, message, redirectUrl string) {
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "success",
	//	"data": data,
	//})
	fmt.Println("message", message, "redirect", redirectUrl)
	c.HTML(200, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (con BaseController) Error(c *gin.Context, message string, redirectUrl string) {

	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
