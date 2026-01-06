package handlers

import (
	"log/slog"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	slog.InfoContext(r.Context(), "Health check performed successfully",
		"operation", "health_check",
	)

	WriteJSON(w, http.StatusOK, response)
}
