package utils

import (
	"math/rand"
	"time"
)

func RandomPId() int64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number with 10 digits
	randomNumber := rand.Intn(9000000000) + 1000000000

	return int64(randomNumber)
}
