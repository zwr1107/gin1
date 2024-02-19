package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {

	fmt.Println("后台首页")
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
