package interactor

import (
	"github.com/manakuro/golang-clean-architecture/domain/model"
	"github.com/manakuro/golang-clean-architecture/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r repository.UserRepository) UserInteractor {
	return &userInteractor{r}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
