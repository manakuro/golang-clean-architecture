package registry

import (
	"golang-clean-architecture/pkg/interface/controller"
	ip "golang-clean-architecture/pkg/interface/presenter"
	ir "golang-clean-architecture/pkg/interface/repository"
	"golang-clean-architecture/pkg/usecase/interactor"
)

func (r *registry) NewUserController() controller.UserController {
	userInteractor := interactor.NewUserInteractor(
		ir.NewUserRepository(r.db),
		ip.NewUserPresenter(),
		ir.NewDBRepository(r.db),
	)

	return controller.NewUserController(userInteractor)
}
