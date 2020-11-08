package repository

import "golang-clean-architecture/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}
