package util

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ClampToOne(x int) int {
	if x != 0 {
		return x / Abs(x)
	}
	return x
}
