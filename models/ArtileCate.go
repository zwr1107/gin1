package models

// ArticleCate references 是主键关联 默认是ID
// ArticleCate 定义一个结构体
type ArticleCate struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	State   int       `json:"state"`
	Article []Article `gorm:"ForeignKey:CateId;AssociationForeignKey:Id" json:"article"`
}

// TableName 设置ArticleCate的表名为`article_cate`
func (articleCate ArticleCate) TableName() string {
	return "article_cate"
}
