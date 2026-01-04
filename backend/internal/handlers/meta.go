package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kmahabeer/tag-management-system/backend/internal/middleware"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		middleware.WriteError(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}
}
