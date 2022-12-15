package main

import (
	"aoc/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Sensor struct {
	x           int
	y           int
	coveredDist int
}

func (s *Sensor) covers(x int, y int) bool {
	return manhattanDistance(x, y, s.x, s.y) <= s.coveredDist
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return util.Abs(x2-x1) + util.Abs(y2-y1)
}

// const PART1_Y = 10

const PART1_Y = 2_000_000

func part1(sensors []Sensor, beacons map[string]struct{}, xStart int, xEnd int) int {
	result := 0
	fmt.Println(xStart, xEnd, beacons)

	for x := xStart; x <= xEnd; x++ {
		if _, ok := beacons[util.SerializeCoord(x, PART1_Y)]; ok {
			continue
		}
		for _, sensor := range sensors {
			if sensor.covers(x, PART1_Y) {
				result++
				break
			}
		}
	}
	fmt.Println()

	return result
}

var inputRegex = regexp.MustCompile("Sensor at x=(.*), y=(.*): closest beacon is at x=(.*), y=(.*)")

func main() {
	lines := util.FileAsLines("input")

	sensors := make([]Sensor, 0)
	beaconSet := make(map[string]struct{})

	smallestXCovered := math.MaxInt
	largestXCovered := math.MinInt

	for _, line := range lines {
		if line != "" {

			match := inputRegex.FindStringSubmatch(line)
			sensorX, _ := strconv.Atoi(match[1])
			sensorY, _ := strconv.Atoi(match[2])
			beaconX, _ := strconv.Atoi(match[3])
			beaconY, _ := strconv.Atoi(match[4])
			distance := manhattanDistance(sensorX, sensorY, beaconX, beaconY)

			sensors = append(sensors, Sensor{sensorX, sensorY, distance})
			beaconSet[util.SerializeCoord(beaconX, beaconY)] = util.EMPTY_STRUCT

			if sensorX-distance < smallestXCovered {
				smallestXCovered = sensorX - distance
			}

			if sensorX+distance > largestXCovered {
				largestXCovered = sensorX + distance
			}
		}
	}

	fmt.Println(part1(sensors, beaconSet, smallestXCovered, largestXCovered))
}
