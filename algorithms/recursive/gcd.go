package recursive

func GreatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GreatestCommonDivisor(b, a%b)
	}
}
