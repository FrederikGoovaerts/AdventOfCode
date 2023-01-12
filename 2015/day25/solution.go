package main

import (
	"aoc/util"
	"fmt"
)

const input int64 = 20151125

func getNext(val int64) int64 {
	return (val * 252533) % 33554393
}

func getOrdinalFor(r int, c int) int {
	rOff := 1 - r
	return util.SumUpTo(c-rOff) + rOff
}

func part1(value int64) int64 {
	curr := value
	iMax := getOrdinalFor(3010, 3019)
	for i := 1; i < iMax; i++ {
		curr = getNext(curr)
	}

	return curr
}

func main() {
	fmt.Println(part1(input))
}
