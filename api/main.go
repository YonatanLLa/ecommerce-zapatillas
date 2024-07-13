package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yonatanlla/crud-go-react/db"
	"github.com/yonatanlla/crud-go-react/models"
)

func main() {
	db.DBConnect()

	db.DB.AutoMigrate(&models.Task{})
	db.DB.AutoMigrate(&models.User{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
