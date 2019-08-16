package repository

import "github.com/manakuro/golang-clean-architecture/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
