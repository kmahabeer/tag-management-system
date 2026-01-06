package middleware

import (
	"log/slog"
	"net/http"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
)

func SecurityHeadersMiddleware(securityConfig config.SecurityConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headersSet := []string{}

			if securityConfig.ContentSecurityPolicy != "" {
				w.Header().Set("Content-Security-Policy", securityConfig.ContentSecurityPolicy)
				headersSet = append(headersSet, "Content-Security-Policy")
			}
			if securityConfig.StrictTransportSecurity != "" {
				w.Header().Set("Strict-Transport-Security", securityConfig.StrictTransportSecurity)
				headersSet = append(headersSet, "Strict-Transport-Security")
			}
			if securityConfig.XContentTypeOptions != "" {
				w.Header().Set("X-Content-Type-Options", securityConfig.XContentTypeOptions)
				headersSet = append(headersSet, "X-Content-Type-Options")
			}
			if securityConfig.XFrameOptions != "" {
				w.Header().Set("X-Frame-Options", securityConfig.XFrameOptions)
				headersSet = append(headersSet, "X-Frame-Options")
			}
			if securityConfig.XXSSProtection != "" {
				w.Header().Set("X-XSS-Protection", securityConfig.XXSSProtection)
				headersSet = append(headersSet, "X-XSS-Protection")
			}
			if securityConfig.ReferrerPolicy != "" {
				w.Header().Set("Referrer-Policy", securityConfig.ReferrerPolicy)
				headersSet = append(headersSet, "Referrer-Policy")
			}
			if securityConfig.PermissionsPolicy != "" {
				w.Header().Set("Permissions-Policy", securityConfig.PermissionsPolicy)
				headersSet = append(headersSet, "Permissions-Policy")
			}
			if securityConfig.CacheControl != "" {
				w.Header().Set("Cache-Control", securityConfig.CacheControl)
				headersSet = append(headersSet, "Cache-Control")
			}
			if securityConfig.ServerHeader != "" {
				w.Header().Set("Server", securityConfig.ServerHeader)
				headersSet = append(headersSet, "Server")
			}
			if securityConfig.RemoveXPoweredBy {
				w.Header().Del("X-Powered-By")
				headersSet = append(headersSet, "X-Powered-By-removed")
			}

			if len(headersSet) > 0 {
				requestID := getRequestID(r.Context())
				slog.DebugContext(r.Context(), "Security headers applied",
					"request_id", requestID,
					"headers", headersSet,
				)
			}

			next.ServeHTTP(w, r)
		})
	}
}
