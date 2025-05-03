package faker

import (
	"math/rand"
	"time"
)

var intRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int) int {
	if max <= min {
		return min
	}
	return intRand.Intn(max-min+1) + min
}

func RandomFloat(min, max float64) float64 {
	if max <= min {
		return min
	}
	return min + intRand.Float64()*(max-min)
}

func RandomIntSlice(count, min, max int) []int {
	out := make([]int, count)
	for i := 0; i < count; i++ {
		out[i] = RandomInt(min, max)
	}
	return out
}
