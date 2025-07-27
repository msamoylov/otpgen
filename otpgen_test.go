package otpgen

import (
	"errors"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		length  int
		wantErr error
	}{
		{"valid_6_digit", 6, nil},
		{"valid_1_digit", 1, nil},
		{"valid_1000_digit", 1000, nil},
		{"invalid_zero_length", 0, ErrInvalidLength},
		{"invalid_negative_length", -1, ErrInvalidLength},
		{"invalid_too_long", 1001, ErrInvalidLength},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			token, err := Generate(tt.length)

			if tt.wantErr != nil {
				if err == nil || !errors.Is(err, tt.wantErr) {
					t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				}
				if token != "" {
					t.Errorf("Generate() token = %q, want empty on error", token)
				}
				return
			}

			if err != nil {
				t.Errorf("Generate() unexpected error = %v", err)
				return
			}

			if len(token) != tt.length {
				t.Errorf("Generate() token length = %d, want %d", len(token), tt.length)
			}

			if !isNumeric(token) {
				t.Errorf("Generate() token contains non-numeric characters")
			}
		})
	}
}

func TestGenerateDefault(t *testing.T) {
	t.Parallel()

	token, err := GenerateDefault()
	if err != nil {
		t.Errorf("GenerateDefault() unexpected error = %v", err)
	}
	if len(token) != DefaultLength {
		t.Errorf("GenerateDefault() token length = %d, want %d", len(token), DefaultLength)
	}
	if !isNumeric(token) {
		t.Errorf("GenerateDefault() token contains non-numeric characters")
	}
}

func TestGenerateUniqueness(t *testing.T) {
	t.Parallel()

	tokens := make(map[string]bool)
	for i := 0; i < 100; i++ {
		token, err := Generate(6)
		if err != nil {
			t.Fatalf("Generate() unexpected error = %v", err)
		}
		tokens[token] = true
	}

	if len(tokens) < 95 {
		t.Errorf("Generate() uniqueness too low: %d unique out of 100", len(tokens))
	}
}

func TestGenerateDistribution(t *testing.T) {
	t.Parallel()

	digitCounts := make(map[rune]int)
	for i := 0; i < 1000; i++ {
		token, err := Generate(1)
		if err != nil {
			t.Fatalf("Generate() unexpected error = %v", err)
		}
		digitCounts[rune(token[0])]++
	}

	for digit := '0'; digit <= '9'; digit++ {
		if digitCounts[digit] == 0 {
			t.Errorf("Digit %c never appeared", digit)
		}
	}
}

func isNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return len(s) > 0
}
