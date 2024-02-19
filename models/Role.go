package models

type Role struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	AddTime     int    `json:"add_time"`
}

func (Role) TableName() string {
	return "role"
}
