package registry

import (
	"github.com/manakuro/golang-clean-architecture/interface/controllers"
	"github.com/manakuro/golang-clean-architecture/interface/presenters"
	ir "github.com/manakuro/golang-clean-architecture/interface/repository"
	"github.com/manakuro/golang-clean-architecture/usecase/interactor"
	"github.com/manakuro/golang-clean-architecture/usecase/presenter"
	ur "github.com/manakuro/golang-clean-architecture/usecase/repository"
)

func (r *registry) NewUserController() controllers.UserController {
	return controllers.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() presenter.UserPresenter {
	return presenters.NewUserPresenter()
}
