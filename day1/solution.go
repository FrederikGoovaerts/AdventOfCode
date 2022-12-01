package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(elves []Elf) {
	curr := 0
	for _, elf := range elves {
		cal := elf.totalCalories()
		if cal > curr {
			curr = cal
		}
	}
	fmt.Println(curr)
}

func part2(elves []Elf) {
	calories := make([]int, len(elves))
	for index, elf := range elves {
		calories[index] = elf.totalCalories()
	}
	sort.Ints(calories)
	fmt.Println(calories[len(elves)-1] + calories[len(elves)-2] + calories[len(elves)-3])
}

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

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	calories := make([]int, 0)
	elves := []Elf{}

	for _, line := range lines {
		if line == "" {
			elf := Elf{calories}
			elves = append(elves, elf)
			calories = make([]int, 0)
		} else {
			val, _ := strconv.Atoi(line)
			calories = append(calories, val)
		}
	}
	part1(elves)
	part2(elves)
}
