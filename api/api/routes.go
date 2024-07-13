package api

import (
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.POST("/", PostUser)

}
