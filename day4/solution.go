package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Section struct {
	start int
	end   int
}

func (s Section) fullyOverlaps(other Section) bool {
	return s.start <= other.start && other.end <= s.end
}

func (s Section) overlaps(other Section) bool {
	return s.start <= other.start && other.start <= s.end || s.start <= other.end && other.end <= s.end
}

type Pair struct {
	first  Section
	second Section
}

func (p Pair) fullyOverlaps() bool {
	return p.first.fullyOverlaps(p.second) || p.second.fullyOverlaps(p.first)
}

func (p Pair) overlaps() bool {
	return p.first.overlaps(p.second) || p.second.overlaps(p.first)
}

func part1(pairs []Pair) int {
	result := 0
	for _, pair := range pairs {
		if pair.fullyOverlaps() {
			result++
		}
	}
	return result
}
func part2(pairs []Pair) int {
	result := 0
	for _, pair := range pairs {
		if pair.overlaps() {
			result++
		}
	}
	return result
}

func main() {
	lines := util.FileAsLines("input")
	pairs := []Pair{}

	for _, line := range lines {
		if line != "" {
			pairStrings := strings.Split(line, ",")
			pair1String := strings.Split(pairStrings[0], "-")
			pair1First, _ := strconv.Atoi(pair1String[0])
			pair1Second, _ := strconv.Atoi(pair1String[1])
			pair2String := strings.Split(pairStrings[1], "-")
			pair2First, _ := strconv.Atoi(pair2String[0])
			pair2Second, _ := strconv.Atoi(pair2String[1])
			pairs = append(pairs, Pair{Section{pair1First, pair1Second}, Section{pair2First, pair2Second}})
		}
	}
	fmt.Println(part1(pairs))
	fmt.Println(part2(pairs))
}
