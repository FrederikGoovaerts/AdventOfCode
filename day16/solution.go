package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	name string
	rate int
	conn []string
}

var valveRegex = regexp.MustCompile("^Valve (.*) has flow rate=(.*); tunnels? leads? to valves? (.*)$")

const IDLE_DEST = "IDLE"

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

type State struct {
	dest           string
	timeToDestOpen int
	currentRate    int
	score          int
	timeLeft       int
	unopened       []string
}

func getNonZeroValveNames(valves []Valve) []string {
	result := make([]string, 0)
	for _, valve := range valves {
		if valve.rate > 0 {
			result = append(result, valve.name)
		}
	}
	return result
}

func getNonZeroValveRates(valves []Valve) map[string]int {
	result := make(map[string]int)
	for _, valve := range valves {
		if valve.rate > 0 {
			result[valve.name] = valve.rate
		}
	}
	return result
}

func getValveDistances(valves []Valve) map[string]int {
	dist := make(map[string]int, len(valves)*len(valves))
	for _, v1 := range valves {
		for _, v2 := range valves {
			if v1.name == v2.name {
				dist[util.Serialize(v1.name, v2.name)] = 0
			} else if util.Contains(v1.conn, v2.name) {
				dist[util.Serialize(v1.name, v2.name)] = 1
			} else {
				dist[util.Serialize(v1.name, v2.name)] = 200 // Arbitrary number above 30
			}
		}
	}

	for _, v1 := range valves {
		for _, v2 := range valves {
			for _, v3 := range valves {
				d21 := dist[util.Serialize(v2.name, v1.name)]
				d13 := dist[util.Serialize(v1.name, v3.name)]
				d23 := dist[util.Serialize(v2.name, v3.name)]
				if d23 > d21+d13 {
					dist[util.Serialize(v2.name, v3.name)] = d21 + d13
				}
			}
		}
	}
	return dist
}

func step(state State, dist map[string]int, rates map[string]int) []State {
	result := make([]State, 0, len(state.unopened))
	if state.timeToDestOpen == 0 {
		if len(state.unopened) > 0 && state.dest != IDLE_DEST {
			for _, next := range state.unopened {
				timeToNextOpen := dist[util.Serialize(state.dest, next)]

				result = append(result, State{
					next,
					timeToNextOpen,
					state.currentRate + rates[state.dest],
					state.score + state.currentRate,
					state.timeLeft - 1,
					util.Remove(state.unopened, next),
				})
			}
		} else {
			return []State{{IDLE_DEST, 100, state.currentRate + rates[state.dest], state.score + state.currentRate, state.timeLeft - 1, state.unopened}}
		}
	} else {
		return []State{{state.dest, state.timeToDestOpen - 1, state.currentRate, state.score + state.currentRate, state.timeLeft - 1, state.unopened}}
	}

	return result
}

func part1(valves []Valve) int {
	result := 0

	dist := getValveDistances(valves)
	valveNames := getNonZeroValveNames(valves)
	valveRates := getNonZeroValveRates(valves)

	states := []State{{"AA", 0, 0, 0, 31, valveNames}}

	for len(states) > 0 {
		curr := states[0]
		states = states[1:]
		for _, steppedState := range step(curr, dist, valveRates) {
			if steppedState.timeLeft > 0 {
				states = append(states, steppedState)
			} else if steppedState.score > result {
				result = steppedState.score
			}
		}

	}

	return result
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")

	valves := parse(lines)
	fmt.Println(part1(valves))
}
