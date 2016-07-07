package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/unchartedsoftware/egol/api/sim"
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

// MarshalState marshals a slice of organisms.
func MarshalState(organisms []*sim.Organism) ([]byte, error) {
	if len(organisms) == 0 {
		return make([]byte, 0), nil
	}
	sample, err := organisms[0].Marshal()
	if err != nil {
		return nil, err
	}
	numBytes := len(sample) * len(organisms)
	bytes := make([]byte, numBytes)
	for index, organism := range organisms {
		buff, err := organism.Marshal()
		if err != nil {
			return nil, err
		}
		copy(bytes[index*numBytes:], buff[0:])
	}
	return bytes, nil
}

// MarshalUpdates marshals a slice of updates.
func MarshalUpdates(updates []*sim.Update) ([]byte, error) {
	if len(updates) == 0 {
		return make([]byte, 0), nil
	}
	sample, err := updates[0].Marshal()
	if err != nil {
		return nil, err
	}
	numBytes := len(sample) * len(updates)
	bytes := make([]byte, numBytes)
	for index, update := range updates {
		buff, err := update.Marshal()
		if err != nil {
			return nil, err
		}
		copy(bytes[index*numBytes:], buff[0:])
	}
	return bytes, nil
}
