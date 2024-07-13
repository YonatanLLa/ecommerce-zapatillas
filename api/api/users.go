package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yonatanlla/crud-go-react/models"
	"github.com/yonatanlla/crud-go-react/services"
	"golang.org/x/crypto/bcrypt"
)

func PostUser(c echo.Context) error {

	var user models.Register
	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
	}

	// Check if user already exists
	existingUser := services.ValideUniqueEmail(user.Email)
	if existingUser != nil {
		return echo.NewHTTPError(http.StatusConflict, "user already exists")
	}

	// Create user
	user.Password = string(hashedPassword)
	createdUser := services.CreateUser(&user)
	if createdUser != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
	}

	return c.JSON(http.StatusOK, user)
}

func PostLogin(c echo.Context) error {

	var login models.Login
	err := c.Bind(&login)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	loginUser := services.LoginUser(&login)

	if loginUser == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	return c.JSON(http.StatusOK, loginUser)

}
