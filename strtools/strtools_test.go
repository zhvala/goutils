package strtools

import (
	"testing"
)

var strData = "abcdefghijklmnopqrstuvwxyz1234567890"
var byteData = []byte(strData)

func TestFastByteToStr(t *testing.T) {
	output := BytesToString(byteData)
	if output != strData {
		t.Fatal(output, strData)
	}
}

func BenchmarkFastByteToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output := BytesToString(byteData)
		if output != strData {
			b.Fatal(output, strData)
		}
	}
}

func BenchmarkByteToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output := string(byteData)
		if output != strData {
			b.Fatal(output, strData)
		}
	}
}
