package main

import "aoc/util"

type Valve struct {
	name string
	rate int
	conn []string
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

func getValveTimeRates(valves []Valve, names []string) map[string]int {
	result := make(map[string]int)
	for _, valve := range valves {
		if util.Contains(names, valve.name) {
			for time := 1; time < 31; time++ {
				key := util.Serialize(valve.name, time)
				result[key] = time * valve.rate
			}
		}
	}
	return result
}
