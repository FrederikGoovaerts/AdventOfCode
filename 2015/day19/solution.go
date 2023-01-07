package main

import (
	"aoc/util"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

type Rep struct {
	target      string
	result      string
	targetRegex *regexp.Regexp
	resultRegex *regexp.Regexp
}

func parse(lines []string) ([]Rep, string) {
	replacements := make([]Rep, 0, len(lines)-1)
	medicine := ""

	for index, line := range lines {
		if index == len(lines)-1 {
			medicine = line
		} else {
			parts := strings.Split(line, " => ")
			targetRegex := regexp.MustCompile(parts[0])
			resultRegex := regexp.MustCompile(parts[1])
			replacements = append(replacements, Rep{parts[0], parts[1], targetRegex, resultRegex})
		}
	}

	return replacements, medicine
}

func splice(old string, start, end int, replacement string) string {
	oldArr := []rune(old)
	replArr := []rune(replacement)
	sizeDiff := len(replacement) - (end - start)

	newArr := make([]rune, len(old)+sizeDiff)

	for i := 0; i < len(old)+sizeDiff; i++ {
		if i < start {
			newArr[i] = oldArr[i]
		} else if i >= end+sizeDiff {
			newArr[i] = oldArr[i-sizeDiff]
		} else {
			newArr[i] = replArr[i-start]
		}
	}

	return string(newArr)
}

func part1(reps []Rep, medicine string) int {
	results := make(util.StringSet)

	for _, rep := range reps {
		matches := rep.targetRegex.FindAllStringIndex(medicine, -1)
		for _, m := range matches {
			results.Add(splice(medicine, m[0], m[1], rep.result))
		}
	}

	return len(results)
}

func contractToSingular(curr string, steps int, reps []Rep, shortest *map[string]int) {
	if len(curr) == 1 {
		val, present := (*shortest)[curr]
		if !present {
			(*shortest)[curr] = steps
		} else {
			(*shortest)[curr] = util.MinInt(val, steps)
		}
	}

	for _, rep := range reps {
		matches := rep.resultRegex.FindAllStringIndex(curr, -1)
		for _, m := range matches {
			replaced := splice(curr, m[0], m[1], rep.target)
			contractToSingular(replaced, steps+1, reps, shortest)
			if len(*shortest) > 0 {
				return
			}
		}
	}
}

func getShortest(m map[string]int) (string, int) {
	s := ""
	i := math.MaxInt
	for k, v := range m {
		if v < i {
			s = k
			i = v
		}
	}
	return s, i
}

func part2(reps []Rep, medicine string) int {
	sort.Slice(reps, func(i, j int) bool {
		return len(reps[i].result) > len(reps[j].result)
	})

	curr := medicine
	ops := 0

	for strings.Contains(curr, "Z") {
		i := strings.Index(curr, "Z") + 1
		length := 4
		theMap := make(map[string]int)
		for len(theMap) == 0 {
			part := curr[i-length : i]
			if !strings.Contains(part, "R") || strings.HasPrefix(part, "R") {
				length++
				continue
			}
			contractToSingular(part, 0, reps, &theMap)
			length++
		}
		r, op := getShortest(theMap)
		curr = splice(curr, i-length+1, i, r)
		ops += op
	}

	return len(curr) - 1 + ops
}

func main() {
	// input := util.FileAsLines("ex1")
	// input := util.FileAsLines("ex2")
	// input := util.FileAsLines("input")
	input := util.FileAsLines("input_simplified")
	replacements, medicine := parse(input)

	fmt.Println(part1(replacements, medicine))
	fmt.Println(part2(replacements, medicine))
}
