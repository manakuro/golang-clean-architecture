package presenter

import "github.com/manakuro/golang-clean-architecture/domain/model"

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = u.Name + "nyan"
	}
	return us
}
