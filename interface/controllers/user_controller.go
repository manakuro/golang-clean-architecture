package controllers

import (
	"github.com/manakuro/golang-clean-architecture/domain/model"
	"github.com/manakuro/golang-clean-architecture/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetUsers() ([]*model.User, error)
}

func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers() ([]*model.User, error) {
	var u []*model.User

	us, err := uc.userService.Get(u)
	if err != nil {
		return nil, err
	}

	return us, nil
}
