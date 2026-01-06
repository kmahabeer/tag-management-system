package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
)

// TokenBucket represents a token bucket for rate limiting
type TokenBucket struct {
	tokens     float64   // Current number of tokens
	lastRefill time.Time // Last time tokens were refilled
	capacity   int       // Maximum tokens (requests per minute)
}

// RateLimiter manages rate limiting for multiple IPs
type RateLimiter struct {
	buckets           map[string]*TokenBucket // IP -> TokenBucket
	mu                sync.RWMutex            // Thread safety
	requestsPerMinute int                     // Configurable limit
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(requestsPerMinute int) *RateLimiter {
	return &RateLimiter{
		buckets:           make(map[string]*TokenBucket),
		requestsPerMinute: requestsPerMinute,
	}
}

// Allow checks if a request from the given IP is allowed
func (rl *RateLimiter) Allow(ip string) (bool, time.Duration) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	bucket, exists := rl.buckets[ip]
	if !exists {
		// Create new bucket
		bucket = &TokenBucket{
			tokens:     float64(rl.requestsPerMinute),
			lastRefill: now,
			capacity:   rl.requestsPerMinute,
		}
		rl.buckets[ip] = bucket
	} else {
		// Refill tokens based on elapsed time
		elapsed := now.Sub(bucket.lastRefill)
		refillAmount := elapsed.Minutes() * float64(rl.requestsPerMinute)

		bucket.tokens += refillAmount
		if bucket.tokens > float64(bucket.capacity) {
			bucket.tokens = float64(bucket.capacity)
		}
	}

	bucket.lastRefill = now

	if bucket.tokens >= 1 {
		bucket.tokens -= 1
		return true, 0
	}

	// Calculate wait time for next token
	waitTime := time.Duration((1 - bucket.tokens) / float64(rl.requestsPerMinute) * float64(time.Minute))
	return false, waitTime
}

// RateLimitMiddleware creates a rate limiting middleware
func RateLimitMiddleware(rateLimitConfig config.RateLimitConfig) func(http.Handler) http.Handler {
	limiter := NewRateLimiter(rateLimitConfig.RequestsPerMinute)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := getClientIP(r)

			allowed, waitTime := limiter.Allow(ip)
			if !allowed {
				// Return 429 Too Many Requests
				w.Header().Set("Retry-After", waitTime.Round(time.Second).String())
				w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", rateLimitConfig.RequestsPerMinute))
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(waitTime).Unix()))

				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// getClientIP extracts the client IP address from the request
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header first (for proxies)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// Take the first IP in the chain
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// Fall back to RemoteAddr
	return r.RemoteAddr
}
