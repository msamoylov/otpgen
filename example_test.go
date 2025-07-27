package otpgen_test

import (
	"fmt"
	"log"

	"github.com/msamoylov/otpgen"
)

func ExampleGenerate() {
	token, err := otpgen.Generate(6)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Token length: %d\n", len(token))
	fmt.Printf("Token is numeric: %t\n", isNumeric(token))
	// Output:
	// Token length: 6
	// Token is numeric: true
}

func ExampleGenerateDefault() {
	token, err := otpgen.GenerateDefault()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Default token length: %d\n", len(token))
	// Output:
	// Default token length: 6
}

func ExampleGenerate_errorHandling() {
	_, err := otpgen.Generate(0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Output:
	// Error: token length must be between 1 and 1000
}

func isNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return len(s) > 0
}
