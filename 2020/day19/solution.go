package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var regexMap map[int]string

func isRawValue(regex string) bool {
	return strings.HasPrefix(regex, "\"") && strings.HasSuffix(regex, "\"")
}

func getRawValue(regex string) string {
	return regex[1 : len(regex)-1]
}

func allElementsPresent(regex string) bool {
	return allNonLoopElementsPresent(regex, "")
}

func allNonLoopElementsPresent(regex string, self string) bool {
	elements := strings.Fields(regex)
	for _, el := range elements {
		if el != "|" && el != self {
			elInt, _ := strconv.Atoi(el)
			_, present := regexMap[elInt]
			if !present {
				return false
			}
		}
	}
	return true
}

func buildRegex(input string, self string, maxDepth int) string {
	regexString := ""

	acc := ""
	loopString := ""

	appendAcc := func() {
		if strings.Contains(acc, "*") {
			loopString = acc
		}
		regexString += acc
		acc = ""
	}

	elements := strings.Fields(input)
	for _, el := range elements {
		if el == "|" {
			appendAcc()
			regexString += "|"
		} else if el == self {
			acc += "*"
		} else {
			elInt, _ := strconv.Atoi(el)
			val, _ := regexMap[elInt]
			acc += val
		}
	}
	appendAcc()

	if strings.Contains(regexString, "*") {
		var b strings.Builder
		original := regexString
		for i := 0; i < maxDepth; i++ {
			newLooped := original
			for j := 0; j < i; j++ {
				newLooped = strings.Replace(newLooped, "*", loopString, 1)
			}
			b.WriteString(strings.Replace(newLooped, "*", "", -1))
			if i < maxDepth-1 {
				b.WriteString("|")
			}
		}
		regexString = b.String()
	}

	return "(" + regexString + ")"
}

func main() {
	regexMap = make(map[int]string, 0)
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	messages := strings.Split(parts[1], "\n")

	regexSet := make(map[string]struct{}, 0)
	for _, regex := range strings.Split(parts[0], "\n") {
		regexSet[regex] = struct{}{}
	}

	for len(regexSet) > 0 {
		for k := range regexSet {
			splitRegex := strings.Split(k, ": ")
			number, _ := strconv.Atoi(splitRegex[0])
			if isRawValue(splitRegex[1]) {
				val := getRawValue(splitRegex[1])
				regexMap[number] = val
				delete(regexSet, k)
			} else if allElementsPresent(splitRegex[1]) {
				reg := buildRegex(splitRegex[1], splitRegex[0], 5)
				regexMap[number] = reg
				delete(regexSet, k)
			}
		}
	}

	count := 0
	zeroRegex := regexp.MustCompile(regexMap[0])
	zeroRegex.Longest()
	for _, input := range messages {
		res := zeroRegex.FindString(input)
		if len(res) == len(input) {
			count++
		}
	}
	fmt.Println(count)

	// --- Part 2 copypasta

	regexMap = make(map[int]string, 0)
	regexSet = make(map[string]struct{}, 0)
	skipped := false
	for _, regex := range strings.Split(parts[0], "\n") {
		if !strings.HasPrefix(regex, "8: ") && !strings.HasPrefix(regex, "11: ") {
			regexSet[regex] = struct{}{}
		} else {
			skipped = true
		}
	}
	if skipped {
		regexSet["8: 42 | 42 8"] = struct{}{}
		regexSet["11: 42 31 | 42 11 31"] = struct{}{}
	}

	for len(regexSet) > 0 {
		for k := range regexSet {
			splitRegex := strings.Split(k, ": ")
			number, _ := strconv.Atoi(splitRegex[0])
			if isRawValue(splitRegex[1]) {
				val := getRawValue(splitRegex[1])
				regexMap[number] = val
				delete(regexSet, k)
			} else if allNonLoopElementsPresent(splitRegex[1], splitRegex[0]) {
				reg := buildRegex(splitRegex[1], splitRegex[0], 5)
				regexMap[number] = reg
				delete(regexSet, k)
			}
		}
	}

	count = 0

	zeroRegex = regexp.MustCompile(regexMap[0])
	zeroRegex.Longest()
	for _, input := range messages {
		res := zeroRegex.FindString(input)
		if len(res) == len(input) {
			count++
		}
	}
	fmt.Println(count)

}
