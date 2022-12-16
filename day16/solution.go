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
	valve    string
	timeLeft int
	score    int
	unopened []string
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

func step(state State, dist map[string]int, rates map[string]int) []State {
	result := make([]State, 0, len(state.unopened))

	for _, next := range state.unopened {
		newTimeLeft := state.timeLeft - (dist[util.Serialize(state.valve, next)] + 1) // distance traveled in minutes and one extra minute to open the valve
		if newTimeLeft >= 0 {
			rate := rates[next]
			score := state.score + newTimeLeft*rate
			unopened := util.Remove(state.unopened, next)
			if len(unopened) == 0 {
				result = append(result, State{next, 0, score, unopened})
			} else {
				result = append(result, State{next, newTimeLeft, score, unopened})
			}
		} else {
			result = append(result, State{next, 0, state.score, state.unopened})
		}
	}

	return result
}

func part1(valves []Valve) int {
	// baby's first Floyd-Warshall
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

	result := 0

	valveNames := getNonZeroValveNames(valves)
	valveRates := getNonZeroValveRates(valves)
	states := []State{{"AA", 30, 0, valveNames}}

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
