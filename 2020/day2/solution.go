package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type password struct {
	min  int
	max  int
	char string
	pass string
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	var passwords []password
	matcher := regexp.MustCompile("([0-9]*)-([0-9]*) (.): (.*)")
	for _, line := range lines {
		res := matcher.FindStringSubmatch(line)
		min, _ := strconv.Atoi(res[1])
		max, _ := strconv.Atoi(res[2])
		passwords = append(passwords, password{min, max, res[3], res[4]})
	}

	part1Count := 0
	for _, password := range passwords {
		matchAmount := strings.Count(password.pass, password.char)
		if matchAmount >= password.min && matchAmount <= password.max {
			part1Count++
		}
	}
	fmt.Println(part1Count)

	part2Count := 0
	for _, password := range passwords {
		minRune := []rune(password.pass)[password.min-1 : password.min]
		maxRune := []rune(password.pass)[password.max-1 : password.max]
		if (string(minRune) == password.char) != (string(maxRune) == password.char) {
			part2Count++
		}
	}
	fmt.Println(part2Count)
}
