package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/manakuro/golang-clean-architecture/interface/controllers"
)

func NewRouter(e *echo.Echo, c controllers.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })

	return e
}
