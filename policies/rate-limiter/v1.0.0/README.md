# Rate Limiter v1.0.0

This is the initial version of the rate limiter policy.

## Features

- Basic rate limiting with configurable limit and window
- Simple allow/deny mechanism

## Usage

```go
rl := ratelimiter.NewRateLimiter(100, time.Minute)
if rl.Allow() {
    // Process request
}
```