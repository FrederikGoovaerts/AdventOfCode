package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

var bluePrintRegex = regexp.MustCompile("^Blueprint (.*): Each ore robot costs (.*) ore. Each clay robot costs (.*) ore. Each obsidian robot costs (.*) ore and (.*) clay. Each geode robot costs (.*) ore and (.*) obsidian.$")

func parse(lines []string) []Blueprint {
	blueprints := make([]Blueprint, len(lines))
	for index, line := range lines {
		matches := bluePrintRegex.FindStringSubmatch(line)
		id, _ := strconv.Atoi(matches[1])
		oreOre, _ := strconv.Atoi(matches[2])
		clayOre, _ := strconv.Atoi(matches[3])
		obsOre, _ := strconv.Atoi(matches[4])
		obsClay, _ := strconv.Atoi(matches[5])
		geoOre, _ := strconv.Atoi(matches[6])
		geoObs, _ := strconv.Atoi(matches[7])
		blueprints[index] = Blueprint{
			id,
			oreOre,
			clayOre,
			obsOre,
			obsClay,
			geoOre,
			geoObs,
		}
	}

	return blueprints
}

func simulate(blueprint Blueprint, path []Robot, timeLimit int) int {
	buildCounter := 0

	oreRobot := 1
	clayRobot := 0
	obsRobot := 0
	geoRobot := 0
	ore := 0
	clay := 0
	obs := 0
	geo := 0
	for t := 1; t <= timeLimit; t++ {
		willCreate := false
		if buildCounter < len(path) {
			robotToCreate := path[buildCounter]
			if blueprint.canCreate(robotToCreate, ore, clay, obs) {
				willCreate = true
			}
		}

		ore += oreRobot
		clay += clayRobot
		obs += obsRobot
		geo += geoRobot

		if willCreate {
			buildCounter++
			switch path[buildCounter-1] {
			case OreRobot:
				oreRobot++
				ore -= blueprint.oreRobotOre
			case ClayRobot:
				clayRobot++
				ore -= blueprint.clayRobotOre
			case ObsidianRobot:
				obsRobot++
				clay -= blueprint.obsRobotClay
				ore -= blueprint.obsRobotOre
			case GeodeRobot:
				geoRobot++
				obs -= blueprint.geoRobotObs
				ore -= blueprint.geoRobotOre
			}
		}
	}
	if buildCounter < len(path) {
		return -1
	} else {
		return geo
	}
}

var advancedRobots = []Robot{ClayRobot, ObsidianRobot, GeodeRobot}

// Expertly crafted with my intellectual insight
var startingPaths = [][]Robot{
	{},
	{OreRobot},
	{OreRobot, OreRobot},
	{ClayRobot, OreRobot},
	{OreRobot, ClayRobot, OreRobot},
	{OreRobot, OreRobot, OreRobot},
	{OreRobot, OreRobot, OreRobot, OreRobot},
	{OreRobot, ClayRobot, OreRobot, OreRobot},
	{OreRobot, OreRobot, ClayRobot, OreRobot},
}

func getBest(blueprint Blueprint, path []Robot, time int) int {
	score := simulate(blueprint, path, time)

	if score == -1 {
		return score
	}
	newPath := make([]Robot, len(path)+1)
	copy(newPath, path)
	for _, robot := range advancedRobots {
		newPath[len(path)] = robot
		newScore := getBest(blueprint, newPath, time)
		if newScore > score {
			score = newScore
		}
	}
	return score
}

func part1(blueprints []Blueprint) int {
	score := make([]int, len(blueprints))

	for index, blueprint := range blueprints {
		for _, p := range startingPaths {
			score[index] = util.MaxInt(getBest(blueprint, p, 24), score[index])
		}
		score[index] *= blueprint.id
	}

	return util.Sum(score)
}

func part2(blueprints []Blueprint) int {
	score := make([]int, len(blueprints))

	for index, blueprint := range blueprints {
		for _, p := range startingPaths {
			score[index] = util.MaxInt(getBest(blueprint, p, 32), score[index])
		}
	}

	return util.Product(score)
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")
	blueprints := parse(lines)

	fmt.Println(part1(blueprints))
	fmt.Println(part2(blueprints[:3]))
}
