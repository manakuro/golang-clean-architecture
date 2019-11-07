package registry

import (
	"github.com/manakuro/golang-clean-architecture/interface/controller"
	ip "github.com/manakuro/golang-clean-architecture/interface/presenter"
	ir "github.com/manakuro/golang-clean-architecture/interface/repository"
	"github.com/manakuro/golang-clean-architecture/usecase/interactor"
	up "github.com/manakuro/golang-clean-architecture/usecase/presenter"
	ur "github.com/manakuro/golang-clean-architecture/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
