# Rate Limiter v1.1.0

This version adds burst support to the rate limiter.

## Features

- Configurable burst capacity
- Backward compatible with v1.0.0

## Usage

```go
rl := ratelimiter.NewRateLimiterWithBurst(100, time.Minute, 200)
if rl.Allow() {
    // Process request
}
```