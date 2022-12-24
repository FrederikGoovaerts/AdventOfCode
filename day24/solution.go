package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Board struct {
	// These max values are not counting the last end square
	maxX          int
	maxY          int
	cycleTime     int
	blockedByTime map[int]*util.StringSet
}

func getSymbolOffset(symbol string) (int, int) {
	switch symbol {
	case ">":
		return 1, 0
	case "<":
		return -1, 0
	case "^":
		return 0, -1
	case "v":
		return 0, 1
	}
	panic("Wrong symbol provided!")
}

func parse(lines []string) Board {
	width := len(lines[0]) - 2
	height := len(lines) - 2
	cycleTime := util.Lcm(width, height)

	// Initialize empty set of blocked positions for each timestep in the cycle
	blockedByTime := make(map[int]*util.StringSet, 600)
	for i := 0; i < cycleTime; i++ {
		set := make(util.StringSet)
		blockedByTime[int(i)] = &set
	}

	// For each blizzard in the input, fill its position per timestep in the corresponding set
	for y := 0; y < height; y++ {
		parts := strings.Split(lines[y+1], "")
		for x := 0; x < width; x++ {
			symbol := parts[x+1]
			if symbol == ">" || symbol == "<" || symbol == "^" || symbol == "v" {
				xOff, yOff := getSymbolOffset(symbol)
				for i := 0; i < cycleTime; i++ {
					(*blockedByTime[int(i)])[util.SerializeCoordRaw(util.PosMod(x+(xOff*i), width), util.PosMod(y+(yOff*i), height))] = util.EMPTY_STRUCT
				}
			}
		}
	}

	return Board{int(width - 1), int(height - 1), int(cycleTime), blockedByTime}
}

type Coord struct {
	x int
	y int
}

func getNeighbors(x, y, maxX, maxY int) []Coord {
	if x == 0 {
		if y == -1 {
			return []Coord{{0, 0}}
		} else if y == maxY {
			return []Coord{{1, y}, {0, y - 1}}
		} else {
			return []Coord{{1, y}, {0, y + 1}, {0, y - 1}}
		}
	} else if x == maxX {
		if y == 0 {
			return []Coord{{maxX, 1}, {maxX - 1, 0}}
		} else if y == maxY+1 {
			return []Coord{{maxX, maxY - 1}}
		} else {
			return []Coord{{maxX, y + 1}, {maxX, y - 1}, {maxX - 1, y}}
		}
	} else {
		if y == 0 {
			return []Coord{{x, 1}, {x + 1, 0}, {x - 1, 0}}
		} else if y == maxY {
			return []Coord{{x + 1, maxY}, {x, maxY - 1}, {x - 1, maxY}}
		} else {
			return []Coord{{x, y + 1}, {x + 1, y}, {x - 1, y}, {x, y - 1}}
		}
	}
}

func getNeighborsAndWait(x, y, maxX, maxY int) []Coord {
	n := getNeighbors(x, y, maxX, maxY)
	return append(n, Coord{x, y})
}

type State struct {
	x    int
	y    int
	step int
}

func part1(board Board) int {
	visited := make(util.StringSet)

	states := make([]State, 0)
	states = append(states, State{0, -1, 0})

	for len(states) > 0 {
		state := states[0]
		states = states[1:]

		if state.x == board.maxX && state.y == board.maxY {
			return state.step + 1
		}

		nextTimeBlockers := board.blockedByTime[(state.step+1)%board.cycleTime]

		for _, dest := range getNeighborsAndWait(state.x, state.y, board.maxX, board.maxY) {
			// Should be an unoccupied space next round
			if !nextTimeBlockers.Has(util.Serialize(dest.x, dest.y)) {
				// Should not be a visited state
				if !visited.Has(util.Serialize(dest.x, dest.y, (state.step+1)%board.cycleTime)) {
					visited.Add(util.Serialize(dest.x, dest.y, (state.step+1)%board.cycleTime))
					states = append(states, State{dest.x, dest.y, state.step + 1})
				}
			}
		}
	}

	return -1
}

func part2(board Board) int {
	visited := make(util.StringSet)

	states := make([]State, 0)
	states = append(states, State{0, -1, 0})

	reached1 := false
	reached2 := false

	for len(states) > 0 {
		state := states[0]
		states = states[1:]

		if state.x == board.maxX && state.y == board.maxY && !reached1 {
			reached1 = true
			visited = make(util.StringSet)
			states = make([]State, 0)
			states = append(states, State{board.maxX, board.maxY + 1, state.step + 1})
			continue
		}

		if state.x == 0 && state.y == 0 && reached1 && !reached2 {
			reached2 = true
			visited = make(util.StringSet)
			states = make([]State, 0)
			states = append(states, State{0, -1, state.step + 1})
			continue
		}

		if state.x == board.maxX && state.y == board.maxY && reached1 && reached2 {
			return state.step + 1
		}

		nextTimeBlockers := board.blockedByTime[(state.step+1)%board.cycleTime]

		for _, dest := range getNeighborsAndWait(state.x, state.y, board.maxX, board.maxY) {
			// Should be an unoccupied space next round
			if !nextTimeBlockers.Has(util.Serialize(dest.x, dest.y)) {
				// Should not be a visited state
				if !visited.Has(util.Serialize(dest.x, dest.y, (state.step+1)%board.cycleTime)) {
					visited.Add(util.Serialize(dest.x, dest.y, (state.step+1)%board.cycleTime))
					states = append(states, State{dest.x, dest.y, state.step + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")
	board := parse(lines)

	fmt.Println(part1(board))
	fmt.Println(part2(board))
}
