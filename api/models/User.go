package models

import "gorm.io/gorm"

type Register struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type Login struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
