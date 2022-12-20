package main

import (
	"aoc/util"
	"fmt"
	"strconv"
)

func parse(lines []string) []int {
	numbers := make([]int, 0, len(lines))

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	return numbers
}

type Element struct {
	value    int
	shift    int
	previous *Element
	next     *Element
}

func (el *Element) getElementInSteps(steps int) *Element {
	curr := el
	for i := 0; i < util.Abs(steps); i++ {
		if steps < 0 {
			curr = curr.previous
		} else {
			curr = curr.next
		}
	}
	return curr
}

func getShortestShift(num, length int) int {
	return num
	// neg := 0
	// pos := 0
	// if num < 0 {
	// 	neg = num % (length - 1)
	// 	pos = neg + length
	// } else if num > 0 {
	// 	pos = num % (length - 1)
	// 	neg = pos - length
	// }
	// if -neg > pos {
	// 	return pos
	// } else {
	// 	return neg
	// }
}

func mix(el *Element) {
	if el.value == 0 {
		return
	}

	prev := el.previous
	next := el.next
	// Uncouple element from circle first
	prev.next = next
	next.previous = prev

	dest := el.getElementInSteps(el.shift)
	if el.shift > 0 {
		el.next = dest.next
		el.previous = dest
		dest.next.previous = el
		dest.next = el
	} else if el.shift < 0 {
		el.previous = dest.previous
		el.next = dest
		dest.previous.next = el
		dest.previous = el
	}
}

func part1(numbers []int) int {
	length := len(numbers)
	originalOrder := make([]*Element, 0, length)

	// Chain setup
	var first *Element = nil
	var zero *Element = nil
	var curr *Element = nil

	for _, num := range numbers {
		shortest := getShortestShift(num, length)
		element := Element{num, shortest, curr, nil}
		curr = &element
		if first == nil {
			first = &element
		}
		if num == 0 {
			zero = &element
		}
		if element.previous != nil {
			element.previous.next = &element
		}

		originalOrder = append(originalOrder, &element)
	}
	curr.next = first
	first.previous = curr

	for _, element := range originalOrder {
		mix(element)
	}
	// for i := 0; i < length; i++ {
	// 	el := zero.getElementInSteps(i)
	// 	fmt.Print(" " + fmt.Sprint(el.value))
	// }
	// fmt.Println()
	// for i := 0; i > -length; i-- {
	// 	el := zero.getElementInSteps(i)
	// 	fmt.Print(" " + fmt.Sprint(el.value))
	// }
	// fmt.Println()

	a := zero.getElementInSteps(1000)
	b := a.getElementInSteps(1000)
	c := b.getElementInSteps(1000)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	return a.value + b.value + c.value
}

func part2(numbers []int) int {
	return 0
}

func main() {
	// lines := util.FileAsLines("ex1")
	// lines := util.FileAsLines("ex2")
	lines := util.FileAsLines("input")
	numbers := parse(lines)

	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}
