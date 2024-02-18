package admin

import (
	"gin1/models"
	"github.com/gin-gonic/gin"
)

type NavController struct {
	BaseController
}

// GetList 获取列表
func (con NavController) GetList(c *gin.Context) {
	//查询所有数据
	navList := []models.Nav{}
	models.DB.Find(&navList)

	// in 查询
	//var ids []int
	models.DB.Model(&models.Nav{}).Where("id in (?)", []int{1, 2, 3}).Find(&navList)
	//between and 查询
	models.DB.Model(&models.Nav{}).Where("id between ? and ?", 1, 3).Find(&navList)
	//or 查询
	models.DB.Model(&models.Nav{}).Where("id = ? or id = ?", 1, 3).Find(&navList)
	models.DB.Model(&models.Nav{}).Where("id = 1").Or("id = 3").Find(&navList)
	//like 查询
	models.DB.Model(&models.Nav{}).Where("title like ?", "%首页%").Find(&navList)
	//排序
	models.DB.Model(&models.Nav{}).Order("id desc").Find(&navList)
	//limit
	models.DB.Model(&models.Nav{}).Limit(2).Find(&navList)
	//offset
	models.DB.Model(&models.Nav{}).Offset(2).Find(&navList)
	//指定返回字段
	var fields []map[string]interface{}
	models.DB.Model(&models.Nav{}).Select("id,title").Find(&fields)
	//查询单个字段
	var title string
	models.DB.Model(&models.Nav{}).Where("id = 1").Pluck("title", &title)
	//查询多个字段
	var result []map[string]interface{}
	models.DB.Model(&models.Nav{}).Where("id = 1").Select("id,title").Find(&result)
	//分页
	var pageList []models.Nav
	models.DB.Model(&models.Nav{}).Limit(2).Offset(2).Find(&pageList)
	//查询总数
	var total int64
	navList = []models.Nav{}
	models.DB.Model(&models.Nav{}).Count(&total).Find(&navList)
	//原生SQL
	var sql = "select * from nav where id = ?"
	sqlData := []models.Nav{}
	models.DB.Raw(sql, 1).Scan(&sqlData)
	var delSql = "delete from nav where id = ?"
	models.DB.Exec(delSql, 1)

	//统计表数据
	var count int64
	models.DB.Raw("select count(*) as count from nav").Scan(&count)

	//返回数据
	c.JSON(200, gin.H{
		"code":   0,
		"msg":    "获取成功",
		"data":   navList,
		"res":    result,
		"fields": fields,
		"page":   pageList,
		"total":  total,
		"sql":    sqlData,
		"count":  count,
	})
}

// GetInfo 获取详情
func (con NavController) GetInfo(c *gin.Context) {
	result := models.Nav{Id: 1}
	models.DB.First(&result)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": result,
	})
}

// Add 添加
func (con NavController) Add(c *gin.Context) {

}

// Edit 编辑
func (con NavController) Edit(c *gin.Context) {

}

// Del 删除
func (con NavController) Del(c *gin.Context) {

}
