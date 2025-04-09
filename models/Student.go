package models

type Student struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Grade    string    `json:"grade"`
	Teachers []Teacher `gorm:"many2many:teacher_students;" json:"teachers,omitempty"` // связь с учителями
}
