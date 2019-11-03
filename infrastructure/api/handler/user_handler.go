package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/manakuro/golang-clean-architecture/interface/controllers"
)

type userHandler struct {
	userController controllers.UserController
}

type UserHandler interface {
	GetUsers(c echo.Context) error
}

func NewUserHandler(uc controllers.UserController) UserHandler {
	return &userHandler{userController: uc}
}

func (uh *userHandler) GetUsers(c echo.Context) error {
	u, err := uh.userController.GetUsers()
	if err != nil {
		// error handling here
		return err
	}

	return c.JSON(http.StatusOK, u)
}
