package otpgen

import (
	"testing"
)

func BenchmarkGenerate6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Generate(6)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerate100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Generate(100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerateDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GenerateDefault()
		if err != nil {
			b.Fatal(err)
		}
	}
}
