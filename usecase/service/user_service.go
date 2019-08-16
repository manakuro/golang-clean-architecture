package service

import (
	"github.com/manakuro/golang-clean-architecture/domain/model"
	"github.com/manakuro/golang-clean-architecture/usecase/presenter"
	"github.com/manakuro/golang-clean-architecture/usecase/repository"
)

type userService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserService interface {
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserService(r repository.UserRepository, p presenter.UserPresenter) UserService {
	return &userService{r, p}
}

func (us *userService) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}
