package controller

import (
	"errors"
	"net/http"

	"golang-clean-architecture/domain/model"
	"golang-clean-architecture/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
	CreateUser(c Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(c Context) error {
	var params model.User

	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	u, err := uc.userInteractor.Create(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
