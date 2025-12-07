package jwtvalidator

import "strings"

// JWTValidator validates JWT tokens.
type JWTValidator struct {
	secret string
}

// NewJWTValidator creates a new validator.
func NewJWTValidator(secret string) *JWTValidator {
	return &JWTValidator{secret: secret}
}

// Validate checks if the token is valid.
func (jv *JWTValidator) Validate(token string) bool {
	// Placeholder: check if token contains secret
	return strings.Contains(token, jv.secret)
}