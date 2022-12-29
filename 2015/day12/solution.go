package main

import (
	"aoc/util"
	"encoding/json"
	"fmt"
)

func isRed(in json.RawMessage) bool {
	redString := new(string)
	err := json.Unmarshal(in, redString)
	if err != nil {
		return false
	}
	return *redString == "red"
}

func count(in json.RawMessage, checkRed bool) int {
	result := 0

	// Try decoding number
	number := new(int)
	err := json.Unmarshal(in, number)
	if err == nil {
		return *number
	}

	// Try decoding array
	arr := new([]json.RawMessage)
	err = json.Unmarshal(in, arr)
	if err == nil {
		for _, el := range *arr {
			result += count(el, checkRed)
		}
		return result
	}

	// Try decoding object
	obj := new(map[string]json.RawMessage)
	err = json.Unmarshal(in, obj)
	if err == nil {
		for _, v := range *obj {
			if checkRed && isRed(v) {
				return 0
			}
			result += count(v, checkRed)
		}
	}
	return result

}

func part1(input string) int {
	return count(json.RawMessage(input), false)
}

func part2(input string) int {
	return count(json.RawMessage(input), true)
}

func main() {
	input := util.FileAsString("input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
