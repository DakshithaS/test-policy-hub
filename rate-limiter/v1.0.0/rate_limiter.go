package ratelimiter

import "time"

// RateLimiter implements a simple rate limiter.
type RateLimiter struct {
	limit  int
	window time.Duration
	// Placeholder for internal state
}

// NewRateLimiter creates a new rate limiter.
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

// Allow checks if the request is allowed based on the rate limit.
func (rl *RateLimiter) Allow() bool {
	// Placeholder implementation: always allow
	return true
}
