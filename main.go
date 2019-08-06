package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/manakuro/golang-clean-architecture/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/manakuro/golang-clean-architecture/domain/model"
)

func connectDB() *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3307",
		DBName:               "golang-clean-architecture",
		AllowNativePasswords: true,
		Params:               map[string]string{"parseTime": "true"},
	}
	fmt.Println(mySqlConfig.FormatDSN())

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func main() {
	config.ReadConfig()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", getUsers)
	})

	fmt.Println("Server listen at http://localhost:8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db := connectDB()

	defer db.Close()

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
