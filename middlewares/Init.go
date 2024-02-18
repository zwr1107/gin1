package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Init(c *gin.Context) {
	fmt.Println("初始化中间件")
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
}
