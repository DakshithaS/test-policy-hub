# Rate Limiter v2.0.0

This version introduces context support and breaking API changes.

## Breaking Changes

- Constructor renamed to `New`
- Burst is now required
- Added `AllowWithContext` method

## Usage

```go
rl := ratelimiter.New(100, time.Minute, 200)
allowed, err := rl.AllowWithContext(ctx)
```