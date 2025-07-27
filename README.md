# otpgen

[![Go Reference](https://pkg.go.dev/badge/github.com/msamoylov/otpgen.svg)](https://pkg.go.dev/github.com/msamoylov/otpgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/msamoylov/otpgen)](https://goreportcard.com/report/github.com/msamoylov/otpgen)

A production-grade Go library for generating cryptographically secure one-time password (OTP) tokens. Perfect for 
two-factor authentication, temporary codes and other security applications requiring unpredictable numeric sequences.

Made for secure authentication systems.

## Features

- 🔒Uses `crypto/rand` for secure random number generation.
- ⚡ ~1,073 ns/op for 6-digit tokens with optimized memory allocation.
- 🧪 >90% test coverage with comprehensive edge case testing.
- 🚀 Thread-safe concurrent operation with parallel test execution.
- 📦 Uses only Go standard library.
- 🎯 Comprehensive error handling, validation and documentation.

## Installation

```bash
go get github.com/msamoylov/otpgen
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/msamoylov/otpgen"
)

func main() {
    // Generate a 6-digit OTP token (most common)
    token, err := otpgen.GenerateDefault()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Your OTP: %s\n", token) // Output: Your OTP: 123456
    
    // Generate custom length token
    longToken, err := otpgen.Generate(12)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Long token: %s\n", longToken) // Output: Long token: 987654321098
}
```

## API Reference

### Functions

#### `Generate(length int) (string, error)`

Generates a cryptographically secure numeric token of the specified length.

**Parameters:**
- `length`: Token length (must be between 1 and 1000)

**Returns:**
- `string`: The generated numeric token
- `error`: Error if length is invalid or crypto operations fail

**Example:**
```go
token, err := otpgen.Generate(8)
if err != nil {
    return err
}
// token is now an 8-digit string like "04275183"
```

#### `GenerateDefault() (string, error)`

Generates a 6-digit token using the default length. Equivalent to `Generate(6)`.

**Returns:**
- `string`: The generated 6-digit numeric token
- `error`: Error if crypto operations fail

**Example:**
```go
token, err := otpgen.GenerateDefault()
if err != nil {
    return err
}
// token is now a 6-digit string like "042751"
```

### Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `MinLength` | 1 | Minimum allowed token length |
| `MaxLength` | 1000 | Maximum allowed token length |
| `DefaultLength` | 6 | Standard OTP length used by most services |

### Errors

| Error | Description |
|-------|-------------|
| `ErrInvalidLength` | Returned when token length is not between 1 and 1000 |
| `ErrCryptoFailure` | Returned when cryptographic random number generation fails |

## Usage Examples

### Basic Usage

```go
// Generate standard 6-digit OTP
token, err := otpgen.GenerateDefault()
if err != nil {
    log.Fatal(err)
}
fmt.Println("OTP:", token)
```

### Custom Length Tokens

```go
// Generate tokens of various lengths
lengths := []int{4, 6, 8, 12}

for _, length := range lengths {
    token, err := otpgen.Generate(length)
    if err != nil {
        log.Printf("Error generating %d-digit token: %v", length, err)
        continue
    }
    fmt.Printf("%d-digit token: %s\n", length, token)
}
```

### Error Handling

```go
// Handle invalid input
token, err := otpgen.Generate(0)
if err != nil {
    if errors.Is(err, otpgen.ErrInvalidLength) {
        fmt.Println("Invalid token length specified")
    } else if errors.Is(err, otpgen.ErrCryptoFailure) {
        fmt.Println("Cryptographic operation failed")
    }
}
```

### Using Constants

```go
// Use package constants for validation
if userLength < otpgen.MinLength || userLength > otpgen.MaxLength {
    return fmt.Errorf("length must be between %d and %d", otpgen.MinLength, otpgen.MaxLength)
}

token, err := otpgen.Generate(userLength)
```

## Performance

Benchmark results on Apple M2 Pro:

| Token Length | Time/op | Memory/op | Allocs/op |
|--------------|---------|-----------|-----------|
| 1 digit | 180.1 ns | 48 B | 4 |
| 6 digits | 1,073 ns | 296 B | 20 |
| 12 digits | 2,078 ns | 608 B | 38 |
| 100 digits | 17,124 ns | 5,024 B | 302 |
| 1000 digits | 170,967 ns | 50,048 B | 3,002 |

**Key Performance Characteristics:**
- Linear time complexity O(n) with token length.
- Consistent memory allocation patterns.
- Excellent parallel performance.
- No memory leaks or unnecessary allocations.

### Running Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run benchmarks with memory allocation stats
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkGenerate6
```

## Testing

The package includes comprehensive tests covering:

- Input validation and boundary conditions.
- Error handling and edge cases.
- Cryptographic randomness distribution.
- Concurrent operation safety.
- Performance benchmarking.
- Example usage verification.

### Running Tests

```bash
# Run all tests
go test

# Run tests with coverage
go test -cover

# Run tests in parallel
go test -parallel 8

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Security Considerations

- Uses `crypto/rand` for cryptographically secure random number generation.
- Each digit is independently and uniformly selected from 0-9.
- No predictable patterns or biases in token generation.
- Suitable for security-sensitive applications like 2FA.
- Thread-safe for concurrent use.

## Contributing

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes.
4. Add tests for new functionality.
5. Ensure tests pass and coverage remains high.
6. Run `go fmt ./...` and `go vet ./...`
7. Commit your changes (`git commit -m 'Add amazing feature'`)
8. Push to the branch (`git push origin feature/amazing-feature`)
9. Open a Pull Request.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/msamoylov/otpgen.git
cd otpgen

# Run tests
go test -v

# Run benchmarks
go test -bench=.

# Check code quality
go fmt ./...
go vet ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
