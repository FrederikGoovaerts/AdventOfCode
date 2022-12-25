package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

func decodeValue(in string) int {
	switch in {
	case "=":
		return -2
	case "-":
		return -1
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	}
	panic("illegal value supplied")
}

func decodeSnafu(in string) int {
	result := 0
	parts := strings.Split(in, "")
	for i := len(parts) - 1; i >= 0; i-- {
		val := decodeValue(parts[len(in)-i-1])
		result += val * (util.PowInt(5, i))
	}
	return result
}

// cheaper than string -> int -> string
func getNext(in string) string {
	switch in {
	case "0":
		return "1"
	case "1":
		return "2"
	case "2":
		return "3"
	case "3":
		return "4"
	case "4":
		return "5"
	}
	panic("illegal input")
}

func encodeSnafu(number int) string {
	stringified := "0" + strconv.FormatInt(int64(number), 5)
	parts := strings.Split(stringified, "")
	for i := len(parts) - 1; i > 0; i-- {
		val := parts[i]
		if val == "3" {
			parts[i] = "="
			parts[i-1] = getNext(parts[i-1])
		} else if val == "4" {
			parts[i] = "-"
			parts[i-1] = getNext(parts[i-1])
		} else if val == "5" {
			parts[i] = "0"
			parts[i-1] = getNext(parts[i-1])
		}
	}
	if parts[0] == "0" {
		parts[0] = ""
	}
	result := strings.Join(parts, "")

	return result
}

func solve(numbers []string) string {
	result := 0
	for _, number := range numbers {
		result += decodeSnafu(number)
	}
	return encodeSnafu(result)
}

func main() {
	// numbers := util.FileAsLines("ex1")
	numbers := util.FileAsLines("input")

	fmt.Println(solve(numbers))
}
