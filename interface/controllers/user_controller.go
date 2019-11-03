package controllers

import (
	"net/http"

	"github.com/manakuro/golang-clean-architecture/domain/model"
	"github.com/manakuro/golang-clean-architecture/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetUsers(c Context) error
}

func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userService.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
