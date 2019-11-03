package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/manakuro/golang-clean-architecture/interface/controllers"
)

func NewRouter(e *echo.Echo, c controllers.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", c.GetUsers)

	return e
}
