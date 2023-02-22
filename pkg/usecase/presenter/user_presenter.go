package presenter

import "golang-clean-architecture/pkg/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
