package random

import (
	"math/rand"
	"time"
)

/* 随机生成字符串 */
const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hexBytes    = "0123456789abcdef"
)

var (
	randomEngine = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// LetterString return a pure letter string.
func LetterString(n int) string {
	result := make([]byte, n)
	for i := range result {
		result[i] = letterBytes[randomEngine.Intn(len(letterBytes))]
	}
	return string(result)
}

// HexString return a n bits hex string.
func HexString(n int) string {
	result := make([]byte, n)
	for i := range result {
		result[i] = hexBytes[randomEngine.Intn(len(hexBytes))]
	}
	return string(result)
}
