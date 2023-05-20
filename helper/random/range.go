package random

import "math/rand"

func Chance(percent float64) bool {
	return rand.Float64() <= percent
}
