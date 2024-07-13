package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"`
}
