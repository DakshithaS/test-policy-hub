package ratelimiter

import "time"

// RateLimiter implements a simple rate limiter with burst allowance.
type RateLimiter struct {
	limit  int
	window time.Duration
	burst  int
	// Placeholder for internal state
}

// NewRateLimiter creates a new rate limiter.
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
		burst:  limit, // Default burst to limit
	}
}

// NewRateLimiterWithBurst creates a rate limiter with burst capacity.
func NewRateLimiterWithBurst(limit int, window time.Duration, burst int) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
		burst:  burst,
	}
}

// Allow checks if the request is allowed based on the rate limit.
func (rl *RateLimiter) Allow() bool {
	// Placeholder implementation: always allow
	return true
}
