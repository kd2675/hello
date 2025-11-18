package utils

import (
	"math/rand"
	"time"
)

var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func PickUniqueNumbersInRange(min, max, count int) []int {
	if max-min+1 < count {
		return nil
	}

	numbers := make(map[int]bool)
	result := make([]int, 0, count)

	for len(result) < count {
		num := rng.Intn(max-min+1) + min
		if !numbers[num] {
			numbers[num] = true
			result = append(result, num)
		}
	}

	return result
}
