package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

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
	sort.Ints(numbers)

	oneJolts := 1
	threeJolts := 1
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1]-numbers[i] == 1 {
			oneJolts++
		}
		if numbers[i+1]-numbers[i] == 3 {
			threeJolts++
		}
	}
	fmt.Println(oneJolts * threeJolts)

	jumpMap := make(map[int]int64)
	jumpMap[0] = 1
	for _, number := range numbers {
		jumpMap[number] = jumpMap[number-1] + jumpMap[number-2] + jumpMap[number-3]
	}
	fmt.Println(jumpMap[numbers[len(numbers)-1]])
}
