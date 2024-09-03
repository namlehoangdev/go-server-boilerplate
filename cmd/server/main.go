package main

import (
	"github.com/namlehoangdev/go-server-boilerplate/internal/server/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRouter()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
