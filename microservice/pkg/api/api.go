package api

import (
	"microservice/pkg/api/controllers/comment"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

func StartAPI(pgdb *pg.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	router.Route("/comments", func(router chi.Router) {
		router.Get("/", comment.GetComments())
	})

	router.Get("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("up and running"))
	})

	return router
}
