package controller

import (
	"net/http"

	"github.com/manakuro/golang-clean-architecture/api"

	"github.com/manakuro/golang-clean-architecture/domain/model"
	"github.com/manakuro/golang-clean-architecture/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
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

	// mapping []*model.User to []*api.User here
	us := mapUser(u)

	return c.JSON(http.StatusOK, us)
}

func mapUser(us []*model.User) []*api.User {
	result := make([]*api.User, len(us))

	for i, v := range us {
		au := api.User{
			ID:   v.ID,
			Name: v.Name,
			Age:  v.Age,
		}
		result[i] = &au
	}

	return result
}
