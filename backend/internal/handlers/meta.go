package handlers

import (
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	WriteJSON(w, http.StatusOK, response)
}
