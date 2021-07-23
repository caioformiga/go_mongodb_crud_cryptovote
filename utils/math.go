package utils

//
/*
	Abs retorna o valor absoluto de x
	Warning: The smallest value of a signed integer doesn’t have a matching positive value.
	math.MinInt64 is -9223372036854775808, but
	math.MaxInt64 is 9223372036854775807.
	Unfortunately, our Abs function returns a negative value in this case.

*/
func Abs(x int) int {
	a := int64(x)
	if a < 0 {
		return int(-a)
	}
	return int(a)
}

func Maior(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
