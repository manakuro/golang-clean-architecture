package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/manakuro/golang-clean-architecture/infrastructure/api/handler"
)

func NewRouter(handler handler.AppHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
	})

	return r
}
