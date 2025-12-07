package ratelimiter

import (
	"context"
	"time"
)

// RateLimiter implements a rate limiter with context support.
type RateLimiter struct {
	limit  int
	window time.Duration
	burst  int
	// Placeholder for internal state
}

// New creates a new rate limiter.
func New(limit int, window time.Duration, burst int) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
		burst:  burst,
	}
}

// Allow checks if the request is allowed.
func (rl *RateLimiter) Allow() bool {
	// Placeholder implementation
	return true
}

// AllowWithContext checks with context cancellation.
func (rl *RateLimiter) AllowWithContext(ctx context.Context) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
		return rl.Allow(), nil
	}
}