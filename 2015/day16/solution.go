package main

import (
	"aoc/util"
	"fmt"
	"regexp"
)

var detection map[string]int = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

type Aunt map[string]int

func (a Aunt) matchesPart1() bool {
	for k, v := range a {
		if v != detection[k] {
			return false
		}
	}
	return true
}

func (a Aunt) matchesPart2() bool {
	for k, v := range a {
		if k == "trees" || k == "cats" {
			if v <= detection[k] {
				return false
			}
		} else if k == "pomeranians" || k == "goldfish" {
			if v >= detection[k] {
				return false
			}
		} else if v != detection[k] {
			return false
		}
	}
	return true
}

var inputRegex = regexp.MustCompile("^Sue (.*): (.*): ([0-9]*), (.*): ([0-9]*), (.*): ([0-9]*)$")

func parse(lines []string) []Aunt {
	aunts := make([]Aunt, 0, len(lines))

	for _, line := range lines {
		aunt := make(Aunt)
		matches := inputRegex.FindStringSubmatch(line)
		firstAmount := util.StringToInt(matches[3])
		aunt[matches[2]] = firstAmount
		secondAmount := util.StringToInt(matches[5])
		aunt[matches[4]] = secondAmount
		thirdAmount := util.StringToInt(matches[7])
		aunt[matches[6]] = thirdAmount

		aunts = append(aunts, aunt)
	}

	return aunts
}

func part1(aunts []Aunt) int {
	for index, aunt := range aunts {
		if aunt.matchesPart1() {
			return index + 1
		}
	}

	return -1
}

func part2(aunts []Aunt) int {
	for index, aunt := range aunts {
		if aunt.matchesPart2() {
			return index + 1
		}
	}

	return -1
}

func main() {
	input := util.FileAsLines("input")
	aunts := parse(input)

	fmt.Println(part1(aunts))
	fmt.Println(part2(aunts))
}
