package utils

import (
	"math/rand"
	"time"
)

const (
	minRandomValue = 0
	maxRandomValue = 9
)

var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func PickNumberInRange(min, max int) int {
	return rng.Intn(max-min+1) + min
}

func GenerateRandomValue() int {
	return PickNumberInRange(minRandomValue, maxRandomValue)
}
