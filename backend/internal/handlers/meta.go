package handlers

import (
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	WriteJSON(w, r, http.StatusOK, response)
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}
