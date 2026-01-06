package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/kmahabeer/tag-management-system/backend/internal/middleware"
)

type Validatable interface {
	Validate() error
}

func WriteJSON(w http.ResponseWriter, r *http.Request, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.ErrorContext(r.Context(), "Failed to encode JSON response",
			"operation", "encode_response",
			"status", status,
			"error", err,
		)
		return err
	}
	return nil
}

func DecodeAndValidate(r *http.Request, v Validatable) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		slog.ErrorContext(r.Context(), "Failed to decode request body",
			"operation", "decode_request",
			"error", err,
		)
		return err
	}
	if err := v.Validate(); err != nil {
		slog.WarnContext(r.Context(), "Validation failed",
			"operation", "validate_request",
			"error", err,
		)
		return err
	}
	return nil
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	middleware.WriteError(w, http.StatusNotImplemented, "Not implemented", nil)
}
