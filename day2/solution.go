package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Round struct {
	first  string
	second string
}

func winsTo(choice string) string {
	if choice == "A" {
		return "Y"
	} else if choice == "B" {
		return "Z"
	}
	return "X"
}

func drawsTo(choice string) string {
	if choice == "A" {
		return "X"
	} else if choice == "B" {
		return "Y"
	}
	return "Z"
}

func losesTo(choice string) string {
	if choice == "A" {
		return "Z"
	} else if choice == "B" {
		return "X"
	}
	return "Y"
}

func choiceScore(round Round) int {
	if round.second == "X" {
		return 1
	} else if round.second == "Y" {
		return 2
	}
	return 3
}

func battleScore(round Round) int {
	if round.second == winsTo(round.first) {
		return 6
	} else if round.second == losesTo(round.first) {
		return 0
	}
	return 3
}

func getPart1Score(rounds []Round) int {
	score := 0
	for _, round := range rounds {
		score += choiceScore(round) + battleScore(round)
	}
	return score
}

func transformRound(round Round) Round {
	if round.second == "Y" {
		return Round{round.first, drawsTo(round.first)}
	} else if round.second == "Z" {
		return Round{round.first, winsTo(round.first)}
	}

	return Round{round.first, losesTo(round.first)}
}

func getPart2Score(rounds []Round) int {
	score := 0
	for _, round := range rounds {
		updatedRound := transformRound(round)
		score += choiceScore(updatedRound) + battleScore(updatedRound)
	}
	return score
}

func main() {
	lines := util.FileAsLines("input")
	rounds := []Round{}

	for _, line := range lines {
		if line != "" {
			splitLine := strings.Split(line, " ")
			round := Round{splitLine[0], splitLine[1]}
			rounds = append(rounds, round)
		}
	}
	fmt.Println(getPart1Score(rounds))
	fmt.Println(getPart2Score(rounds))
}
