package registry

import (
	"golang-clean-architecture/interface/controller"
	ip "golang-clean-architecture/interface/presenter"
	ir "golang-clean-architecture/interface/repository"
	"golang-clean-architecture/usecase/interactor"
	up "golang-clean-architecture/usecase/presenter"
	ur "golang-clean-architecture/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
