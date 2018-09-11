package routine

import (
	"testing"
)

func BenchmarkGoroutineID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetID()
	}
}
