package main

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func parse(lines []string) [][]int {
	result := make([][]int, len(lines))

	for index, line := range lines {
		res := make([]int, 3)
		parts := strings.Split(line, "x")
		res[0], _ = strconv.Atoi(parts[0])
		res[1], _ = strconv.Atoi(parts[1])
		res[2], _ = strconv.Atoi(parts[2])
		result[index] = res
	}

	return result
}

func getWrappingWithSlack(in []int) int {
	sort.Ints(in)
	surface := (2 * in[0] * in[1]) + (2 * in[1] * in[2]) + (2 * in[0] * in[2])
	return surface + (in[0] * in[1])
}

func getRibbonAndBow(in []int) int {
	sort.Ints(in)
	circum := (2*in[0] + 2*in[1])
	vol := in[0] * in[1] * in[2]
	return circum + vol
}

func part1(sizes [][]int) int {
	result := 0

	for _, size := range sizes {
		result += getWrappingWithSlack(size)
	}

	return result
}

func part1Concurrent(sizes [][]int) int {
	channel := make(chan int, len(sizes))
	var wg sync.WaitGroup
	for _, size := range sizes {
		wg.Add(1)
		go func(in []int) {
			channel <- getWrappingWithSlack(in)
			wg.Done()
		}(size)
	}
	wg.Wait()
	close(channel)

	result := 0

	for len(channel) > 0 {
		result += <-channel
	}

	return result
}

func part2(sizes [][]int) int {
	result := 0

	for _, size := range sizes {
		result += getRibbonAndBow(size)
	}

	return result
}

func main() {
	input := util.FileAsLines("input")
	sizes := parse(input)

	fmt.Println(part1(sizes))
	fmt.Println(part2(sizes))
}
