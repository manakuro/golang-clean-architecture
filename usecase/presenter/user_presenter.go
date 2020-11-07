package presenter

import "golang-clean-architecture/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
