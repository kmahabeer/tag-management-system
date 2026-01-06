package middleware

import (
	"net/http"
	"strings"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
)

func CORSMiddleware(corsConfig config.CORSConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			isOriginAllowed := isOriginAllowed(origin, corsConfig.AllowedOrigins)

			if r.Method == "OPTIONS" {
				setCORSHeaders(w, corsConfig, origin, isOriginAllowed)
				w.WriteHeader(http.StatusOK)
				return
			}

			setCORSHeaders(w, corsConfig, origin, isOriginAllowed)
			next.ServeHTTP(w, r)
		})
	}
}

func isOriginAllowed(origin string, allowedOrigins []string) bool {
	if len(allowedOrigins) == 0 {
		return false
	}
	for _, allowed := range allowedOrigins {
		if allowed == "*" || allowed == origin {
			return true
		}
	}
	return false
}

func setCORSHeaders(w http.ResponseWriter, corsConfig config.CORSConfig, origin string, isOriginAllowed bool) {
	if isOriginAllowed {
		if len(corsConfig.AllowedOrigins) == 1 && corsConfig.AllowedOrigins[0] == "*" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
	}

	w.Header().Set("Access-Control-Allow-Methods", strings.Join(corsConfig.AllowedMethods, ","))
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(corsConfig.AllowedHeaders, ","))

	if corsConfig.AllowCredentials && isOriginAllowed && (len(corsConfig.AllowedOrigins) != 1 || corsConfig.AllowedOrigins[0] != "*") {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}
