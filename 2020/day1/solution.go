package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(values []int) {
	for i := 0; i < len(values); i++ {
		a := values[i]
		for j := i + 1; j < len(values); j++ {
			b := values[j]
			if a+b == 2020 {
				fmt.Println(a * b)
			}
		}
	}
}

func part2(values []int) {
	for i := 0; i < len(values); i++ {
		a := values[i]
		for j := i + 1; j < len(values); j++ {
			b := values[j]
			for k := j + 1; k < len(values); k++ {
				c := values[k]
				if a+b+c == 2020 {
					fmt.Println(a * b * c)
				}
			}
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	var values []int
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		values = append(values, val)
	}
	part1(values)
	part2(values)
}
