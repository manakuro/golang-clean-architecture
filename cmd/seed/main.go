package main

import (
	"golang-clean-architecture/pkg/config"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/infrastructure/datastore"
	"log"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	u := &model.User{
		ID:        1,
		Name:      "Tom",
		Age:       "20",
		CreatedAt: nil,
		UpdatedAt: nil,
		DeletedAt: nil,
	}
	if err := db.Create(u).Error; err != nil {
		log.Fatalf("failed to seed user data: %v", err)
	}

}
