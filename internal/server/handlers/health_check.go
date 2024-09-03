package handlers

import (
	"encoding/json"
	"github.com/namlehoangdev/go-server-boilerplate/internal/models"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(models.HealthCheckResponse{Message: "Good"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
