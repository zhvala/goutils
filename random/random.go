package random

import (
	"math/rand"
	"time"
)

/* 随机生成字符串 */
const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes = "0123456789abcdef"
)

// LetterString return pure letter string
func LetterString(n int) string {
	var randomEngine = rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, n)
	for i := range result {
		result[i] = letterBytes[randomEngine.Intn(len(letterBytes))]
	}
	return string(result)
}

// HexString return hex string
func HexString(n int) string {
	var randomEngine = rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, n)
	for i := range result {
		result[i] = numberBytes[randomEngine.Intn(len(numberBytes))]
	}
	return string(result)
}
