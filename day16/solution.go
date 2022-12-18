package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var valveRegex = regexp.MustCompile("^Valve (.*) has flow rate=(.*); tunnels? leads? to valves? (.*)$")

func parse(lines []string) []Valve {
	valves := make([]Valve, 0, len(lines))
	for _, line := range lines {
		matches := valveRegex.FindStringSubmatch(line)
		rate, _ := strconv.Atoi(matches[2])
		connections := strings.Split(matches[3], ", ")
		valves = append(valves, Valve{matches[1], rate, connections})
	}

	return valves
}

func step(loc string, free int, score int, dist map[string]int, rates map[string]int, left []string) int {
	if len(left) == 0 {
		return score
	}
	if free > 30 {
		return score
	}

	best := 0
	for _, dest := range left {
		activate := free + dist[util.Serialize(loc, dest)] + 1

		newScore := score

		if activate <= 30 {
			newScore += rates[util.Serialize(dest, 30-activate)]
		}

		curr := step(
			dest,
			activate,
			newScore,
			dist, rates,
			util.Remove(left, dest),
		)

		if curr > best {
			best = curr
		}
	}
	return best

}

func part1(valves []Valve) int {
	valveNames := getNonZeroValveNames(valves)
	dist := getValveDistances(valves)
	rates := getValveTimeRates(valves, valveNames)

	return step("AA", 0, 0, dist, rates, valveNames)
}

func isDoomed(best int, free int, loc string, score int, left []string, rates map[string]int, dist map[string]int) bool {
	miracleRun := score
	for _, el := range left {
		activate := free + dist[util.Serialize(loc, el)]
		miracleRun += rates[util.Serialize(el, 26-activate)]
	}
	return miracleRun <= best
}

func stepDuo(loc1 string, free1 int, loc2 string, free2 int, score int, dist map[string]int, rates map[string]int, left []string, knownBest int) int {
	if len(left) == 0 {
		return score
	}
	if free1 > 26 && free2 > 26 {
		return score
	}
	if free1 == free2 {
		if isDoomed(knownBest, free1, loc1, score, left, rates, dist) {
			return knownBest
		}
		best := knownBest
		for _, dest1 := range left {
			for _, dest2 := range util.Remove(left, dest1) {
				activate1 := free1 + dist[util.Serialize(loc1, dest1)] + 1
				activate2 := free2 + dist[util.Serialize(loc2, dest2)] + 1

				newScore := score

				if activate1 <= 26 {
					newScore += rates[util.Serialize(dest1, 26-activate1)]
				}
				if activate2 <= 26 {
					newScore += rates[util.Serialize(dest2, 26-activate2)]
				}

				curr := stepDuo(
					dest1,
					activate1,
					dest2,
					activate2,
					newScore,
					dist, rates,
					util.Remove(left, dest1, dest2),
					best,
				)

				if curr > best {
					best = curr
				}
			}
		}
		return best
	} else if free1 < free2 {
		if isDoomed(knownBest, free1, loc1, score, left, rates, dist) {
			return knownBest
		}
		best := knownBest
		for _, dest := range left {
			activate := free1 + dist[util.Serialize(loc1, dest)] + 1

			newScore := score

			if activate <= 26 {
				newScore += rates[util.Serialize(dest, 26-activate)]
			}

			curr := stepDuo(
				dest,
				activate,
				loc2,
				free2,
				newScore,
				dist, rates,
				util.Remove(left, dest),
				best,
			)

			if curr > best {
				best = curr
			}
		}
		return best
	} else {
		if isDoomed(knownBest, free2, loc2, score, left, rates, dist) {
			return knownBest
		}
		best := knownBest
		for _, dest := range left {
			activate := free2 + dist[util.Serialize(loc2, dest)] + 1

			newScore := score

			if activate <= 26 {
				newScore += rates[util.Serialize(dest, 26-activate)]
			}

			curr := stepDuo(
				loc1,
				free1,
				dest,
				activate,
				newScore,
				dist, rates,
				util.Remove(left, dest),
				best,
			)

			if curr > best {
				best = curr
			}
		}
		return best
	}
}

func part2(valves []Valve) int {
	valveNames := getNonZeroValveNames(valves)
	dist := getValveDistances(valves)
	rates := getValveTimeRates(valves, valveNames)

	return stepDuo("AA", 0, "AA", 0, 0, dist, rates, valveNames, 0)
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")

	valves := parse(lines)
	fmt.Println(part1(valves))
	fmt.Println(part2(valves))
}
