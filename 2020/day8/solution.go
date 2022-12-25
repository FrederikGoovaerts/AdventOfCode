package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type operation struct {
	op  string
	val int
}

func tweakPrograms(channel chan []operation, base []operation) {
	for i := 0; i < len(base); i++ {
		curr := base[i]
		if curr.op != "acc" {
			tweaked := make([]operation, len(base))
			copy(tweaked, base)
			if curr.op == "jmp" {
				tweaked[i] = operation{"nop", 0}
			} else {
				tweaked[i] = operation{"jmp", curr.val}
			}
			channel <- tweaked
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	operations := make([]operation, 0)
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		val, _ := strconv.Atoi(splitLine[1])
		operations = append(operations, operation{splitLine[0], val})
	}
	currentPointer := 0
	acc := 0
	visited := make(map[int]struct{}, 0)
	done := false

	for !done {
		visited[currentPointer] = struct{}{}
		current := operations[currentPointer]
		switch current.op {
		case "nop":
			currentPointer++
		case "acc":
			currentPointer++
			acc += current.val
		case "jmp":
			currentPointer += current.val
		}
		_, match := visited[currentPointer]
		done = match
	}
	fmt.Println(acc)

	programChannel := make(chan []operation)
	go tweakPrograms(programChannel, operations)
	attempt := <-programChannel
	done = false
	for !done {
		currentPointer = 0
		acc = 0
		visited = make(map[int]struct{}, 0)
		running := true

		for running && currentPointer < len(attempt) {
			visited[currentPointer] = struct{}{}
			current := attempt[currentPointer]
			switch current.op {
			case "nop":
				currentPointer++
			case "acc":
				currentPointer++
				acc += current.val
			case "jmp":
				currentPointer += current.val
			}
			_, match := visited[currentPointer]
			running = !match
		}
		if currentPointer == len(operations) {
			done = true
			fmt.Println(acc)
		} else {
			attempt = <-programChannel
		}
	}
}
