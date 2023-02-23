package controller

import (
	"golang-clean-architecture/pkg/usecase/usecase"
	"net/http"

	"golang-clean-architecture/pkg/domain/model"
)

type userController struct {
	userUsecase usecase.User
}

type User interface {
	GetUsers(c Context) error
	CreateUser(c Context) error
}

func NewUserController(us usecase.User) User {
	return &userController{us}
}

func (uc *userController) GetUsers(ctx Context) error {
	var u []*model.User

	u, err := uc.userUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(ctx Context) error {
	var params model.User

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := uc.userUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
