package admin

import (
	"gin1/models"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UserController struct{}

func (con UserController) GetList(c *gin.Context) {

	userList := []models.User{}

	var where = make(map[string]interface{})
	where = map[string]interface{}{"age": 18}

	//获取用户列表
	models.DB.Where(where).Find(&userList)
	//total := userList.RowsAffected

	//随机生成100条数据
	//for i := 0; i < 100; i++ {
	//	user := models.User{
	//		Username: "test" + string(i),
	//		// //随机生成1-100的年龄
	//		Age:     i + 1,
	//		Email:   "test" + string(i) + "@qq.com",
	//		AddTime: int(time.Now().Unix()),
	//	}
	//	models.DB.Create(&user)
	//}

	//查询总数
	var total int64
	models.DB.Model(&models.User{}).Where(where).Count(&total)

	//返回数据
	c.JSON(200, gin.H{
		"code":  0,
		"msg":   "获取成功",
		"data":  userList,
		"total": total,
	})
}

func (con UserController) Add(c *gin.Context) {

	//添加数据
	user := models.User{

		Username: "test",
		Age:      18,
		Email:    "7566@qq.com",
		AddTime:  int(time.Now().Unix()),
	}
	models.DB.Create(&user)

	//获取自增id
	id := user.Id
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "添加成功",
		"id":   id,
		"data": user,
	})

	//c.HTML(200, "admin/user/add.html", gin.H{})
}

// DoAdd 添加操作
func (con UserController) DoAdd(c *gin.Context) {

	//获取参数
	username := c.PostForm("username")

	//头像1
	//file, err := c.FormFile("avatar[]")
	////头像2
	//file2, err := c.FormFile("avatar2")
	////头像3
	//file3, err := c.FormFile("avatar3")

	var fileData []string

	//多文件上传
	files := c.Request.MultipartForm.File["avatar[]"]
	for _, file := range files {

		//指定目录
		dir := "./static/upload/" + time.Now().Format("2006-01-02")
		if !IsExist(dir) {
			err := CreateDir(dir)
			if err != nil {
				c.JSON(200, gin.H{
					"code":  1,
					"msg":   "创建目录失败",
					"error": err,
				})
				return
			}
		}

		//获取文件后缀
		ext := filepath.Ext(file.Filename)
		//后缀转小写
		ext = strings.ToLower(ext)
		//允许的后缀 jpg,png,jpeg
		allowExt := map[string]bool{
			".jpg":  true,
			".png":  true,
			".jpeg": true,
			".gif":  true,
		}

		//判断文件后缀是否允许
		if _, ok := allowExt[ext]; !ok {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "文件后缀不允许",
			})
		}

		//文件名
		fileName := time.Now().Format("20060102150405") + ext

		// 上传文件到指定的目录
		filePath := dir + "/" + fileName

		fileData = append(fileData, filePath)

		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "上传失败",
			})
			return
		}
	}

	//	// 上传文件到指定的目录
	//	err = c.SaveUploadedFile(file, filePath)
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"code": 1,
	//			"msg":  "上传失败",
	//		})
	//		return
	//	}
	//}
	//
	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"code": 1,
	//		"msg":  "上传失败",
	//	})
	//	return
	//}
	//
	//dir := "upload/" + time.Now().Format("2006-01-02")
	//
	////获取文件后缀
	//ext := filepath.Ext(file.Filename)
	//fileName := time.Now().Format("20060102150405") + ext

	////获取文件后缀
	//ext2 := filepath.Ext(file2.Filename)
	//fileName2 := time.Now().Format("20060102150405") + ext2
	//
	////获取文件后缀
	//ext3 := filepath.Ext(file3.Filename)
	//fileName3 := time.Now().Format("20060102150405") + ext3

	// 上传文件到指定的目录
	//filePath := dir + "/" + fileName

	//filePath2 := dir + "/" + fileName2
	//
	//filePath3 := dir + "/" + fileName3

	//判断目录是否存在
	//if !IsExist(dir) {
	//	err := CreateDir(dir)
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"code": 1,
	//			"msg":  "创建目录失败",
	//		})
	//		return
	//	}
	//}

	// 上传文件到指定的目录

	//err = c.SaveUploadedFile(file, filePath)
	//err = c.SaveUploadedFile(file2, filePath2)
	//err = c.SaveUploadedFile(file3, filePath3)

	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"code": 1,
	//		"msg":  "上传失败",
	//	})
	//	return
	//}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "上传成功",
		"data": gin.H{
			"username": username,
			"avatar":   fileData,
		},
	})
}

// CreateDir 创建目录
func CreateDir(s string) interface{} {
	err := os.MkdirAll(s, 0777)
	if err != nil {
		return err
	}
	return nil
}

// IsExist 判断目录是否存在
func IsExist(s string) bool {
	_, err := os.Stat(s)
	if err != nil {
		return os.IsExist(err)
	}
	return true

}

func (con UserController) Edit(c *gin.Context) {
	//c.String(200, "我是一个admin接口-编辑")

	//获取get参数
	id := c.Param("id")
	//修改数据
	user := models.User{}
	models.DB.Where("id = ?", id).First(&user)

	//更新数据
	models.DB.Model(&user).Where("id = ?", id).Update("username", "test111")

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": user,
		"id":   id,
	})
}

func (con UserController) Del(c *gin.Context) {
	//c.String(200, "我是一个admin接口-删除")

	//删除id为1的数据
	models.DB.Where("id = ?", 1).Delete(&models.User{})

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "删除成功",
	})
}
