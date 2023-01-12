package main

import (
	"aoc/util"
	"fmt"
	"regexp"
)

var vowelRegex = regexp.MustCompile("a|e|i|o|u")

func containsVowels(in string, amount int) bool {
	return len(vowelRegex.FindAllString(in, amount)) >= amount
}

func containsDouble(in string) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] == in[i+1] {
			return true
		}
	}
	return false
}

var comboRegex = regexp.MustCompile("(ab)|(cd)|(pq)|(xy)")

func containsNaughtyCombo(in string) bool {
	return comboRegex.FindStringIndex(in) != nil
}

func isNice(in string) bool {
	return containsVowels(in, 3) && containsDouble(in) && !containsNaughtyCombo(in)
}

func part1(strings []string) int {
	result := 0
	for _, str := range strings {
		if isNice(str) {
			result++
		}
	}
	return result
}

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var precomputedRegexes = make([]*regexp.Regexp, 0, util.PowInt(len(letters), 2))

func precomputeRegexes() {
	for _, first := range letters {
		for _, second := range letters {
			duplicateRegex := regexp.MustCompile(first + second)
			precomputedRegexes = append(precomputedRegexes, duplicateRegex)
		}
	}
}

func containsDuplicate(in string) bool {
	for _, exp := range precomputedRegexes {
		if len(exp.FindAllString(in, 2)) >= 2 {
			return true
		}
	}
	return false
}

func containsStaggeredDouble(in string) bool {
	for i := 0; i < len(in)-2; i++ {
		if in[i] == in[i+2] {
			return true
		}
	}
	return false
}

func isNewNice(in string) bool {
	return containsDuplicate(in) && containsStaggeredDouble(in)
}

func part2(strings []string) int {
	precomputeRegexes()

	result := 0
	for _, str := range strings {
		if isNewNice(str) {
			result++
		}
	}
	return result
}

func main() {
	input := util.FileAsLines("input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
