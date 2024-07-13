package services

import (
	"github.com/yonatanlla/crud-go-react/db"
	"github.com/yonatanlla/crud-go-react/models"
)

func ValideUniqueEmail(email string) error {
	var user models.User
	result := db.DB.Where("email = ?", email).First(&user)

	if result.RowsAffected > 0 {
		return result.Error
	}

	return nil
}

func CreateUser(user *models.User) error {
	err := db.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
