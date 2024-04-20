package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TimeToDate 时间戳转时间
func TimeToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// GetTime 获取时间戳
func GetTime() int64 {
	return time.Now().Unix()
}

// GetTimeStr 获取当前时间
func GetTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Md5 MD5加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// StrToInt 字符串 转整形
func StrToInt(str string) int {
	var result int
	_, err := fmt.Sscanf(str, "%d", &result)
	if err != nil {
		return 0
	}
	return result
}

// InArray 判断元素是否在数组中
func InArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

// 定义消息结构体
type SuccessData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 返回成功
func Success(msg string, data interface{}) SuccessData {
	return SuccessData{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
}

// Error 返回失败
func Error(code int, msg string) SuccessData {
	return SuccessData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func JsonResponse(c *gin.Context, code int, message string, data interface{}) {
	response := SuccessData{
		Code: code,
		Msg:  message,
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}
