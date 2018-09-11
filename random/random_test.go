package random

import (
	"testing"
)

func BenchmarkLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LetterString(10)
	}
}

func BenchmarkHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HexString(10)
	}
}
