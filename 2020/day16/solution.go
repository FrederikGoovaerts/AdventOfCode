package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type numRange struct {
	lower int
	upper int
}

type field struct {
	name   string
	ranges []numRange
}

func parseField(line string) field {
	split := strings.Split(line, ": ")
	ranges := make([]numRange, 0)
	for _, rangeString := range strings.Split(split[1], " or ") {
		splitRangeString := strings.Split(rangeString, "-")
		val1, _ := strconv.Atoi(splitRangeString[0])
		val2, _ := strconv.Atoi(splitRangeString[1])
		ranges = append(ranges, numRange{val1, val2})
	}
	return field{name: split[0], ranges: ranges}
}

func parseFields(val string) []field {
	result := make([]field, 0)
	for _, fieldString := range strings.Split(val, "\n") {
		result = append(result, parseField(fieldString))
	}
	return result
}

func matchesRanges(val int, ranges []numRange) bool {
	found := false
	for _, fieldRange := range ranges {
		if val >= fieldRange.lower && val <= fieldRange.upper {
			found = true
		}
	}
	return found
}

func findOffendingField(vals []int, fields []field) (int, bool) {
	for _, val := range vals {
		found := false
		for _, field := range fields {
			if matchesRanges(val, field.ranges) {
				found = true
			}
		}
		if !found {
			return val, true
		}
	}
	return -1, false
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	fields := parseFields(parts[0])
	sum := 0
	splitTickets := strings.Split(parts[2], "\n")
	validTickets := make([][]int, 0)
	for i := 1; i < len(splitTickets); i++ {
		vals := make([]int, 0)
		for _, str := range strings.Split(splitTickets[i], ",") {
			converted, _ := strconv.Atoi(str)
			vals = append(vals, converted)
		}
		offendingVal, offending := findOffendingField(vals, fields)
		if offending {
			sum += offendingVal
		} else {
			validTickets = append(validTickets, vals)
		}
	}
	fmt.Println(sum)

	nbFields := len(fields)
	fieldOrder := make([]string, nbFields, nbFields)
	fieldSet := make(map[string][]numRange, 0)
	for _, field := range fields {
		fieldSet[field.name] = field.ranges
	}
	for len(fieldSet) > 0 {
		for i := 0; i < nbFields; i++ {
			if fieldOrder[i] == "" {
				candidates := make(map[string][]numRange, 0)
				for k, v := range fieldSet {
					candidates[k] = v
				}
				for _, ticket := range validTickets {
					val := ticket[i]
					for name, ranges := range candidates {
						if !matchesRanges(val, ranges) {
							delete(candidates, name)
						}
					}
				}
				if len(candidates) == 0 {
					panic("no candidates left!")
				}
				if len(candidates) == 1 {
					for name := range candidates {
						fieldOrder[i] = name
						delete(fieldSet, name)
					}
				}
			}
		}
	}

	ownTicket := strings.Split(strings.Split(parts[1], "\n")[1], ",")
	total := 1
	for index, el := range fieldOrder {
		if strings.HasPrefix(el, "departure") {
			val, _ := strconv.Atoi(ownTicket[index])
			total *= val
		}
	}
	fmt.Println(total)

}
