package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/manakuro/golang-clean-architecture/interface/controllers"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controllers.AppController {
	return r.NewUserController()
}
