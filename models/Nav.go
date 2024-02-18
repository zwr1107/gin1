package models

type Nav struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Url    string `json:"url"`
	Sort   int    `json:"sort"`
}

// TableName 设置Nav的表名为`nav`
func (nav Nav) TableName() string {
	return "nav"
}
