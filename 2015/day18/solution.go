package main

import (
	"aoc/util"
	"fmt"
)

func parse(lines []string) (util.StringSet, int) {
	activeLights := make(util.StringSet)

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				activeLights[util.SerializeCoordRaw(x, y)] = util.EMPTY_STRUCT
			}
		}
	}

	return activeLights, len(lines)
}

func getNeighbors(x int, y int, size int) []string {
	n := make([]string, 0, 8)
	// Cardinals
	if x > 0 {
		n = append(n, util.SerializeCoordRaw(x-1, y))
	}
	if y > 0 {
		n = append(n, util.SerializeCoordRaw(x, y-1))
	}
	if x < size-1 {
		n = append(n, util.SerializeCoordRaw(x+1, y))
	}
	if y < size-1 {
		n = append(n, util.SerializeCoordRaw(x, y+1))
	}

	// Diagonals
	if x > 0 && y > 0 {
		n = append(n, util.SerializeCoordRaw(x-1, y-1))
	}
	if x > 0 && y < size-1 {
		n = append(n, util.SerializeCoordRaw(x-1, y+1))
	}
	if x < size-1 && y > 0 {
		n = append(n, util.SerializeCoordRaw(x+1, y-1))
	}
	if x < size-1 && y < size-1 {
		n = append(n, util.SerializeCoordRaw(x+1, y+1))
	}

	return n
}

func doStep(lights util.StringSet, size int) util.StringSet {
	neighborLights := make(map[string]int)

	for k := range lights {
		x, y := util.DeserializeCoordRaw(k)
		for _, n := range getNeighbors(x, y, size) {
			val, present := neighborLights[n]
			if present {
				neighborLights[n] = val + 1
			} else {
				neighborLights[n] = 1
			}
		}
	}

	newLights := make(util.StringSet)
	for loc, litNeighbors := range neighborLights {
		_, wasLit := lights[loc]
		if (wasLit && (litNeighbors == 2 || litNeighbors == 3)) || (!wasLit && litNeighbors == 3) {
			newLights[loc] = util.EMPTY_STRUCT
		}
	}
	return newLights
}

func part1(lights util.StringSet, size int, steps int) int {
	currLights := lights
	for step := 1; step <= steps; step++ {
		currLights = doStep(currLights, size)
	}
	return len(currLights)
}

func setCornersOn(set util.StringSet, size int) util.StringSet {
	set[util.SerializeCoordRaw(0, 0)] = util.EMPTY_STRUCT
	set[util.SerializeCoordRaw(0, size-1)] = util.EMPTY_STRUCT
	set[util.SerializeCoordRaw(size-1, 0)] = util.EMPTY_STRUCT
	set[util.SerializeCoordRaw(size-1, size-1)] = util.EMPTY_STRUCT

	return set
}

func part2(lights util.StringSet, size int, steps int) int {
	currLights := setCornersOn(lights, size)
	for step := 1; step <= steps; step++ {
		currLights = setCornersOn(doStep(currLights, size), size)
	}
	return len(currLights)
}

func main() {
	// input, part1Steps, part2Steps := util.FileAsLines("ex1"), 4, 5
	input, part1Steps, part2Steps := util.FileAsLines("input"), 100, 100
	lights, size := parse(input)

	fmt.Println(part1(lights, size, part1Steps))
	fmt.Println(part2(lights, size, part2Steps))
}
