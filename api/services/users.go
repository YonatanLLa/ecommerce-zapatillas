package services

import (
	"github.com/yonatanlla/crud-go-react/db"
	"github.com/yonatanlla/crud-go-react/models"
	"golang.org/x/crypto/bcrypt"
)

func ValideUniqueEmail(email string) error {
	var user models.Register
	result := db.DB.Where("email = ?", email).First(&user)
	println(result.RowsAffected, "rows affected")
	if result.RowsAffected > 0 {
		return result.Error
	}

	return nil
}

func CreateUser(user *models.Register) error {
	err := db.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(user *models.Login) *models.Register {
	var loginUser models.Register
	err := db.DB.Where("email = ?", user.Email).First(&loginUser).Error
	if err != nil {
		return nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(user.Password))
	if err != nil {
		return nil
	}
	return &loginUser
}
