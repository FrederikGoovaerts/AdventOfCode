package main

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	calories []int
}

func (e Elf) totalCalories() int {
	calories := 0
	for _, v := range e.calories {
		calories += v
	}
	return calories
}

func parse(line string) []Elf {
	elves := []Elf{}

	for _, pack := range strings.Split(line, "\n\n") {
		calories := make([]int, 0)
		for _, cal := range strings.Split(pack, "\n") {
			val, _ := strconv.Atoi(cal)
			calories = append(calories, val)
		}
		elves = append(elves, Elf{calories})
	}
	return elves
}

func part1(elves []Elf) int {
	curr := 0
	for _, elf := range elves {
		cal := elf.totalCalories()
		if cal > curr {
			curr = cal
		}
	}
	return curr
}

func part2(elves []Elf) int {
	calories := make([]int, len(elves))
	for index, elf := range elves {
		calories[index] = elf.totalCalories()
	}
	sort.Ints(calories)
	return calories[len(elves)-1] + calories[len(elves)-2] + calories[len(elves)-3]
}

func main() {
	// line := util.FileAsString("ex1")
	line := util.FileAsString("input")

	elves := parse(line)

	fmt.Println(part1(elves))
	fmt.Println(part2(elves))
}
