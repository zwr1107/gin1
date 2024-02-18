package admin

import (
	"gin1/models"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) GetList(c *gin.Context) {
	//查询所有数据
	var articleList []models.Article
	models.DB.Preload("ArticleCate").Find(&articleList)

	//c.JSON(200, gin.H{
	//	"code": 0,
	//	"msg":  "获取成功",
	//	"data": articleList,
	//})

	//获取文章分类
	var articleCateList []models.ArticleCate
	models.DB.Preload("Article").Find(&articleCateList)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": articleCateList,
	})

}

func (con ArticleController) Add(c *gin.Context) {
	con.Success(c, "我是一个文章接口-添加")
}

func (con ArticleController) Edit(c *gin.Context) {
	con.Success(c, "我是一个文章接口-编辑")
}

func (con ArticleController) Del(c *gin.Context) {
	con.Success(c, "我是一个文章接口-删除")
}
