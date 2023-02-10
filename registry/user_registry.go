package registry

import (
	"golang-clean-architecture/interface/controller"
	ip "golang-clean-architecture/interface/presenter"
	ir "golang-clean-architecture/interface/repository"
	"golang-clean-architecture/usecase/interactor"
)

func (r *registry) NewUserController() controller.UserController {
	userInteractor := interactor.NewUserInteractor(
		ir.NewUserRepository(r.db),
		ip.NewUserPresenter(),
		ir.NewDBRepository(r.db),
	)

	return controller.NewUserController(userInteractor)
}
