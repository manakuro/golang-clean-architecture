package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/manakuro/golang-clean-architecture/infrastructure/api/handler"
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
