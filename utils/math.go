package utils

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Maior(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
