package models

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// TableName 设置User的表名为`user`
func (user User) TableName() string {
	return "user"
}
