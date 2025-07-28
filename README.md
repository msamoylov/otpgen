# otpgen

[![Go Reference](https://pkg.go.dev/badge/github.com/msamoylov/otpgen.svg)](https://pkg.go.dev/github.com/msamoylov/otpgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/msamoylov/otpgen)](https://goreportcard.com/report/github.com/msamoylov/otpgen)
[![codecov](https://codecov.io/gh/msamoylov/otpgen/graph/badge.svg?token=X9AYIXE082)](https://codecov.io/gh/msamoylov/otpgen)

Cryptographically secure OTP token generation for Go.

## Installation

```bash
go get github.com/msamoylov/otpgen
```

## Usage

```go
import "github.com/msamoylov/otpgen"

// Generate 6-digit token
token, err := otpgen.GenerateDefault()

// Generate custom length
token, err := otpgen.Generate(8)
```

## API

- `Generate(length int) (string, error)` - Generate token of specified length (1-1000)
- `GenerateDefault() (string, error)` - Generate 6-digit token

Constants: `MinLength`, `MaxLength`, `DefaultLength`  
Errors: `ErrInvalidLength`, `ErrCryptoFailure`

## Features

- Cryptographically secure using `crypto/rand`
- ~100ns for 6-digit tokens (10x faster)
- Thread-safe
- Zero dependencies

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
