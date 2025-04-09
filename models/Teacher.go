package models

type Teacher struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Subject    string    `json:"subject"`
	Experience int       `json:"experience"`
	PhotoURL   string    `json:"photo_url"`
	Students   []Student `gorm:"many2many:teacher_students;" json:"students,omitempty"` // связь many-to-many
}
