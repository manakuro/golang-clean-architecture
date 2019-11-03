package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

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

	e := echo.New()
	e = router.NewRouter(e, i.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
