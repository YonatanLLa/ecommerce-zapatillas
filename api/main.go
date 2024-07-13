package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yonatanlla/crud-go-react/api"
	"github.com/yonatanlla/crud-go-react/db"
	"github.com/yonatanlla/crud-go-react/models"
)

func main() {
	db.DBConnect()

	db.DB.AutoMigrate(&models.Task{})
	db.DB.AutoMigrate(&models.User{})

	e := echo.New()
	api.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
