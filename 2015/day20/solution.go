package main

import (
	"aoc/util"
	"fmt"
	"math"
)

const INPUT = 29000000

func getApplicableElves(number int) []int {

	result := make([]int, 0)
	result = append(result, 1)
	result = append(result, number)

	for i := 2; float64(i) <= math.Sqrt(float64(number)); i++ {
		if number%i == 0 {
			result = append(result, i)
			if i != number/i {
				result = append(result, number/i)
			}
		}
	}
	return result
}

func getPresentsForHouse(house int) int {
	return util.Sum(getApplicableElves(house)) * 10
}

func part1(input int) int {
	for house := 1; ; house++ {
		if getPresentsForHouse(house) >= input {
			return house
		}
	}
}

func getApplicableLazyElves(number int) []int {
	result := make([]int, 0)
	for _, e := range getApplicableElves(number) {
		if e*50 >= number {
			result = append(result, e)
		}
	}
	return result
}

func getLazyPresentsForHouse(house int) int {
	return util.Sum(getApplicableLazyElves(house)) * 11
}

func part2(input int) int {
	for house := 1; ; house++ {
		if getLazyPresentsForHouse(house) >= input {
			return house
		}
	}
}

func main() {
	fmt.Println(part1(INPUT))
	fmt.Println(part2(INPUT))
}
