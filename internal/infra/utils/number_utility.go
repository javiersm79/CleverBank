package utils

import (
	"math/rand"
	"time"
)

func GenerateNumber() int {
	min := 950000
	max := 9500000
	// set seed
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

func GenerateNumber2() int {
	min := 8850000
	max := 950000000
	// set seed
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}
