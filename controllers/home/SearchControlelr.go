package home

import (
	"context"
	"fmt"
	"gin1/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
}

// 初始化的时候判断goods是否存在  创建索引配置映射
func (con SearchController) Index(c *gin.Context) {

	exists, err := models.EsClient.IndexExists("goods").Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	print(exists)
	if !exists {
		// 配置映射
		mapping := `
		{
			"settings": {
			  "number_of_shards": 1,
			  "number_of_replicas": 0
			},
			"mappings": {
			  "properties": {
				"Content": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
				},
				"Title": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
				}
			  }
			}
		  }
		`
		//注意：增加的写法
		_, err := models.EsClient.CreateIndex("goods").Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			fmt.Println(err)
		}
	}

	c.String(200, "创建索引配置映射成功")
}

// 增加商品数据
func (con SearchController) AddGoods(c *gin.Context) {
	goods := []models.Goods{}
	models.DB.Find(&goods)

	addResult, err := models.EsClient.Index().
		Index("goods").
		Id(strconv.Itoa(goods[1].Id)).
		BodyJson(goods[1]).
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", addResult.Id, addResult.Index, addResult.Type)

	models.JsonResponse(c, 200, "增加商品数据 success", nil)
}

// 更新数据
func (con SearchController) UpdateGoods(c *gin.Context) {

	goods := []models.Goods{}
	models.DB.Find(&goods)
	goods[0].Title = "我是修改后的数据"
	goods[0].GoodsContent = "我是修改后的数据GoodsContent"

	_, err := models.EsClient.Update().
		Index("goods").
		Id("19").
		Doc(goods[0]).
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	c.String(200, "修改数据 success")
}

// 删除
func (con SearchController) DeleteGoods(c *gin.Context) {

	_, err := models.EsClient.Delete().
		Index("goods").
		Id("19").
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	c.String(200, "删除成功 success")
}

// 查询一条数据
func (con SearchController) GetOne(c *gin.Context) {
	c.String(200, "GetOne")
}

func (con SearchController) Query(c *gin.Context) {
	c.String(200, "Query")
}

// 条件筛选查询
func (con SearchController) FilterQuery(c *gin.Context) {
	c.String(200, "filter Query")
}

// 分页查询
func (con SearchController) PagingQuery(c *gin.Context) {
	c.String(200, "filter 分页查询")
}
