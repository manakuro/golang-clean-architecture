package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/usecase"
)

func (r *registry) NewUserController() controller.User {
	userInteractor := usecase.NewUserUsecase(
		repository.NewUserRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewUserController(userInteractor)
}
