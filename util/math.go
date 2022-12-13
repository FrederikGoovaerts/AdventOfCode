package util

const MAX_INT = int(^uint(0) >> 1)

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
