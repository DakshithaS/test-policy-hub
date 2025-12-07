# JWT Validator v2.1.0

Improved JWT validation with proper HMAC-SHA256 verification.

## Features

- HMAC-SHA256 signature verification
- Context support

## Usage

```go
validator := jwtvalidator.New("secret")
valid, err := validator.ValidateWithContext(ctx, token)
```