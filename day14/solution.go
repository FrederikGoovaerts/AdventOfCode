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

func serialize(x int, y int) string {
	return fmt.Sprint(x) + " " + fmt.Sprint(y)
}

type Cave map[string]CaveFilling

func (c *Cave) inputPlugged() bool {
	return (*c)[serialize(500, 0)] == Sand
}
func (c *Cave) dropSand(bedrock int) (string, bool) {
	x := 500
	y := 0
	moved := true
	for moved {
		if y >= bedrock {
			return "", true
		}
		if _, occupied := (*c)[serialize(x, y+1)]; !occupied {
			y++
		} else if _, occupied := (*c)[serialize(x-1, y+1)]; !occupied {
			y++
			x--
		} else if _, occupied := (*c)[serialize(x+1, y+1)]; !occupied {
			y++
			x++
		} else {
			moved = false
		}
	}
	loc := serialize(x, y)
	(*c)[loc] = Sand
	return loc, false
}

func part1(cave Cave, bedrock int) int {
	done := false
	iteration := 0
	for !cave.inputPlugged() && !done {
		_, doneThisRound := cave.dropSand(bedrock)
		done = doneThisRound
		iteration++
	}
	return iteration - 1
}

func getLineCoord(in string) (int, int) {
	coords := strings.Split(in, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return x, y
}

func main() {
	lines := util.FileAsLines("input")
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
						cave[serialize(startX, y)] = Rock
					}
				} else if startY > endY {
					for y := startY; y >= endY; y-- {
						cave[serialize(startX, y)] = Rock
					}
				} else if startX < endX {
					for x := startX; x <= endX; x++ {
						cave[serialize(x, startY)] = Rock
					}
				} else {
					for x := startX; x >= endX; x-- {
						cave[serialize(x, startY)] = Rock
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

	fmt.Println(part1(cave, bedrock))
}
