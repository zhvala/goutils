package timestamp

import (
	"testing"
)

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Second()
	}
}

func BenchmarkSecondStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SecondStr()
	}
}

func BenchmarkMillisecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Millisecond()
	}
}

func BenchmarkMillsecondStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MillisecondStr()
	}
}
