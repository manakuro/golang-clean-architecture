package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/manakuro/golang-clean-architecture/registry"

	"github.com/manakuro/golang-clean-architecture/infrastructure/api/router"

	"github.com/jinzhu/gorm"

	"github.com/manakuro/golang-clean-architecture/infrastructure/datastore"

	"github.com/manakuro/golang-clean-architecture/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/manakuro/golang-clean-architecture/domain/model"
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	db.Debug().Preload("CreditCards").Find(&users)

	fmt.Printf("%+v", users)

	respond(w, http.StatusOK, users)
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
