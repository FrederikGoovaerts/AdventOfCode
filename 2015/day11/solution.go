package main

import (
	"fmt"
	"strings"
)

var nextMap = map[string]string{
	"a": "b",
	"b": "c",
	"c": "d",
	"d": "e",
	"e": "f",
	"f": "g",
	"g": "h",
	"h": "i",
	"i": "j",
	"j": "k",
	"k": "l",
	"l": "m",
	"m": "n",
	"n": "o",
	"o": "p",
	"p": "q",
	"q": "r",
	"r": "s",
	"s": "t",
	"t": "u",
	"u": "v",
	"v": "w",
	"w": "x",
	"x": "y",
	"y": "z",
	"z": "a",
}

func getNext(in string) string {
	result := strings.Split(in, "")
	carry := true

	for i := len(in) - 1; carry && i >= 0; i-- {
		carry = false
		result[i] = nextMap[result[i]]
		if result[i] == "a" {
			carry = true
		}
	}

	return strings.Join(result, "")
}

func isIncreasing(b1, b2, b3 byte) bool {
	return b1+1 == b2 && b1+2 == b3
}

func isSecure(in string) bool {
	hasIncreasing := false
	pairsIndex := make([]int, 0, 2)

	for i := 0; i < len(in); i++ {
		if in[i] == byte('i') || in[i] == byte('o') || in[i] == byte('l') {
			return false
		}
		if !hasIncreasing && i < len(in)-2 && isIncreasing(in[i], in[i+1], in[i+2]) {
			hasIncreasing = true
		}
		if len(pairsIndex) == 0 && i < len(in)-1 && in[i] == in[i+1] {
			pairsIndex = append(pairsIndex, i)
		} else if len(pairsIndex) == 1 && i < len(in)-1 && pairsIndex[0] != i-1 && in[i] == in[i+1] {
			pairsIndex = append(pairsIndex, i)
		}
	}

	return hasIncreasing && len(pairsIndex) == 2
}

func solve(in string) string {
	curr := in

	for {
		curr = getNext(curr)
		if isSecure(curr) {
			return curr
		}
	}
}

func main() {
	// password := ex1
	// password := ex2
	password := input

	part1 := solve(password)
	fmt.Println(part1)
	part2 := solve(part1)
	fmt.Println(part2)
}
