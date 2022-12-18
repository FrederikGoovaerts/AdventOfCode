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

func (s *Sensor) cover(y int, xCovers [][]int) [][]int {
	spread := s.coveredDist - util.Abs(y-s.y)
	if spread < 0 {
		return xCovers
	}
	xStart := s.x - spread
	xEnd := s.x + spread

	result := make([][]int, 0)

	for _, xCover := range xCovers {
		if xCover[0] >= xStart && xCover[1] <= xEnd {
			// fully covered
		} else if xCover[1] < xStart || xCover[0] > xEnd {
			result = append(result, xCover)
		} else if xCover[0] > xStart && xCover[1] > xEnd {
			result = append(result, []int{xEnd + 1, xCover[1]})
		} else if xCover[0] < xStart && xCover[1] < xEnd {
			result = append(result, []int{xCover[0], xStart - 1})
		} else {
			result = append(result, []int{xCover[0], xStart - 1}, []int{xEnd + 1, xCover[1]})
		}
	}

	return result
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return util.Abs(x2-x1) + util.Abs(y2-y1)
}

var inputRegex = regexp.MustCompile("Sensor at x=(.*), y=(.*): closest beacon is at x=(.*), y=(.*)")

func parse(lines []string) ([]Sensor, util.StringSet, int, int) {
	sensors := make([]Sensor, 0)
	beaconSet := make(util.StringSet)

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
	return sensors, beaconSet, smallestXCovered, largestXCovered
}

func part1(sensors []Sensor, beacons util.StringSet, xStart int, xEnd int, y int) int {
	result := 0

	for x := xStart; x <= xEnd; x++ {
		if _, ok := beacons[util.SerializeCoord(x, y)]; ok {
			continue
		}
		for _, sensor := range sensors {
			if sensor.covers(x, y) {
				result++
				break
			}
		}
	}

	return result
}

func part2(sensors []Sensor, size int) int {
	for y := 0; y <= size; y++ {
		uncovered := [][]int{{0, size}}
		for _, sensor := range sensors {
			uncovered = sensor.cover(y, uncovered)

			if len(uncovered) == 0 {
				break
			}
		}

		if len(uncovered) == 1 && uncovered[0][0] == uncovered[0][1] {
			return uncovered[0][0]*4_000_000 + y
		}
	}

	return -1
}

func main() {
	// const part1y = 10
	// const part2Size = 20
	// lines := util.FileAsLines("ex1")

	const part1y = 2_000_000
	const part2Size = 4_000_000
	lines := util.FileAsLines("input")
	sensors, beaconSet, smallestXCovered, largestXCovered := parse(lines)

	fmt.Println(part1(sensors, beaconSet, smallestXCovered, largestXCovered, part1y))
	fmt.Println(part2(sensors, part2Size))
}
