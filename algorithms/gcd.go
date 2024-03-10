package algorithms

import (
	"math"
)

func NaiveGCD(a, b int) int {
	gcd := 1
	for d := 2; d < int(math.Max(float64(a), float64(b))); d++ {
		if a%d == 0 && b%d == 0 {
			gcd = d
		}
	}
	return gcd
}
