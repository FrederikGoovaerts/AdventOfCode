package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bagContents struct {
	color  string
	amount int
}

var colorMap map[string][]bagContents

func canContain(origin string, destination string) bool {
	checked := make(map[string]struct{}, 0)
	next := make([]string, 0)
	for _, el := range colorMap[origin] {
		if el.color == destination {
			return true
		}
		next = append(next, el.color)
	}
	for len(next) > 0 {
		var nextEl string
		nextEl, next = next[0], next[1:]

		for _, el := range colorMap[nextEl] {
			_, alreadyChecked := checked[el.color]
			if el.color == destination {
				return true
			}
			if !alreadyChecked {
				next = append(next, el.color)
			}
			checked[el.color] = struct{}{}
		}
	}
	return false
}

func containsAmount(origin string) int {
	count := 0
	for _, el := range colorMap[origin] {
		count += (containsAmount(el.color) * el.amount)
	}
	return count + 1
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	colorMap = make(map[string][]bagContents, 0)
	for _, line := range lines {
		splitLine := strings.Split(line, " bags contain ")
		if strings.Contains(line, "no other") {
			colorMap[splitLine[0]] = make([]bagContents, 0)
		} else {
			contentsStrings := strings.Split(splitLine[1][:len(splitLine[1])-1], ", ")
			contents := make([]bagContents, 0)
			for _, el := range contentsStrings {
				amount, _ := strconv.Atoi(string(el[0]))
				name := strings.ReplaceAll(strings.ReplaceAll(el[2:], " bags", ""), " bag", "")
				contents = append(contents, bagContents{name, amount})
			}
			colorMap[splitLine[0]] = contents
		}
	}

	count := 0
	for k := range colorMap {
		if canContain(k, "shiny gold") {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(containsAmount("shiny gold") - 1)
}
