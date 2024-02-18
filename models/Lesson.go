package models

type Lesson struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Student []Student `gorm:"many2many:lesson_student" json:"student"`
}

// TableName 设置Lesson的表名为`lesson`
func (lesson Lesson) TableName() string {
	return "lesson"
}
