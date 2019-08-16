package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/manakuro/golang-clean-architecture/registry"

	"github.com/manakuro/golang-clean-architecture/infrastructure/api/router"

	"github.com/jinzhu/gorm"

	"github.com/manakuro/golang-clean-architecture/infrastructure/datastore"

	"github.com/manakuro/golang-clean-architecture/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	config.ReadConfig()

	db = datastore.NewDB()
	defer db.Close()

	i := registry.NewInteractor(db)

	r := router.NewRouter(i.NewAppHandler())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := http.ListenAndServe(":"+config.C.Server.Address, r); err != nil {
		log.Fatalln(err)
	}
}
