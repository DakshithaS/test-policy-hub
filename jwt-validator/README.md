# JWT Validator v1.0.0

Basic JWT validation policy.

## Features

- Simple secret-based validation

## Usage

```go
validator := jwtvalidator.NewJWTValidator("secret")
if validator.Validate(token) {
    // Valid
}
```