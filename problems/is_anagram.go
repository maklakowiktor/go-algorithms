package problems

import (
	"math"
)

func IsAnagram(a string, b string) bool {
	count := make(map[string]uint, int(math.Max(float64(len(a)), float64(len(b)))))

	for c := range a {
		count[string(c)] += 1
	}

	for c := range b {
		count[string(c)] -= 1
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
