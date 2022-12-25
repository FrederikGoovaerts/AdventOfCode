package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func slopeIncidence(x int, y int) int {
	incidence := 0
	currX := 0
	currY := 0
	for currY < len(treeMap)-1 {
		currX = ((x + currX) % len(treeMap[0]))
		currY += y
		incidence += treeMap[currY][currX]
	}
	return incidence
}

var treeMap [][]int

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	treeMap = make([][]int, 0)

	for _, line := range lines {
		treeLine := make([]int, 0)
		for _, el := range line {
			if el == '.' {
				treeLine = append(treeLine, 0)
			} else {
				treeLine = append(treeLine, 1)
			}
		}
		treeMap = append(treeMap, treeLine)
	}
	fmt.Println(slopeIncidence(3, 1))
	fmt.Println(slopeIncidence(1, 1) * slopeIncidence(3, 1) * slopeIncidence(5, 1) * slopeIncidence(7, 1) * slopeIncidence(1, 2))
}
