package models

type Article struct {
	Id          int
	Title       string
	CateId      int
	State       int
	ArticleCate ArticleCate `gorm:"ForeignKey:CateId;AssociationForeignKey:Id"`
}

// TableName 设置Article的表名为`article`
func (article Article) TableName() string {
	return "article"
}
