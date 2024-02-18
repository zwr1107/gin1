package models

// Student 定义一个结构体
//
//many2many:lesson_student 为多对多关联表
type Student struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Pass    string   `json:"pass"`
	ClassId int      `json:"class_id"`
	Number  int      `json:"number"`
	Lesson  []Lesson `gorm:"many2many:lesson_student" json:"lesson"`
}

// TableName 设置Student的表名为`student`
func (student Student) TableName() string {
	return "student"
}
