package firewall

import (
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// RateLimiter struct to track limits per IP
type RateLimiter struct {
	clients map[string]*rate.Limiter
	mu      sync.Mutex
	r       rate.Limit
	b       int
}

// NewRateLimiter initializes rate limiting (e.g., 5 requests per second, burst of 10)
func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		clients: make(map[string]*rate.Limiter),
		r:       r,
		b:       b,
	}
}

// GetLimiter returns the rate limiter for an IP
func (rl *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.clients[ip]
	if !exists {
		limiter = rate.NewLimiter(rl.r, rl.b)
		rl.clients[ip] = limiter
	}

	return limiter
}

// Middleware to enforce rate limiting
func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		limiter := rl.GetLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, "429 - Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
