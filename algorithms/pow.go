package algorithms

func EffectivePow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y < 0 {
		x = 1 / x
		y = -y
	}
	result := 1
	for y > 0 {
		if y%2 == 1 {
			result *= x
			x *= x
			y /= 2
		}
	}
	return result
}
