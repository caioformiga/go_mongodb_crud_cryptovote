package utils

/*
	Abs returns the absolute value of x, case -1, return 1.

	Warning: The smallest value of a signed integer does not have a
	corresponding positive value, see details at:
	[https://yourbasic.org/golang/absolute-value-int-float/]
*/
func Abs(x int64) int64 {
	a := int64(x)
	if a < 0 {
		return (-a)
	}
	return (a)
}
