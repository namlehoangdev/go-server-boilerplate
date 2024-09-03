package routes

import (
	"github.com/gorilla/mux"
	"github.com/namlehoangdev/go-server-boilerplate/internal/server/handlers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	apiV1 := router.PathPrefix("/v1").Subrouter()
	apiV1.HandleFunc("/health-check", handlers.HealthCheckHandler)
	return router
}
