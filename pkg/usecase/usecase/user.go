package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type userUsecase struct {
	UserRepository repository.UserRepository
	DBRepository   repository.DBRepository
}

type User interface {
	Get(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}

func NewUserUsecase(r repository.UserRepository, d repository.DBRepository) User {
	return &userUsecase{r, d}
}

func (uu *userUsecase) Get(u []*model.User) ([]*model.User, error) {
	u, err := uu.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) Create(u *model.User) (*model.User, error) {
	data, err := uu.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.UserRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	user, ok := data.(*model.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
