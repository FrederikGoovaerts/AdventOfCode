package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

// For reference: we see right as X+ and up as Y+

// EX1
// const MIN_X = 0
// const MAX_X = 5
// const MIN_Y = 0
// const MAX_Y = 5

// EX2
// const MIN_X = -11
// const MAX_X = 14
// const MIN_Y = -5
// const MAX_Y = 15

// func visualize(ropeX []int, ropeY []int) {
// 	for y := MAX_Y; y >= MIN_Y; y-- {
// 		for x := MIN_X; x <= MAX_X; x++ {
// 			found := false
// 			for id := 0; id < len(ropeX); id++ {
// 				if ropeX[id] == x && ropeY[id] == y {
// 					found = true
// 					fmt.Print(id)
// 					break
// 				}
// 			}
// 			if !found {
// 				fmt.Print(".")
// 			}
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

type Move struct {
	direction string
	distance  int
}

func addTailLocation(x int, y int, locations *map[string]struct{}) {
	location := fmt.Sprint(x) + " " + fmt.Sprint(y)
	(*locations)[location] = util.EMPTY_STRUCT
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

func solve(moves []Move, ropeLength int) int {
	tailLocations := make(map[string]struct{})
	ropeX := make([]int, ropeLength)
	ropeY := make([]int, ropeLength)

	for _, move := range moves {
		xDiff, yDiff := getHeadMove(move.direction)
		for i := 0; i < move.distance; i++ {
			// Move head
			ropeX[0] += xDiff
			ropeY[0] += yDiff
			// Drag tails
			for i := 1; i < ropeLength; i++ {
				if util.Abs(ropeX[i]-ropeX[i-1]) > 1 || util.Abs(ropeY[i]-ropeY[i-1]) > 1 {
					ropeX[i] += util.ClampToOne(ropeX[i-1] - ropeX[i])
					ropeY[i] += util.ClampToOne(ropeY[i-1] - ropeY[i])
				}
			}
			// Note tail location
			addTailLocation(ropeX[ropeLength-1], ropeY[ropeLength-1], &tailLocations)
		}

	}
	return len(tailLocations)
}

func main() {
	lines := util.FileAsLines("input")
	moves := make([]Move, 0)

	for _, line := range lines {
		if line != "" {
			parts := strings.Split(line, " ")
			distance, _ := strconv.Atoi(parts[1])
			moves = append(moves, Move{parts[0], distance})
		}
	}

	fmt.Println(solve(moves, 2))
	fmt.Println(solve(moves, 10))
}
