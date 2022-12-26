package main

import (
	"crypto/md5"
	"fmt"
	"strings"
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

func part1(in string) int {
	return findNum(in, 5)
}

func part2(in string) int {
	return findNum(in, 6)
}

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
