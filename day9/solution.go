package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// For reference: we see right as X+ and up as Y+

var empty = struct{}{}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Move struct {
	direction string
	distance  int
}

func addTailLocation(x int, y int, locations *map[string]struct{}) {
	location := fmt.Sprint(x) + " " + fmt.Sprint(y)
	(*locations)[location] = empty
}

func getHeadMove(direction string) (int, int) {
	switch direction {
	case "R":
		return 1, 0
	case "L":
		return -1, 0
	case "U":
		return 0, 1
	case "D":
		return 0, -1
	default:
		panic("illegal direction")
	}
}

func transSyncDiff(headX int, headY int, tailX int, tailY int, direction string) (int, int) {
	switch direction {
	case "R":
		return 0, headY - tailY
	case "L":
		return 0, headY - tailY
	case "U":
		return headX - tailX, 0
	case "D":
		return headX - tailX, 0
	default:
		panic("illegal direction")
	}
}

func part1(moves []Move) int {
	tailLocations := make(map[string]struct{})
	headX := 0
	headY := 0
	tailX := 0
	tailY := 0

	for _, move := range moves {
		xDiff, yDiff := getHeadMove(move.direction)
		for i := 0; i < move.distance; i++ {
			// Move head
			headX += xDiff
			headY += yDiff
			// Drag tail
			if Abs(tailX-headX) > 1 || Abs(tailY-headY) > 1 {
				syncX, syncY := transSyncDiff(headX, headY, tailX, tailY, move.direction)
				tailX += xDiff + syncX
				tailY += yDiff + syncY
			}
			// Note tail location
			addTailLocation(tailX, tailY, &tailLocations)
		}

	}
	return len(tailLocations)
}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	moves := make([]Move, 0)

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		if line != "" {
			parts := strings.Split(line, " ")
			distance, _ := strconv.Atoi(parts[1])
			moves = append(moves, Move{parts[0], distance})
		}
	}

	fmt.Println(part1(moves))
}
