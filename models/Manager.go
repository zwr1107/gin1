package models

type Manager struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   int    `json:"mobile"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	RoleId   int    `json:"role_id"`
	AddTime  int    `json:"add_time"`
	IsSuper  int    `json:"is_super"`
	Role     Role   `gorm:"foreignKey:RoleId;references:Id" json:"role"`
}

// TableName 设置Manager的表名为`manager`
func (manager Manager) TableName() string {
	return "manager"
}
