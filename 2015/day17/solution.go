package main

import (
	"aoc/util"
	"fmt"
	"math"
	"strconv"
)

func parse(lines []string) (int, []int) {
	total := 0
	containers := make([]int, 0, len(lines)-1)

	for i, line := range lines {
		if i == 0 {
			total = util.StringToInt(line)
		} else {
			containers = append(containers, util.StringToInt(line))
		}
	}

	return total, containers
}

func part1(total int, containers []int) int {
	fmtString := "%0" + fmt.Sprint(len(containers)) + "s"
	result := 0

	for i := 1; i < util.PowInt(2, len(containers)); i++ {
		binaryString := fmt.Sprintf(fmtString, strconv.FormatInt(int64(i), 2))
		capacity := 0
		for index, val := range binaryString {
			if val == '1' {
				capacity += containers[index]
			}
		}
		if capacity == total {
			result++
		}
	}

	return result
}

func part2(total int, containers []int) int {
	fmtString := "%0" + fmt.Sprint(len(containers)) + "s"
	result := 0
	amount := math.MaxInt

	for i := 1; i < util.PowInt(2, len(containers)); i++ {
		binaryString := fmt.Sprintf(fmtString, strconv.FormatInt(int64(i), 2))
		nbContainers := 0
		capacity := 0
		for index, val := range binaryString {
			if val == '1' {
				capacity += containers[index]
				nbContainers++
			}
		}
		if capacity == total {
			if nbContainers == amount {
				result++
			} else if nbContainers < amount {
				amount = nbContainers
				result = 1
			}
		}
	}

	return result
}
func main() {
	// input := util.FileAsLines("ex1")
	input := util.FileAsLines("input")
	total, containers := parse(input)

	fmt.Println(part1(total, containers))
	fmt.Println(part2(total, containers))
}
