package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
)

func TestNewRateLimiter(t *testing.T) {
	limiter := NewRateLimiter(60)
	if limiter.requestsPerMinute != 60 {
		t.Errorf("expected requestsPerMinute to be 60, got %d", limiter.requestsPerMinute)
	}
	if len(limiter.buckets) != 0 {
		t.Errorf("expected buckets to be empty, got %d", len(limiter.buckets))
	}
}

func TestRateLimiter_Allow(t *testing.T) {
	limiter := NewRateLimiter(1) // 1 request per minute

	// First request should be allowed
	allowed, waitTime := limiter.Allow("192.168.1.1")
	if !allowed {
		t.Error("expected first request to be allowed")
	}
	if waitTime != 0 {
		t.Errorf("expected waitTime to be 0, got %v", waitTime)
	}

	// Second request should be denied
	allowed, waitTime = limiter.Allow("192.168.1.1")
	if allowed {
		t.Error("expected second request to be denied")
	}
	if waitTime <= 0 {
		t.Errorf("expected waitTime to be positive, got %v", waitTime)
	}
}

func TestRateLimiter_Refill(t *testing.T) {
	limiter := NewRateLimiter(60) // 60 requests per minute = 1 per second

	// Use up all tokens
	for i := 0; i < 60; i++ {
		allowed, _ := limiter.Allow("192.168.1.1")
		if !allowed {
			t.Errorf("expected request %d to be allowed", i+1)
		}
	}

	// Next request should be denied
	allowed, waitTime := limiter.Allow("192.168.1.1")
	if allowed {
		t.Error("expected request to be denied after using all tokens")
	}

	// Simulate time passing (1 second = 1 token)
	bucket := limiter.buckets["192.168.1.1"]
	bucket.lastRefill = bucket.lastRefill.Add(-time.Second)

	// Now request should be allowed again
	allowed, waitTime = limiter.Allow("192.168.1.1")
	if !allowed {
		t.Error("expected request to be allowed after refill")
	}
	if waitTime != 0 {
		t.Errorf("expected waitTime to be 0, got %v", waitTime)
	}
}

func TestRateLimitMiddleware(t *testing.T) {
	tests := []struct {
		name          string
		requests      int
		expectStatus  int
		expectHeaders bool
	}{
		{
			name:          "first request allowed",
			requests:      1,
			expectStatus:  200,
			expectHeaders: false,
		},
		{
			name:          "rate limited",
			requests:      2, // Exceed 1 per minute
			expectStatus:  429,
			expectHeaders: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestsPerMinute := 60
			if tt.name == "rate limited" {
				requestsPerMinute = 1
			}
			cfg := config.RateLimitConfig{RequestsPerMinute: requestsPerMinute}
			middleware := RateLimitMiddleware(cfg)

			nextCalled := false
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
				w.WriteHeader(http.StatusOK)
			})

			handler := middleware(next)

			var lastRecorder *httptest.ResponseRecorder
			for i := 0; i < tt.requests; i++ {
				req := httptest.NewRequest("GET", "/test", nil)
				req.RemoteAddr = "192.168.1.1:12345"
				w := httptest.NewRecorder()

				handler.ServeHTTP(w, req)
				lastRecorder = w
			}

			if tt.expectStatus == 200 && !nextCalled {
				t.Error("expected next handler to be called")
			}

			if lastRecorder.Code != tt.expectStatus {
				t.Errorf("expected status %d, got %d", tt.expectStatus, lastRecorder.Code)
			}

			if tt.expectHeaders {
				if lastRecorder.Header().Get("Retry-After") == "" {
					t.Error("expected Retry-After header")
				}
				if lastRecorder.Header().Get("X-RateLimit-Limit") == "" {
					t.Error("expected X-RateLimit-Limit header")
				}
			}
		})
	}
}

func TestGetClientIP(t *testing.T) {
	tests := []struct {
		name       string
		headers    map[string]string
		remoteAddr string
		expected   string
	}{
		{
			name:       "X-Forwarded-For present",
			headers:    map[string]string{"X-Forwarded-For": "203.0.113.1, 198.51.100.1"},
			remoteAddr: "127.0.0.1:12345",
			expected:   "203.0.113.1",
		},
		{
			name:       "X-Forwarded-For single IP",
			headers:    map[string]string{"X-Forwarded-For": "203.0.113.1"},
			remoteAddr: "127.0.0.1:12345",
			expected:   "203.0.113.1",
		},
		{
			name:       "no X-Forwarded-For",
			headers:    map[string]string{},
			remoteAddr: "127.0.0.1:12345",
			expected:   "127.0.0.1:12345",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test", nil)
			req.RemoteAddr = tt.remoteAddr
			for k, v := range tt.headers {
				req.Header.Set(k, v)
			}

			ip := getClientIP(req)
			if ip != tt.expected {
				t.Errorf("expected IP %s, got %s", tt.expected, ip)
			}
		})
	}
}
