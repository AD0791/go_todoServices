package sqlmodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"` // Hashed & salted password
}
