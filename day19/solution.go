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

func run(bl Blueprint, inv Inventory, minute int) int {
	if minute == 24 {
		minedInv := inv.mine()
		return bl.id * minedInv.geode
	}

	craftable := inv.shouldCreate(bl)
	minedInv := inv.mine()

	bestQuality := 0
	for _, craftableRobot := range craftable {
		newInv := minedInv.makeRobot(bl, craftableRobot)
		bestQuality = util.MaxInt(bestQuality, run(bl, newInv, minute+1))
	}

	return bestQuality
}

func part1(blueprints []Blueprint) int {
	initialInv := Inventory{}
	initialInv.oreRobot = 1

	qualityLevels := make([]int, len(blueprints))

	for index, blueprint := range blueprints {
		qualityLevels[index] = run(blueprint, initialInv, 1)
	}

	return util.Sum(qualityLevels)
}

func part2(blueprints []Blueprint) int {
	return 0
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")
	blueprints := parse(lines)

	fmt.Println(part1(blueprints))
	fmt.Println(part2(blueprints))
}
