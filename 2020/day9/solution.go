package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFirstNonSum(numbers []int, window int) int {
	currentIndex := 0
	windowMap := make(map[int]int, 0)
	for i := 0; i < window; i++ {
		windowMap[numbers[i]]++
	}
	for (currentIndex + window) < len(numbers) {
		found := false
		curr := numbers[currentIndex+window]
		for i := currentIndex; i < currentIndex+window; i++ {
			a := numbers[i]
			b := curr - a
			_, match := windowMap[b]
			if a != b && match {
				found = true
			}
		}
		if !found {
			return curr
		}
		if windowMap[numbers[currentIndex]] > 1 {
			windowMap[numbers[currentIndex]]--
		} else {
			delete(windowMap, numbers[currentIndex])
		}
		windowMap[curr]++
		currentIndex++
	}

	return -1
}

func findContiguous(numbers []int, val int) (int, int) {
	for first := 0; first < len(numbers); first++ {
		currVal := numbers[first]
		largest := currVal
		smallest := currVal
		nextIndex := first + 1
		for currVal < val {
			if numbers[nextIndex] > largest {
				largest = numbers[nextIndex]
			}
			if numbers[nextIndex] < smallest {
				smallest = numbers[nextIndex]
			}
			currVal += numbers[nextIndex]
			if currVal == val {
				return smallest, largest
			}
			nextIndex++
		}
	}

	return -1, -1
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	numbers := make([]int, 0)
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		numbers = append(numbers, val)
	}
	invalid := getFirstNonSum(numbers, 25)
	fmt.Println(invalid)
	a, b := findContiguous(numbers, invalid)
	fmt.Println(a + b)
}
