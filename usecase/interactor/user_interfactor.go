package interactor

import (
	"golang-clean-architecture/domain/model"
	"golang-clean-architecture/usecase/presenter"
	"golang-clean-architecture/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) Create(u *model.User) (*model.User, error) {
	u, err := us.UserRepository.Create(u)

	return u, err
}
