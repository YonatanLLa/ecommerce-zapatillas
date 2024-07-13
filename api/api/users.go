package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yonatanlla/crud-go-react/models"
	"github.com/yonatanlla/crud-go-react/services"
)

func PostUser(c echo.Context) error {

	var user models.User

	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	existingUser := services.ValideUniqueEmail(user.Email)

	if existingUser != nil {
		return echo.NewHTTPError(http.StatusConflict, "user already exists")
	}
	createdUser := services.CreateUser(&user)

	if createdUser != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
	}

	return c.JSON(http.StatusOK, user)

}
