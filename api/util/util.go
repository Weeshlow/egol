package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-gl/mathgl/mgl32"
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

// RandomPosition returns a random vec3
func RandomPosition() mgl32.Vec3 {
	return mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		0.0,
	}
}

// RandomDirection returns a random unit vec3
func RandomDirection() mgl32.Vec3 {
	return mgl32.Vec3{
		(rand.Float32() * 2) - 1,
		(rand.Float32() * 2) - 1,
		0.0,
	}.Normalize()
}
