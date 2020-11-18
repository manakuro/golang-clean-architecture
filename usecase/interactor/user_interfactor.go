package interactor

import (
	"errors"
	"golang-clean-architecture/domain/model"
	"golang-clean-architecture/usecase/presenter"
	"golang-clean-architecture/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
	DBRepository   repository.DBRepository
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter, d repository.DBRepository) UserInteractor {
	return &userInteractor{r, p, d}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) Create(u *model.User) (*model.User, error) {
	data, err := us.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := us.UserRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	user, ok := data.(*model.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return user, nil
}
