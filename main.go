package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/manakuro/golang-clean-architecture/config"
	"github.com/manakuro/golang-clean-architecture/infrastructure/api/router"
	"github.com/manakuro/golang-clean-architecture/infrastructure/datastore"
	"github.com/manakuro/golang-clean-architecture/registry"
)

var db *gorm.DB

func main() {
	config.ReadConfig()

	db = datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	i := registry.NewInteractor(db)

	r := router.NewRouter(i.NewAppHandler())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := http.ListenAndServe(":"+config.C.Server.Address, r); err != nil {
		log.Fatalln(err)
	}
}
