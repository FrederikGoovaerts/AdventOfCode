package main

import (
	"aoc/util"
	"fmt"
	"math"
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

type DuoState struct {
	dest1           string
	timeToDest1Open int
	dest2           string
	timeToDest2Open int
	currentRate     int
	score           int
	timeLeft        int
	unopened        []string
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

func stepDuo(state DuoState, dist map[string]int, rates map[string]int) int {
	if state.timeLeft <= 0 {
		return state.score
	}

	dest1 := state.dest1
	dest2 := state.dest2
	dest1Time := state.timeToDest1Open
	dest2Time := state.timeToDest2Open
	unopened := state.unopened

	if state.timeToDest1Open > 0 && state.timeToDest2Open > 0 {
		change := 1
		if dest1Time > dest2Time {
			change = int(math.Min(float64(dest2Time), float64(state.timeLeft)))
		} else {
			change = int(math.Min(float64(dest1Time), float64(state.timeLeft)))
		}
		return stepDuo(DuoState{
			state.dest1,
			state.timeToDest1Open - change,
			state.dest2,
			state.timeToDest2Open - change,
			state.currentRate,
			state.score + (state.currentRate * change),
			state.timeLeft - change,
			state.unopened,
		}, dist, rates)
	}

	if dest1Time == 0 && dest2Time == 0 {
		if len(unopened) > 1 {
			best := 0
			for _, un1 := range unopened {
				for _, un2 := range util.Remove(unopened, un1) {
					curr := stepDuo(DuoState{
						un1,
						dist[util.Serialize(dest1, un1)],
						un2,
						dist[util.Serialize(dest2, un2)],
						state.currentRate + rates[dest1] + rates[dest2],
						state.score + state.currentRate,
						state.timeLeft - 1,
						util.Remove(state.unopened, un1, un2),
					}, dist, rates)
					if curr > best {
						best = curr
					}
				}
			}
			return best
		} else if len(unopened) == 1 {
			dest := unopened[0]
			first := stepDuo(DuoState{
				IDLE_DEST,
				100,
				dest,
				dist[util.Serialize(dest2, dest)],
				state.currentRate + rates[dest1] + rates[dest2],
				state.score + state.currentRate,
				state.timeLeft - 1,
				[]string{},
			}, dist, rates)

			second := stepDuo(DuoState{
				dest,
				dist[util.Serialize(dest1, dest)],
				IDLE_DEST,
				100,
				state.currentRate + rates[dest1] + rates[dest2],
				state.score + state.currentRate,
				state.timeLeft - 1,
				[]string{},
			}, dist, rates)
			if first > second {
				return first
			} else {
				return second
			}
		} else {
			return stepDuo(DuoState{
				IDLE_DEST,
				100,
				IDLE_DEST,
				100,
				state.currentRate + rates[dest1] + rates[dest2],
				state.score + state.currentRate,
				state.timeLeft - 1,
				state.unopened,
			}, dist, rates)
		}
	} else if dest1Time == 0 { // dest2Time > 0
		if len(unopened) > 0 {
			best := 0
			for _, un := range unopened {
				curr := stepDuo(DuoState{
					un,
					dist[util.Serialize(dest1, un)],
					dest2,
					dest2Time - 1,
					state.currentRate + rates[dest1],
					state.score + state.currentRate,
					state.timeLeft - 1,
					util.Remove(state.unopened, un),
				}, dist, rates)

				if curr > best {
					best = curr
				}
			}
			return best
		} else {
			return stepDuo(DuoState{
				IDLE_DEST,
				100,
				dest2,
				dest2Time - 1,
				state.currentRate + rates[dest1],
				state.score + state.currentRate,
				state.timeLeft - 1,
				state.unopened,
			}, dist, rates)
		}
	} else { // dest1Time > 0, dest2Time == 0
		if len(unopened) > 0 {
			best := 0
			for _, un := range unopened {
				curr := stepDuo(DuoState{
					dest1,
					dest1Time - 1,
					un,
					dist[util.Serialize(dest2, un)],
					state.currentRate + rates[dest2],
					state.score + state.currentRate,
					state.timeLeft - 1,
					util.Remove(state.unopened, un),
				}, dist, rates)

				if curr > best {
					best = curr
				}
			}
			return best
		} else {
			return stepDuo(DuoState{
				dest1,
				dest1Time - 1,
				IDLE_DEST,
				100,
				state.currentRate + rates[dest2],
				state.score + state.currentRate,
				state.timeLeft - 1,
				state.unopened,
			}, dist, rates)
		}
	}
}

func part2(valves []Valve) int {
	// result := 0

	dist := getValveDistances(valves)
	valveNames := getNonZeroValveNames(valves)
	valveRates := getNonZeroValveRates(valves)

	// states := []DuoState{{"AA", 0, "AA", 0, 0, 0, 27, valveNames}}

	// for len(states) > 0 {
	// 	curr := states[0]
	// 	states = states[1:]
	// 	for _, steppedState := range stepDuo(curr, dist, valveRates) {
	// 		if steppedState.timeLeft > 0 {
	// 			states = append([]DuoState{steppedState}, states...)
	// 		} else if steppedState.score > result {
	// 			result = steppedState.score
	// 			fmt.Println(result)
	// 		}
	// 	}

	// }

	// return result
	return stepDuo(DuoState{"AA", 0, "AA", 0, 0, 0, 27, valveNames}, dist, valveRates)
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")

	valves := parse(lines)
	fmt.Println(part1(valves))
	fmt.Println(part2(valves))
}

// 2469 too low
