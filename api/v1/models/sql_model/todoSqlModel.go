package sqlmodel

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title" gorm:"not null" validate:"required,min=3"`
	Completed bool   `json:"completed" gorm:"default:false" validate:"required"`
}
