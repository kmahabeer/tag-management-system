package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kmahabeer/tag-management-system/backend/internal/middleware"
)

type Validatable interface {
	Validate() error
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func DecodeAndValidate(r *http.Request, v Validatable) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.Validate()
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	middleware.WriteError(w, http.StatusNotImplemented, "Not implemented", nil)
}
