package main

import (
	"aoc/util"
	"fmt"
	"regexp"
)

type DistMap map[string]int

func (d *DistMap) add(first, second string, dist int) {
	(*d)[util.Serialize(first, second)] = dist
}

func (d *DistMap) get(first, second string) int {
	return (*d)[util.Serialize(first, second)]
}

var inputRegex = regexp.MustCompile("^(.*) would (lose|gain) ([0-9]*) happiness units by sitting next to (.*).$")

func parse(lines []string) ([]string, DistMap) {
	nameSet := make(util.StringSet)
	distances := make(DistMap)

	for _, line := range lines {
		matches := inputRegex.FindStringSubmatch(line)
		dist := util.StringToInt(matches[3])
		if matches[2] == "lose" {
			dist = -dist
		}
		distances.add(matches[1], matches[4], dist)
		nameSet.Add(matches[1])
	}

	names := make([]string, 0, len(nameSet))
	for name := range nameSet {
		names = append(names, name)
	}
	return names, distances
}

type Parameters struct {
	startLoc string
	addSelf  bool
	dists    DistMap
}

func findHighest(currDist int, currLoc string, remaining []string, currHighest int, p Parameters) int {
	if len(remaining) == 0 {
		if p.addSelf {
			return util.MaxInt(currHighest, currDist)
		} else {
			closed := currDist + p.dists.get(currLoc, p.startLoc) + p.dists.get(p.startLoc, currLoc)
			return util.MaxInt(currHighest, closed)
		}
	}
	best := currHighest
	for _, next := range remaining {
		mutualDists := p.dists.get(currLoc, next) + p.dists.get(next, currLoc)
		best = util.MaxInt(best, findHighest(currDist+mutualDists, next, util.Remove(remaining, next), best, p))
	}
	return best
}

func part1(names []string, dists DistMap) int {
	best := 0
	for _, name := range names {
		best = util.MaxInt(best, findHighest(0, name, util.Remove(names, name), best, Parameters{name, false, dists}))
	}
	return best
}

func part2(names []string, dists DistMap) int {
	best := 0
	for _, name := range names {
		best = util.MaxInt(best, findHighest(0, name, util.Remove(names, name), best, Parameters{name, true, dists}))
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
