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

func main() {
	dat, err := ioutil.ReadFile("ex2")
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

	coincidences := make([]coincidence, 0)
	firstBus, _ := strconv.Atoi(busses[0])
	for i := 1; i < len(busses); i++ {
		if busses[i] != "x" {
			otherBus, _ := strconv.Atoi(busses[i])
			for j := 1; j < firstBus*otherBus; j++ {
				if (j%firstBus == 0) && (j%otherBus == i) {
					coincidences = append(coincidences, coincidence{firstBus * otherBus, j})
				}
			}
		}
	}
	fmt.Println(coincidences)
	for len(coincidences) > 1 {
		newCoincidences := make([]coincidence, 0)
		first := coincidences[0]
		for i := 1; i < len(coincidences); i++ {
			second := coincidences[i]
			matched := false
			curr := first.occurence
			for !matched {
				if curr%second.interval == (second.interval - second.occurence) {
					matched = true
					newCoincidences = append(newCoincidences, coincidence{first.interval * second.interval, curr})
				} else {
					curr += first.interval
				}
			}
		}
		coincidences = newCoincidences
	}
	fmt.Println(coincidences[0])
}
