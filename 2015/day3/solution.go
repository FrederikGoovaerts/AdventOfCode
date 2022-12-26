package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func getDirOffset(dir string) (int, int) {
	switch dir {
	case "<":
		return -1, 0
	case ">":
		return 1, 0
	case "^":
		return 0, 1
	case "v":
		return 0, -1
	}
	panic("Not a valid direction")
}

func part1(directions []string) int {
	delivered := make(map[string]int)
	x, y := 0, 0
	delivered[util.SerializeCoordRaw(0, 0)] = 1
	for _, dir := range directions {
		xOff, yOff := getDirOffset(dir)
		x += xOff
		y += yOff
		val, present := delivered[util.SerializeCoordRaw(x, y)]
		if present {
			delivered[util.SerializeCoordRaw(x, y)] = val + 1
		} else {
			delivered[util.SerializeCoordRaw(x, y)] = 1
		}
	}
	return len(delivered)
}

func part2(directions []string) int {
	delivered := make(map[string]int)
	x1, y1 := 0, 0
	x2, y2 := 0, 0
	delivered[util.SerializeCoordRaw(0, 0)] = 2
	for index, dir := range directions {
		xOff, yOff := getDirOffset(dir)
		var x int
		var y int
		if index%2 == 0 {
			x1 += xOff
			y1 += yOff
			x, y = x1, y1
		} else {
			x2 += xOff
			y2 += yOff
			x, y = x2, y2

		}
		val, present := delivered[util.SerializeCoordRaw(x, y)]
		if present {
			delivered[util.SerializeCoordRaw(x, y)] = val + 1
		} else {
			delivered[util.SerializeCoordRaw(x, y)] = 1
		}
	}
	return len(delivered)
}

func main() {
	input := util.FileAsString("input")
	directions := strings.Split(input, "")

	fmt.Println(part1(directions))
	fmt.Println(part2(directions))
}
