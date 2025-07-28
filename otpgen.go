// Package otpgen provides cryptographically secure OTP token generation.
package otpgen

import (
	"crypto/rand"
	"errors"
)

const (
	MinLength     = 1
	MaxLength     = 1000
	DefaultLength = 6
)

var (
	ErrInvalidLength = errors.New("token length must be between 1 and 1000")
	ErrCryptoFailure = errors.New("failed to generate cryptographically secure random number")
)

// Generate generates a cryptographically secure numeric token of the specified length.
func Generate(length int) (string, error) {
	if length < MinLength || length > MaxLength {
		return "", ErrInvalidLength
	}

	// Generate all random bytes in a single call
	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", ErrCryptoFailure
	}

	// Convert to digits
	token := make([]byte, length)
	for i := 0; i < length; i++ {
		token[i] = '0' + (randomBytes[i] % 10)
	}

	return string(token), nil
}

// GenerateDefault generates a 6-digit token using the default length.
func GenerateDefault() (string, error) {
	return Generate(DefaultLength)
}
