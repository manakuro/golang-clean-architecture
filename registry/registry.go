package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/manakuro/golang-clean-architecture/infrastructure/api/handler"
	"github.com/manakuro/golang-clean-architecture/interface/controllers"
	"github.com/manakuro/golang-clean-architecture/interface/presenters"
	ir "github.com/manakuro/golang-clean-architecture/interface/repository"
	"github.com/manakuro/golang-clean-architecture/usecase/presenter"
	"github.com/manakuro/golang-clean-architecture/usecase/repository"
	"github.com/manakuro/golang-clean-architecture/usecase/service"
)

type interactor struct {
	db *gorm.DB
}

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(db *gorm.DB) Interactor {
	return &interactor{db}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return i.NewUserHandler()
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserController())
}

func (i *interactor) NewUserController() controllers.UserController {
	return controllers.NewUserController(i.NewUserService())
}

func (i *interactor) NewUserService() service.UserService {
	return service.NewUserService(i.NewUserRepository(), i.NewUserPresenter())
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return ir.NewUserRepository(i.db)
}

func (i *interactor) NewUserPresenter() presenter.UserPresenter {
	return presenters.NewUserPresenter()
}
