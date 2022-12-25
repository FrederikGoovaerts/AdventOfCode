package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var direction = []struct {
	n int
	e int
}{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

type command struct {
	dir    string
	amount int
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	commands := make([]command, 0)
	for _, line := range lines {
		dir := line[:1]
		amount, _ := strconv.Atoi(line[1:])
		commands = append(commands, command{dir, amount})
	}
	angleIndex := 0
	n := 0
	e := 0
	for _, command := range commands {
		switch command.dir {
		case "N":
			n += command.amount
		case "S":
			n += command.amount
		case "E":
			e += command.amount
		case "W":
			e -= command.amount
		case "L":
			angleIndex = (angleIndex - (command.amount / 90) + 4) % 4
		case "R":
			angleIndex = (angleIndex + (command.amount / 90) + 4) % 4
		case "F":
			n += (direction[angleIndex].n * command.amount)
			e += (direction[angleIndex].e * command.amount)
		}
	}
	fmt.Println(math.Abs(float64(n)) + math.Abs(float64(e)))

	n = 0
	e = 0
	waypointN := 1
	waypointE := 10
	for _, command := range commands {
		switch command.dir {
		case "N":
			waypointN += command.amount
		case "S":
			waypointN -= command.amount
		case "E":
			waypointE += command.amount
		case "W":
			waypointE -= command.amount
		case "L":
			times := command.amount / 90
			for i := 0; i < times; i++ {
				lastN := waypointN
				lastE := waypointE
				waypointN = lastE
				waypointE = -lastN
			}
		case "R":
			times := command.amount / 90
			for i := 0; i < times; i++ {
				lastN := waypointN
				lastE := waypointE
				waypointN = -lastE
				waypointE = lastN
			}
		case "F":
			n += waypointN * command.amount
			e += waypointE * command.amount
		}
	}
	fmt.Println(math.Abs(float64(n)) + math.Abs(float64(e)))

}
