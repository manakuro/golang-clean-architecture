package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	"github.com/manakuro/golang-clean-architecture/config"
	"github.com/manakuro/golang-clean-architecture/infrastructure/datastore"
	"github.com/manakuro/golang-clean-architecture/infrastructure/router"
	"github.com/manakuro/golang-clean-architecture/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
