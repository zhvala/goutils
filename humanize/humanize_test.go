package humanize

import (
	"testing"
)

func BenchmarkBytes(b *testing.B) {
	input := int64(64 * 1024 * 1024 * 1024)
	expect := "64.0 GB"
	for i := 0; i < b.N; i++ {
		output := Bytes(input)
		if output != expect {
			b.Fatal(output, expect)
		}
	}
}
