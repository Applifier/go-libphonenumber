package phonenumber

import (
	"testing"
)

func BenchmarkIsPossibleNumberPositiveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPossibleNumber("+358401231234", "FI")
	}
}

func BenchmarkIsPossibleNumberPositiveHardCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPossibleNumber("040-123 4321", "FI")
	}
}

func BenchmarkIsPossibleNumberNegativeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPossibleNumber("asbdgjwoprg", "FI")
	}
}

func BenchmarkParseEasyPositiveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("+358401231234", "FI")
	}
}

func BenchmarkParseHardPositiveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("040-123 4321", "FI")
	}
}

func BenchmarkParseNegativeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("aeoighowHG", "FI")
	}
}
