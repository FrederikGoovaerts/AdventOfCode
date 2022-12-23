package main

import (
	"aoc/util"
	"fmt"
	"math"
)

func parse(lines []string) util.StringSet {
	startingCoordinates := make(util.StringSet)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				startingCoordinates[util.SerializeCoord(x, y)] = util.EMPTY_STRUCT
			}
		}

	}

	return startingCoordinates
}

var dirOrder = []string{"n", "s", "w", "e"}

func isDirectionFree(x, y int, dir string, coordinates util.StringSet) bool {
	switch dir {
	case "n":
		_, nePresent := coordinates[util.SerializeCoord(x+1, y-1)]
		_, nPresent := coordinates[util.SerializeCoord(x, y-1)]
		_, nwPresent := coordinates[util.SerializeCoord(x-1, y-1)]
		return !nePresent && !nPresent && !nwPresent
	case "s":
		_, sePresent := coordinates[util.SerializeCoord(x+1, y+1)]
		_, sPresent := coordinates[util.SerializeCoord(x, y+1)]
		_, swPresent := coordinates[util.SerializeCoord(x-1, y+1)]
		return !sePresent && !sPresent && !swPresent
	case "e":
		_, nePresent := coordinates[util.SerializeCoord(x+1, y-1)]
		_, ePresent := coordinates[util.SerializeCoord(x+1, y)]
		_, sePresent := coordinates[util.SerializeCoord(x+1, y+1)]
		return !sePresent && !ePresent && !nePresent
	case "w":
		_, nwPresent := coordinates[util.SerializeCoord(x-1, y-1)]
		_, wPresent := coordinates[util.SerializeCoord(x-1, y)]
		_, swPresent := coordinates[util.SerializeCoord(x-1, y+1)]
		return !swPresent && !wPresent && !nwPresent
	}
	panic("Invalid direction provided")
}

func isFullyFree(x, y int, coordinates util.StringSet) bool {
	_, nePresent := coordinates[util.SerializeCoord(x+1, y-1)]
	_, nPresent := coordinates[util.SerializeCoord(x, y-1)]
	_, nwPresent := coordinates[util.SerializeCoord(x-1, y-1)]
	_, wPresent := coordinates[util.SerializeCoord(x-1, y)]
	_, swPresent := coordinates[util.SerializeCoord(x-1, y+1)]
	_, sPresent := coordinates[util.SerializeCoord(x, y+1)]
	_, sePresent := coordinates[util.SerializeCoord(x+1, y+1)]
	_, ePresent := coordinates[util.SerializeCoord(x+1, y)]
	return !nePresent && !nPresent && !nwPresent && !ePresent && !sePresent && !sPresent && !swPresent && !wPresent
}

func getProposal(x, y int, dir string) string {
	switch dir {
	case "n":
		return util.SerializeCoord(x, y-1)
	case "s":
		return util.SerializeCoord(x, y+1)
	case "e":
		return util.SerializeCoord(x+1, y)
	case "w":
		return util.SerializeCoord(x-1, y)
	}
	panic("Invalid direction provided")
}

func doRound(dirCounter int, coordinates util.StringSet) util.StringSet {
	proposals := make(map[string]string, len(coordinates))
	proposalCounters := make(map[string]int, len(coordinates))

	for coord := range coordinates {
		x, y := util.DeserializeCoord(coord)
		proposal := coord
		if !isFullyFree(x, y, coordinates) {
			for dirIndex := dirCounter; dirIndex < dirCounter+4; dirIndex++ {
				direction := dirOrder[dirIndex%4]
				if isDirectionFree(x, y, direction, coordinates) {
					proposal = getProposal(x, y, direction)
					break
				}
			}
		}
		proposals[coord] = proposal
		val, present := proposalCounters[proposal]
		if !present {
			proposalCounters[proposal] = 1
		} else {
			proposalCounters[proposal] = val + 1
		}
	}

	result := make(util.StringSet, len(coordinates))
	for origin, proposal := range proposals {
		if proposalCounters[proposal] == 1 {
			result[proposal] = util.EMPTY_STRUCT
		} else {
			result[origin] = util.EMPTY_STRUCT
		}
	}
	return result
}

func vis(coordinates util.StringSet) {
	for y := -3; y < 10; y++ {
		for x := -4; x < 12; x++ {
			_, present := coordinates[util.SerializeCoord(x, y)]
			if present {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getEmptySquares(coordinates util.StringSet) int {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for c := range coordinates {
		x, y := util.DeserializeCoord(c)
		minX = util.MinInt(minX, x)
		maxX = util.MaxInt(maxX, x)
		minY = util.MinInt(minY, y)
		maxY = util.MaxInt(maxY, y)
	}
	return ((maxX - minX + 1) * (maxY - minY + 1)) - len(coordinates)
}

func part1(startingCoordinates util.StringSet) int {
	dirCounter := 0
	coordinates := startingCoordinates

	for i := 1; i <= 10; i++ {
		coordinates = doRound(dirCounter, coordinates)
		dirCounter++
	}
	return getEmptySquares(coordinates)
}

func containsSame(oldCoordinates, coordinates util.StringSet) bool {
	for c := range oldCoordinates {
		_, present := coordinates[c]
		if !present {
			return false
		}
	}
	return true
}

func part2(startingCoordinates util.StringSet) int {
	result := 0
	dirCounter := 0
	coordinates := startingCoordinates

	for round := 1; true; round++ {
		oldCoordinates := coordinates
		coordinates = doRound(dirCounter, coordinates)
		if containsSame(oldCoordinates, coordinates) {
			result = round
			break
		}
		dirCounter++
	}

	return result
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")
	startingCoordinates := parse(lines)

	fmt.Println(part1(startingCoordinates))
	fmt.Println(part2(startingCoordinates))
}
