package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
)

func TestCORSMiddleware(t *testing.T) {
	tests := []struct {
		name             string
		method           string
		origin           string
		corsConfig       config.CORSConfig
		expectNextCalled bool
		expectStatusCode int
		expectHeaders    map[string]string
	}{
		{
			name:   "Allowed origin exact match GET",
			method: "GET",
			origin: "https://example.com",
			corsConfig: config.CORSConfig{
				AllowedOrigins:   []string{"https://example.com", "https://other.com"},
				AllowedMethods:   []string{"GET", "POST"},
				AllowedHeaders:   []string{"Content-Type", "Authorization"},
				AllowCredentials: true,
			},
			expectNextCalled: true,
			expectStatusCode: 200,
			expectHeaders: map[string]string{
				"Access-Control-Allow-Origin":      "https://example.com",
				"Access-Control-Allow-Methods":     "GET,POST",
				"Access-Control-Allow-Headers":     "Content-Type,Authorization",
				"Access-Control-Allow-Credentials": "true",
			},
		},
		{
			name:   "Allowed origin wildcard GET",
			method: "GET",
			origin: "https://example.com",
			corsConfig: config.CORSConfig{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{"GET", "POST"},
				AllowedHeaders:   []string{"Content-Type"},
				AllowCredentials: false,
			},
			expectNextCalled: true,
			expectStatusCode: 200,
			expectHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET,POST",
				"Access-Control-Allow-Headers": "Content-Type",
			},
		},
		{
			name:   "Disallowed origin GET",
			method: "GET",
			origin: "https://bad.com",
			corsConfig: config.CORSConfig{
				AllowedOrigins:   []string{"https://example.com"},
				AllowedMethods:   []string{"GET"},
				AllowedHeaders:   []string{"Content-Type"},
				AllowCredentials: false,
			},
			expectNextCalled: true,
			expectStatusCode: 200,
			expectHeaders: map[string]string{
				"Access-Control-Allow-Methods": "GET",
				"Access-Control-Allow-Headers": "Content-Type",
			},
		},
		{
			name:   "Preflight OPTIONS allowed origin",
			method: "OPTIONS",
			origin: "https://example.com",
			corsConfig: config.CORSConfig{
				AllowedOrigins:   []string{"https://example.com"},
				AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
				AllowedHeaders:   []string{"Content-Type"},
				AllowCredentials: true,
			},
			expectNextCalled: false,
			expectStatusCode: 200,
			expectHeaders: map[string]string{
				"Access-Control-Allow-Origin":      "https://example.com",
				"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
				"Access-Control-Allow-Headers":     "Content-Type",
				"Access-Control-Allow-Credentials": "true",
			},
		},
		{
			name:   "Preflight OPTIONS disallowed origin",
			method: "OPTIONS",
			origin: "https://bad.com",
			corsConfig: config.CORSConfig{
				AllowedOrigins:   []string{"https://example.com"},
				AllowedMethods:   []string{"GET", "OPTIONS"},
				AllowedHeaders:   []string{"Content-Type"},
				AllowCredentials: false,
			},
			expectNextCalled: false,
			expectStatusCode: 200,
			expectHeaders: map[string]string{
				"Access-Control-Allow-Methods": "GET,OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type",
			},
		},
		{
			name:   "Credentials not set with wildcard",
			method: "GET",
			origin: "https://example.com",
			corsConfig: config.CORSConfig{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{"GET"},
				AllowedHeaders:   []string{"Content-Type"},
				AllowCredentials: true,
			},
			expectNextCalled: true,
			expectStatusCode: 200,
			expectHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET",
				"Access-Control-Allow-Headers": "Content-Type",
			},
		},
		{
			name:   "No origin header",
			method: "GET",
			origin: "",
			corsConfig: config.CORSConfig{
				AllowedOrigins:   []string{"https://example.com"},
				AllowedMethods:   []string{"GET"},
				AllowedHeaders:   []string{"Content-Type"},
				AllowCredentials: false,
			},
			expectNextCalled: true,
			expectStatusCode: 200,
			expectHeaders: map[string]string{
				"Access-Control-Allow-Methods": "GET",
				"Access-Control-Allow-Headers": "Content-Type",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextCalled := false
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
				w.WriteHeader(http.StatusOK)
			})

			middleware := CORSMiddleware(tt.corsConfig)
			handler := middleware(next)

			req := httptest.NewRequest(tt.method, "/test", nil)
			if tt.origin != "" {
				req.Header.Set("Origin", tt.origin)
			}
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)

			if nextCalled != tt.expectNextCalled {
				t.Errorf("expected next called: %v, got: %v", tt.expectNextCalled, nextCalled)
			}

			if w.Code != tt.expectStatusCode {
				t.Errorf("expected status code: %d, got: %d", tt.expectStatusCode, w.Code)
			}

			for key, expectedValue := range tt.expectHeaders {
				actualValue := w.Header().Get(key)
				if actualValue != expectedValue {
					t.Errorf("expected header %s: %s, got: %s", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestIsOriginAllowed(t *testing.T) {
	tests := []struct {
		name           string
		origin         string
		allowedOrigins []string
		expected       bool
	}{
		{
			name:           "exact match",
			origin:         "https://example.com",
			allowedOrigins: []string{"https://example.com", "https://other.com"},
			expected:       true,
		},
		{
			name:           "wildcard",
			origin:         "https://example.com",
			allowedOrigins: []string{"*"},
			expected:       true,
		},
		{
			name:           "no match",
			origin:         "https://bad.com",
			allowedOrigins: []string{"https://example.com"},
			expected:       false,
		},
		{
			name:           "empty allowed origins",
			origin:         "https://example.com",
			allowedOrigins: []string{},
			expected:       false,
		},
		{
			name:           "empty origin",
			origin:         "",
			allowedOrigins: []string{"https://example.com"},
			expected:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isOriginAllowed(tt.origin, tt.allowedOrigins)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
