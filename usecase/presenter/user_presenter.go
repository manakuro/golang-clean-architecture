package presenter

import "github.com/manakuro/golang-clean-architecture/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
