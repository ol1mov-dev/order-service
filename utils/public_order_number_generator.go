package utils

import (
	"crypto/rand"
	"fmt"
)

const (
	Letters      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaNum     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	PrefixLength = 3
	SuffixLength = 6
)

func randomChars(n int, charset string) string {
	b := make([]byte, n)
	for i := range b {
		num := randomInt(len(charset))
		b[i] = charset[num]
	}
	return string(b)
}

func randomInt(max int) int {
	// криптографически безопасное число
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return int(b[0]) % max
}

// Public Order Number: ABC-7KJF0W

func GeneratePublicOrderNumber() string {
	prefix := randomChars(PrefixLength, Letters)
	suffix := randomChars(SuffixLength, AlphaNum)
	return fmt.Sprintf("%s-%s", prefix, suffix)
}
