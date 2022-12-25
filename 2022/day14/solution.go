package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type CaveFilling int8

const (
	Rock CaveFilling = 0
	Sand CaveFilling = 1
)

type Cave map[string]CaveFilling

func (c *Cave) inputPlugged() bool {
	return (*c)[util.SerializeCoordRaw(500, 0)] == Sand
}

func parse(lines []string) (map[string]CaveFilling, int) {
	cave := Cave(make(map[string]CaveFilling))
	bedrock := 0

	for _, line := range lines {
		if line != "" {
			linePoints := strings.Split(line, " -> ")
			for i := 0; i < len(linePoints)-1; i++ {
				startX, startY := getLineCoord(linePoints[i])
				endX, endY := getLineCoord(linePoints[i+1])
				if startY < endY {
					for y := startY; y <= endY; y++ {
						cave[util.SerializeCoordRaw(startX, y)] = Rock
					}
				} else if startY > endY {
					for y := startY; y >= endY; y-- {
						cave[util.SerializeCoordRaw(startX, y)] = Rock
					}
				} else if startX < endX {
					for x := startX; x <= endX; x++ {
						cave[util.SerializeCoordRaw(x, startY)] = Rock
					}
				} else {
					for x := startX; x >= endX; x-- {
						cave[util.SerializeCoordRaw(x, startY)] = Rock
					}
				}

				if startY+1 > bedrock {
					bedrock = startY + 1
				}
				if endY+1 > bedrock {
					bedrock = endY + 1
				}
			}
		}
	}

	return cave, bedrock
}

func (c *Cave) dropSand(bedrock int, bedrockIsFloor bool) (string, bool) {
	x := 500
	y := 0
	moved := true
	for moved {
		if y >= bedrock {
			if bedrockIsFloor {
				moved = false
			} else {
				return "", true
			}
		} else if _, occupied := (*c)[util.SerializeCoordRaw(x, y+1)]; !occupied {
			y++
		} else if _, occupied := (*c)[util.SerializeCoordRaw(x-1, y+1)]; !occupied {
			y++
			x--
		} else if _, occupied := (*c)[util.SerializeCoordRaw(x+1, y+1)]; !occupied {
			y++
			x++
		} else {
			moved = false
		}
	}
	loc := util.SerializeCoordRaw(x, y)
	(*c)[loc] = Sand
	return loc, false
}

func part1(cave Cave, bedrock int) int {
	done := false
	iteration := 0
	for !cave.inputPlugged() && !done {
		_, doneThisRound := cave.dropSand(bedrock, false)
		done = doneThisRound
		iteration++
	}
	return iteration - 1
}

func part2(cave Cave, bedrock int) int {
	iteration := 0
	for !cave.inputPlugged() {
		cave.dropSand(bedrock, true)
		iteration++
	}
	return iteration
}

func getLineCoord(in string) (int, int) {
	coords := strings.Split(in, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return x, y
}

func main() {
	lines := util.FileAsLines("input")
	cave, bedrock := parse(lines)

	part1Res := part1(cave, bedrock)
	fmt.Println(part1Res)
	// We could run this on a clean copy of the cave, but the part1 steps would be done the same way
	part2Res := part2(cave, bedrock)
	fmt.Println(part1Res + part2Res)
}
