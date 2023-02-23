package repository

import (
	"golang-clean-architecture/pkg/usecase/repository"
	"log"

	"github.com/jinzhu/gorm"
)

type dbRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) repository.DBRepository {
	return &dbRepository{db}
}

func (r *dbRepository) Transaction(txFunc func(interface{}) (interface{}, error)) (data interface{}, err error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			log.Print("recover")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Print("rollback")
			tx.Rollback()
			panic("error")
		} else {
			err = tx.Commit().Error
		}
	}()

	data, err = txFunc(tx)
	return data, err
}
