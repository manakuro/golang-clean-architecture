package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/manakuro/golang-clean-architecture/infrastructure/api/handler"
)

func NewRouter(e *echo.Echo, handler handler.AppHandler) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", handler.GetUsers)

	return e
}
