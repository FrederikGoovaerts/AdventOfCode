package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type schedule struct {
	length int
	offset int
}

type coincidence struct {
	interval  int
	occurence int
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func findSync(xPeriod int, xOffset int, yPeriod int, yOffset int) (int, int) {
	initialOffset := min(xOffset, yOffset)
	normXOff := xOffset - initialOffset
	normYOff := yOffset - initialOffset
	currX := normXOff + xPeriod
	found := false
	for !found {
		isSync := (currX-normYOff)%yPeriod == 0
		if isSync {
			found = true
		} else {
			currX += xPeriod
		}
	}
	return (xPeriod * yPeriod), currX + initialOffset
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	arrival, _ := strconv.Atoi(lines[0])
	busses := strings.Split(lines[1], ",")
	lowest := math.MaxInt64
	lowestBus := ""
	for _, line := range busses {
		if line != "x" {
			departure, _ := strconv.Atoi(line)
			diff := departure - (arrival % departure)

			if diff < lowest {
				lowest = diff
				lowestBus = line
			}
		}
	}
	bus, _ := strconv.Atoi(lowestBus)
	fmt.Println(lowest * bus)

	currentPeriod, _ := strconv.Atoi(busses[0])
	currentOffset := 0
	for i := 1; i < len(busses); i++ {
		if busses[i] != "x" {
			nextPeriod, _ := strconv.Atoi(busses[i])
			currentPeriod, currentOffset = findSync(currentPeriod, currentOffset, nextPeriod, i)
		}
	}
	fmt.Println(currentPeriod - currentOffset)

}
