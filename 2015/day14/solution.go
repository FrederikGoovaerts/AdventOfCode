package main

import (
	"aoc/util"
	"fmt"
	"regexp"
)

type Reindeer struct {
	name     string
	runTime  int
	restTime int
	speed    int
}

func (r Reindeer) getLoopTime() int {
	return r.runTime + r.restTime
}

func (r Reindeer) getCycleDist() int {
	return r.runTime * r.speed
}

var inputRegex = regexp.MustCompile("^(.*) can fly ([0-9]*) km/s for ([0-9]*) seconds, but then must rest for ([0-9]*) seconds.$")

func parse(lines []string) []Reindeer {
	reindeer := make([]Reindeer, 0, len(lines))

	for _, line := range lines {
		matches := inputRegex.FindStringSubmatch(line)
		speed := util.StringToInt(matches[2])
		runTime := util.StringToInt(matches[3])
		restTime := util.StringToInt(matches[4])

		reindeer = append(reindeer, Reindeer{matches[1], runTime, restTime, speed})
	}

	return reindeer
}

func getDistAfter(time int, r Reindeer) int {
	totalDist := 0

	fullCycles := time / r.getLoopTime()
	totalDist += fullCycles * r.getCycleDist()

	remainingTime := time % r.getLoopTime()
	if remainingTime >= r.runTime {
		totalDist += r.getCycleDist()
	} else {
		totalDist += remainingTime * r.speed
	}

	return totalDist
}

func part1(reindeer []Reindeer, time int) int {
	best := 0

	for _, r := range reindeer {
		dist := getDistAfter(time, r)
		// fmt.Println(dist, r.name)
		best = util.MaxInt(best, dist)
	}

	return best
}

func part2(reindeer []Reindeer, time int) int {
	dist := make(map[string]int)
	points := make(map[string]int)

	for _, r := range reindeer {
		dist[r.name] = 0
		points[r.name] = 0
	}

	for i := 0; i < time; i++ {
		bestDist := 0
		for _, r := range reindeer {
			isRunning := i%r.getLoopTime() < r.runTime
			if isRunning {
				dist[r.name] += r.speed
			}
			if dist[r.name] > bestDist {
				bestDist = dist[r.name]
			}
		}
		for k, v := range dist {
			if v == bestDist {
				points[k]++
			}
		}
	}

	best := 0

	for _, v := range points {
		best = util.MaxInt(best, v)
	}

	return best
}

func main() {
	// input, time := util.FileAsLines("ex1"), 1000
	input, time := util.FileAsLines("input"), 2503
	reindeer := parse(input)

	fmt.Println(part1(reindeer, time))
	fmt.Println(part2(reindeer, time))
}
