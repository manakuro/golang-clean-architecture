package handler

import (
	"net/http"

	"github.com/manakuro/golang-clean-architecture/interface/controllers"
)

type userHandler struct {
	userController controllers.UserController
}

type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(uc controllers.UserController) UserHandler {
	return &userHandler{userController: uc}
}

func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	u, _ := uh.userController.GetUsers()

	ResponseHandler(w, http.StatusOK, u)
}
