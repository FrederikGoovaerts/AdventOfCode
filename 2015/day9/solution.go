package main

import (
	"aoc/util"
	"fmt"
	"math"
	"regexp"
)

type DistMap map[string]int

func (d *DistMap) add(first, second string, dist int) {
	(*d)[util.Serialize(first, second)] = dist
	(*d)[util.Serialize(second, first)] = dist
}

func (d *DistMap) get(first, second string) int {
	return (*d)[util.Serialize(first, second)]
}

var inputRegex = regexp.MustCompile("^(.*) to (.*) = (.*)$")

func parse(lines []string) ([]string, DistMap) {
	nameSet := make(util.StringSet)
	distances := make(DistMap)

	for _, line := range lines {
		matches := inputRegex.FindStringSubmatch(line)
		dist := util.StringToInt(matches[3])
		distances.add(matches[1], matches[2], dist)
		nameSet.Add(matches[1])
		nameSet.Add(matches[2])
	}

	names := make([]string, 0, len(nameSet))
	for name := range nameSet {
		names = append(names, name)
	}
	return names, distances
}

func findShortest(currDist int, currLoc string, remaining []string, dists DistMap, currBest int) int {
	if currDist > currBest {
		return currBest
	}
	if len(remaining) == 0 {
		return util.MinInt(currBest, currDist)
	}
	best := currBest
	for _, next := range remaining {
		distTo := dists.get(currLoc, next)
		best = util.MinInt(best, findShortest(currDist+distTo, next, util.Remove(remaining, next), dists, best))
	}
	return best
}

func part1(names []string, dists DistMap) int {
	best := math.MaxInt
	for _, name := range names {
		best = util.MinInt(best, findShortest(0, name, util.Remove(names, name), dists, best))
	}
	return best
}

func findLongest(currDist int, currLoc string, remaining []string, dists DistMap, currBest int) int {
	if len(remaining) == 0 {
		return util.MaxInt(currBest, currDist)
	}
	best := currBest
	for _, next := range remaining {
		distTo := dists.get(currLoc, next)
		best = util.MaxInt(best, findLongest(currDist+distTo, next, util.Remove(remaining, next), dists, best))
	}
	return best
}

func part2(names []string, dists DistMap) int {
	best := 0
	for _, name := range names {
		best = util.MaxInt(best, findLongest(0, name, util.Remove(names, name), dists, best))
	}
	return best
}

func main() {
	// input := util.FileAsLines("ex1")
	input := util.FileAsLines("input")
	names, distances := parse(input)

	fmt.Println(part1(names, distances))
	fmt.Println(part2(names, distances))
}
