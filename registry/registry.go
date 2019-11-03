package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/manakuro/golang-clean-architecture/interface/controllers"
)

type interactor struct {
	db *gorm.DB
}

type Interactor interface {
	NewAppController() controllers.AppController
}

func NewInteractor(db *gorm.DB) Interactor {
	return &interactor{db}
}

func (i *interactor) NewAppController() controllers.AppController {
	return i.NewUserController()
}
