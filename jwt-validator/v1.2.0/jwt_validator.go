package jwtvalidator

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// JWTValidator validates JWT tokens with HMAC.
type JWTValidator struct {
	secret []byte
}

// New creates a new validator.
func New(secret string) *JWTValidator {
	return &JWTValidator{secret: []byte(secret)}
}

// Validate checks the token.
func (jv *JWTValidator) Validate(token string) bool {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}
	headerPayload := parts[0] + "." + parts[1]
	signature, _ := base64.RawURLEncoding.DecodeString(parts[2])
	expected := hmac.New(sha256.New, jv.secret).Sum([]byte(headerPayload))
	return hmac.Equal(signature, expected)
}

// ValidateWithContext validates with context.
func (jv *JWTValidator) ValidateWithContext(ctx context.Context, token string) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
		return jv.Validate(token), nil
	}
}

// ValidateWithClaims validates and returns claims (new feature in v1.1.0).
func (jv *JWTValidator) ValidateWithClaims(token string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}
	if !jv.Validate(token) {
		return nil, errors.New("token validation failed")
	}
	return claims, nil
}
