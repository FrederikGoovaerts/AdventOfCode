package main

import (
	"aoc/util"
	"fmt"
	"math"
	"strings"
)

func parse(lines []string) util.StringSet {
	cubes := make(util.StringSet)

	for _, line := range lines {
		serCoord3 := strings.ReplaceAll(line, ",", " ")
		cubes[serCoord3] = util.EMPTY_STRUCT
	}

	return cubes
}

func part1(cubes util.StringSet) int {
	result := 0

	for serCoord3 := range cubes {
		result += 6
		cube := util.DeserializeCoord3(serCoord3)
		for _, neigh := range cube.GetNeighbors() {
			_, present := cubes[util.SerializeCoord3(neigh)]
			if present {
				result--
			}
		}
	}

	return result
}

type Bounds struct {
	minX int
	minY int
	minZ int
	maxX int
	maxY int
	maxZ int
}

func (b *Bounds) isOutOfBounds(cube util.Coord3) bool {
	return cube.X < b.minX || cube.X > b.maxX || cube.Y < b.minY || cube.Y > b.maxY || cube.Z < b.minZ || cube.Z > b.maxZ
}

func (b *Bounds) surface() int {
	// The coordinates represent squares with a size of 1x1x1, not discrete points, hence the +1's
	return 2 * ((b.maxX-b.minX+1)*(b.maxY-b.minY+1) + (b.maxX-b.minX+1)*(b.maxZ-b.minZ+1) + (b.maxZ-b.minZ+1)*(b.maxY-b.minY+1))
}

func floodFillOuter(coord util.Coord3, bounds Bounds, droplet util.StringSet) util.StringSet {
	filled := make(util.StringSet)
	filled[util.SerializeCoord3(coord)] = util.EMPTY_STRUCT
	curr := coord
	next := []util.Coord3{coord}

	for len(next) > 0 {
		curr, next = next[0], next[1:]
		for _, n := range curr.GetNeighbors() {
			serN := util.SerializeCoord3(n)
			_, alreadyFilled := filled[serN]
			_, inDroplet := droplet[serN]
			if !bounds.isOutOfBounds(n) && !alreadyFilled && !inDroplet {
				filled[serN] = util.EMPTY_STRUCT
				next = append(next, n)
			}
		}
	}

	return filled
}

func part2(cubes util.StringSet) int {
	bounds := Bounds{
		math.MaxInt,
		math.MaxInt,
		math.MaxInt,
		math.MinInt,
		math.MinInt,
		math.MinInt,
	}

	for serCoord3 := range cubes {
		cube := util.DeserializeCoord3(serCoord3)
		bounds.minX = util.MinInt(bounds.minX, cube.X)
		bounds.minY = util.MinInt(bounds.minY, cube.Y)
		bounds.minZ = util.MinInt(bounds.minZ, cube.Z)
		bounds.maxX = util.MaxInt(bounds.maxX, cube.X)
		bounds.maxY = util.MaxInt(bounds.maxY, cube.Y)
		bounds.maxZ = util.MaxInt(bounds.maxZ, cube.Z)
	}

	bounds.minX--
	bounds.minY--
	bounds.minZ--
	bounds.maxX++
	bounds.maxY++
	bounds.maxZ++

	outer := floodFillOuter(util.Coord3{X: bounds.minX, Y: bounds.minY, Z: bounds.minZ}, bounds, cubes)

	result := part1(outer) - bounds.surface()

	return result
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")
	cubes := parse(lines)

	fmt.Println(part1(cubes))
	fmt.Println(part2(cubes))
}
