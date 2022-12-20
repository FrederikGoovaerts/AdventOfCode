package util

import (
	"strconv"
	"strings"
)

func MinInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

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

func Sum(list []int) int {
	result := 0
	for _, el := range list {
		result += el
	}
	return result
}

func Product(list []int) int {
	result := 1
	for _, el := range list {
		result *= el
	}
	return result
}

func SerializeCoord(x int, y int) string {
	return Serialize(x, y)
}

func DeserializeCoord(ser string) (int, int) {
	parts := strings.Split(ser, " ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

type Coord3 struct {
	X int
	Y int
	Z int
}

func (c Coord3) GetNeighbors() []Coord3 {
	return []Coord3{
		{c.X + 1, c.Y, c.Z},
		{c.X - 1, c.Y, c.Z},
		{c.X, c.Y + 1, c.Z},
		{c.X, c.Y - 1, c.Z},
		{c.X, c.Y, c.Z + 1},
		{c.X, c.Y, c.Z - 1},
	}
}

func SerializeCoord3(c Coord3) string {
	return SerializeCoord3Raw(c.X, c.Y, c.Z)
}

func SerializeCoord3Raw(x int, y int, z int) string {
	return Serialize(x, y, z)
}

func DeserializeCoord3(ser string) Coord3 {
	x, y, z := DeserializeCoord3Raw(ser)
	return Coord3{x, y, z}
}

func DeserializeCoord3Raw(ser string) (int, int, int) {
	parts := strings.Split(ser, " ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return x, y, z
}
