// Package otpgen provides cryptographically secure OTP token generation.
package otpgen

import (
	"crypto/rand"
	"errors"
	"math/big"
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

	const digits = "0123456789"
	digitsLen := big.NewInt(int64(len(digits)))
	token := make([]byte, length)

	for i := range token {
		num, err := rand.Int(rand.Reader, digitsLen)
		if err != nil {
			return "", ErrCryptoFailure
		}
		token[i] = digits[num.Int64()]
	}

	return string(token), nil
}

// GenerateDefault generates a 6-digit token using the default length.
func GenerateDefault() (string, error) {
	return Generate(DefaultLength)
}
