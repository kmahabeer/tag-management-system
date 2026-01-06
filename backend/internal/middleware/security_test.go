package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
)

func TestSecurityHeadersMiddleware(t *testing.T) {
	tests := []struct {
		name                string
		securityConfig      config.SecurityConfig
		setXPoweredByBefore bool
		expectHeaders       map[string]string
		expectXPoweredBy    bool // true if should be present, false if removed
	}{
		{
			name: "All headers set with values",
			securityConfig: config.SecurityConfig{
				ContentSecurityPolicy:   "default-src 'self'",
				StrictTransportSecurity: "max-age=31536000",
				XContentTypeOptions:     "nosniff",
				XFrameOptions:           "DENY",
				XXSSProtection:          "1; mode=block",
				ReferrerPolicy:          "no-referrer",
				PermissionsPolicy:       "geolocation=(), microphone=()",
				CacheControl:            "no-cache",
				ServerHeader:            "MyServer/1.0",
				RemoveXPoweredBy:        false,
			},
			setXPoweredByBefore: false,
			expectHeaders: map[string]string{
				"Content-Security-Policy":   "default-src 'self'",
				"Strict-Transport-Security": "max-age=31536000",
				"X-Content-Type-Options":    "nosniff",
				"X-Frame-Options":           "DENY",
				"X-XSS-Protection":          "1; mode=block",
				"Referrer-Policy":           "no-referrer",
				"Permissions-Policy":        "geolocation=(), microphone=()",
				"Cache-Control":             "no-cache",
				"Server":                    "MyServer/1.0",
			},
			expectXPoweredBy: false, // not set
		},
		{
			name: "All headers empty",
			securityConfig: config.SecurityConfig{
				ContentSecurityPolicy:   "",
				StrictTransportSecurity: "",
				XContentTypeOptions:     "",
				XFrameOptions:           "",
				XXSSProtection:          "",
				ReferrerPolicy:          "",
				PermissionsPolicy:       "",
				CacheControl:            "",
				ServerHeader:            "",
				RemoveXPoweredBy:        false,
			},
			setXPoweredByBefore: false,
			expectHeaders:       map[string]string{},
			expectXPoweredBy:    false,
		},
		{
			name: "Server header set",
			securityConfig: config.SecurityConfig{
				ServerHeader: "TestServer",
			},
			setXPoweredByBefore: false,
			expectHeaders: map[string]string{
				"Server": "TestServer",
			},
			expectXPoweredBy: false,
		},
		{
			name: "Remove X-Powered-By header",
			securityConfig: config.SecurityConfig{
				RemoveXPoweredBy: true,
			},
			setXPoweredByBefore: true,
			expectHeaders:       map[string]string{},
			expectXPoweredBy:    false,
		},
		{
			name: "Default config behavior",
			securityConfig: config.SecurityConfig{
				ContentSecurityPolicy:   "default-src 'self'",
				StrictTransportSecurity: "max-age=31536000; includeSubDomains",
				XContentTypeOptions:     "nosniff",
				XFrameOptions:           "DENY",
				XXSSProtection:          "1; mode=block",
				ReferrerPolicy:          "strict-origin-when-cross-origin",
				PermissionsPolicy:       "",
				CacheControl:            "no-cache, no-store, must-revalidate",
				ServerHeader:            "",
				RemoveXPoweredBy:        true,
			},
			setXPoweredByBefore: true,
			expectHeaders: map[string]string{
				"Content-Security-Policy":   "default-src 'self'",
				"Strict-Transport-Security": "max-age=31536000; includeSubDomains",
				"X-Content-Type-Options":    "nosniff",
				"X-Frame-Options":           "DENY",
				"X-XSS-Protection":          "1; mode=block",
				"Referrer-Policy":           "strict-origin-when-cross-origin",
				"Cache-Control":             "no-cache, no-store, must-revalidate",
			},
			expectXPoweredBy: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextCalled := false
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
				w.WriteHeader(http.StatusOK)
			})

			middleware := SecurityHeadersMiddleware(tt.securityConfig)
			handler := middleware(next)

			req := httptest.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()

			if tt.setXPoweredByBefore {
				w.Header().Set("X-Powered-By", "Express")
			}

			handler.ServeHTTP(w, req)

			if !nextCalled {
				t.Errorf("expected next to be called")
			}

			if w.Code != http.StatusOK {
				t.Errorf("expected status code 200, got %d", w.Code)
			}

			for key, expectedValue := range tt.expectHeaders {
				actualValue := w.Header().Get(key)
				if actualValue != expectedValue {
					t.Errorf("expected header %s: %s, got: %s", key, expectedValue, actualValue)
				}
			}

			xPoweredBy := w.Header().Get("X-Powered-By")
			if tt.expectXPoweredBy && xPoweredBy == "" {
				t.Errorf("expected X-Powered-By header to be present")
			}
			if !tt.expectXPoweredBy && xPoweredBy != "" {
				t.Errorf("expected X-Powered-By header to be removed, but got: %s", xPoweredBy)
			}
		})
	}
}
