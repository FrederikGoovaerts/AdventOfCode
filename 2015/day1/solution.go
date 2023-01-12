package main

import (
	"aoc/util"
	"fmt"
)

func part1(in string) int {
	result := 0
	for _, char := range in {
		if char == '(' {
			result++
		} else {
			result--
		}
	}
	return result
}

func part2(in string) int {
	floor := 0
	for index, char := range in {
		if char == '(' {
			floor++
		} else {
			floor--
		}

		if floor < 0 {
			return index + 1
		}
	}
	return -1
}

func main() {
	// input := ex1a
	// input := ex5a
	input := util.FileAsString("input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
