package models

type LessonStudent struct {
	LessonId  int     `json:"lesson_id"`
	StudentId int     `json:"student_id"`
	Student   Student `gorm:"ForeignKey:StudentId;AssociationForeignKey:Id" json:"student"`
}
