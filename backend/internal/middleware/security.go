package middleware

import (
	"net/http"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
)

func SecurityHeadersMiddleware(securityConfig config.SecurityConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if securityConfig.ContentSecurityPolicy != "" {
				w.Header().Set("Content-Security-Policy", securityConfig.ContentSecurityPolicy)
			}
			if securityConfig.StrictTransportSecurity != "" {
				w.Header().Set("Strict-Transport-Security", securityConfig.StrictTransportSecurity)
			}
			if securityConfig.XContentTypeOptions != "" {
				w.Header().Set("X-Content-Type-Options", securityConfig.XContentTypeOptions)
			}
			if securityConfig.XFrameOptions != "" {
				w.Header().Set("X-Frame-Options", securityConfig.XFrameOptions)
			}
			if securityConfig.XXSSProtection != "" {
				w.Header().Set("X-XSS-Protection", securityConfig.XXSSProtection)
			}
			if securityConfig.ReferrerPolicy != "" {
				w.Header().Set("Referrer-Policy", securityConfig.ReferrerPolicy)
			}
			if securityConfig.PermissionsPolicy != "" {
				w.Header().Set("Permissions-Policy", securityConfig.PermissionsPolicy)
			}
			if securityConfig.CacheControl != "" {
				w.Header().Set("Cache-Control", securityConfig.CacheControl)
			}
			if securityConfig.ServerHeader != "" {
				w.Header().Set("Server", securityConfig.ServerHeader)
			}
			if securityConfig.RemoveXPoweredBy {
				w.Header().Del("X-Powered-By")
			}
			next.ServeHTTP(w, r)
		})
	}
}
