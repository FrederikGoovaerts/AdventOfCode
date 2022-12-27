package main

import (
	"fmt"
	"strings"
)

const input = "1113122113"

var input_el = []string{"Fr"}

const conway_constant_rounded float64 = 1.303577269034

func getNext(in string) string {
	chars := strings.Split(in, "")
	result := ""

	count := 1
	el := chars[0]

	for index := 1; index < len(in); index++ {
		if chars[index] == el {
			count++
		} else if chars[index] != el {
			result += fmt.Sprint(count) + el
			el = chars[index]
			count = 1
		}
	}
	result += fmt.Sprint(count) + el

	return result
}

func part1(in string) int {
	curr := in
	for i := 0; i < 40; i++ {
		curr = getNext(curr)
	}
	return len(curr)
}

func getNextElements(elements []string) []string {
	result := make([]string, 0)

	for _, el := range elements {
		result = append(result, evolutionMap[el]...)
	}

	return result
}

func getLengthAt(in []string, turns int) int {
	curr := in
	for i := 0; i < turns; i++ {
		curr = getNextElements(curr)
	}

	result := 0
	for _, el := range curr {
		result += len(elMap[el])
	}
	return result
}

func part1Elements(in []string) int {
	return getLengthAt(in, 40)
}

func part2Elements(in []string) int {
	return getLengthAt(in, 50)
}

func main() {
	fmt.Println(part1Elements(input_el))
	fmt.Println(part2Elements(input_el))
}
