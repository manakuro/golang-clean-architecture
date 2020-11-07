package presenter

import (
	"golang-clean-architecture/domain/model"
	"golang-clean-architecture/usecase/presenter"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
