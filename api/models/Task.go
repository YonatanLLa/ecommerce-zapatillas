package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      int    `json:"user_id"`
}
