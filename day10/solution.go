package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var noop = Op{"noop", 0}

type Op struct {
	command string
	value   int
}

func isRelevantCycle(cycle int) bool {
	return (cycle-20)%40 == 0
}

func part1(ops []Op) int {
	result := 0
	x := 1

	for cycleZeroIndexed, op := range ops {
		cycle := cycleZeroIndexed + 1
		if isRelevantCycle(cycle) {
			result += (x * cycle)
		}
		switch op.command {
		case "add":
			x += op.value
		}
	}

	return result
}

func part2(ops []Op) string {
	result := ""
	drawing := make([]bool, len(ops))
	x := 0

	for cycle, op := range ops {
		if cycle%40 >= x && cycle%40 <= x+2 {
			drawing[cycle] = true
		}
		switch op.command {
		case "add":
			x += op.value
		}
	}

	for row := 0; row < 6; row++ {
		for index := 0; index < 40; index++ {
			if drawing[row*40+index] {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}

	return result
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	ops := make([]Op, 0)

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		if line != "" {
			parts := strings.Split(line, " ")
			switch parts[0] {
			case "addx":
				value, _ := strconv.Atoi(parts[1])

				ops = append(ops, noop)
				ops = append(ops, Op{"add", value})
			case "noop":
				ops = append(ops, noop)

			}
		}
	}

	fmt.Println(part1(ops))
	fmt.Println(part2(ops))

}
