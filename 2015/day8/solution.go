package main

import (
	"aoc/util"
	"fmt"
	"regexp"
)

var escapeRegex = regexp.MustCompile(`\\\\|\\"|\\x[0-9a-f]{2}`)

func getDecodeDifference(in string) int {
	unquoted := in[1 : len(in)-1]
	escapes := escapeRegex.FindAllStringIndex(unquoted, -1)
	diff := 0
	for _, escape := range escapes {
		diff += escape[1] - escape[0] - 1
	}
	return diff + 2
}

var encodeRegex = regexp.MustCompile(`\\|\"`)

func getEncodeDifference(in string) int {
	escapes := encodeRegex.FindAllStringIndex(in, -1)
	return len(escapes) + 2
}

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		diff := getDecodeDifference(line)
		result += diff
	}
	return result
}

func part2(lines []string) int {
	result := 0
	for _, line := range lines {
		diff := getEncodeDifference(line)
		result += diff
	}
	return result
}

func main() {
	// input := util.FileAsLines("ex1")
	input := util.FileAsLines("input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
