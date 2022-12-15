package util

import (
	"fmt"
	"strconv"
	"strings"
)

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

func SerializeCoord(x int, y int) string {
	return fmt.Sprint(x) + " " + fmt.Sprint(y)
}

func DeserializeCoord(ser string) (int, int) {
	parts := strings.Split(ser, " ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}
