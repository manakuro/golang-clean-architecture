package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
	dBRepository   repository.DBRepository
}

type User interface {
	List(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}

func NewUserUsecase(r repository.UserRepository, d repository.DBRepository) User {
	return &userUsecase{r, d}
}

func (uu *userUsecase) List(u []*model.User) ([]*model.User, error) {
	u, err := uu.userRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) Create(u *model.User) (*model.User, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.userRepository.Create(u)

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
