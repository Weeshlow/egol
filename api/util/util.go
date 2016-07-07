package util

import (
	"fmt"
	"math/rand"
	"time"
)

// RandString returns a random string of length N.
func RandString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := len(letters)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(length)]
	}
	return string(b)
}

// RandID returns a random ID.
func RandID() string {
	return fmt.Sprintf("%s-%s-%s-%s",
		RandString(4),
		RandString(2),
		RandString(4),
		RandString(4))
}

// Timestamp returns the timestamp in milliseconds
func Timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
