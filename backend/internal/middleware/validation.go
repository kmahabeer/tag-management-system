package middleware

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func isValidUUID(s string) bool {
	if len(s) != 36 {
		return false
	}
	for i := 0; i < len(s); i++ {
		b := s[i]
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if b != '-' {
				return false
			}
		} else {
			if !((b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')) {
				return false
			}
		}
	}
	return true
}

func ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		if ctx == nil {
			next.ServeHTTP(w, r)
			return
		}

		for _, param := range ctx.URLParams.Keys {
			if strings.HasSuffix(param, "_id") || param == "id" {
				value := chi.URLParam(r, param)
				if value != "" {
					if !isValidUUID(value) {
						WriteError(w, http.StatusBadRequest, "Invalid UUID format for parameter "+param, nil)
						return
					}
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
