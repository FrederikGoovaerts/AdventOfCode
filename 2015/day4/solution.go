package main

import (
	"aoc/util"
	"crypto/md5"
	"fmt"
	"math"
	"strings"
	"sync"
)

func findNum(in string, zeroes int) int {
	pref := strings.Repeat("0", zeroes)
	for i := 0; ; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(in+fmt.Sprint(i))))
		if strings.HasPrefix(hash, pref) {
			return i
		}
	}
}

const JOB_SIZE = 2000
const CONCURRENT_JOBS = 8
const BATCH_SIZE = JOB_SIZE * CONCURRENT_JOBS

func findNumConcurrent(in string, zeroes int) int {
	channel := make(chan int, CONCURRENT_JOBS)
	for i := 0; ; i++ {
		var wg sync.WaitGroup
		for job := 0; job < CONCURRENT_JOBS; job++ {
			wg.Add(1)
			go func(start, stop int) {
				pref := strings.Repeat("0", zeroes)
				for i := start; i < stop; i++ {
					hash := fmt.Sprintf("%x", md5.Sum([]byte(in+fmt.Sprint(i))))
					if strings.HasPrefix(hash, pref) {
						channel <- i
					}
				}
				wg.Done()
			}(i*BATCH_SIZE+job*JOB_SIZE, i*BATCH_SIZE+(job+1)*JOB_SIZE-1)
		}
		wg.Wait()
		if len(channel) > 0 {
			result := math.MaxInt
			for len(channel) > 0 {
				result = util.MinInt(result, <-channel)
			}
			return result
		}
	}
}

func part1(in string) int {
	return findNum(in, 5)
}

func part2(in string) int {
	return findNum(in, 6)
}

func part2Concurrent(in string) int {
	return findNumConcurrent(in, 6)
}

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
