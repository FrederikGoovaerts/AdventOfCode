package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	crates []string
}

func (s *Stack) addToTop(c []string) {
	s.crates = append(s.crates, c...)
}

func (s *Stack) removeFromTop(amount int) []string {
	removed := s.crates[len(s.crates)-amount:]
	s.crates = s.crates[:len(s.crates)-amount]
	return removed
}

func (s *Stack) addToBottom(c string) {
	s.crates = append([]string{c}, s.crates...)
}

type Move struct {
	amount int
	from   int
	to     int
}

// To make life easier, hardcode the amount of stacks
func parse(lines []string, nbStacks int) ([]Stack, []Move) {
	stacks := []Stack{}
	moves := []Move{}

	for i := 0; i < nbStacks; i++ {
		stacks = append(stacks, Stack{})
	}

	readMovesMode := false
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.HasPrefix(line, " 1") {
			readMovesMode = true
		} else if !readMovesMode {
			startSpacingRegex := regexp.MustCompile(`^    `)
			otherSpacingRegex := regexp.MustCompile(`    `)
			startFixed := startSpacingRegex.ReplaceAllString(line, "--- ")
			fixed := otherSpacingRegex.ReplaceAllString(startFixed, " ---")

			crates := strings.Split(fixed, " ")
			for i := 0; i < nbStacks; i++ {
				crate := crates[i]
				if crate != "---" {
					stacks[i].addToBottom(crate)
				}
			}
		} else {
			r, _ := regexp.Compile("move (.*) from (.*) to (.*)")
			match := r.FindStringSubmatch(line)
			amount, _ := strconv.Atoi(match[1])
			from, _ := strconv.Atoi(match[2])
			to, _ := strconv.Atoi(match[3])
			moves = append(moves, Move{amount, from, to})
		}
	}

	return stacks, moves
}

func part1(stacks []Stack, moves []Move) string {
	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			moved := stacks[move.from-1].removeFromTop(1)
			stacks[move.to-1].addToTop(moved)
		}
	}
	result := ""
	for _, stack := range stacks {
		result += stack.crates[len(stack.crates)-1][1:2]
	}
	return result
}

func part2(stacks []Stack, moves []Move) string {
	for _, move := range moves {
		moved := stacks[move.from-1].removeFromTop(move.amount)
		stacks[move.to-1].addToTop(moved)
	}
	result := ""
	for _, stack := range stacks {
		result += stack.crates[len(stack.crates)-1][1:2]
	}
	return result
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")

	stacksPart1, moves := parse(lines, 9)
	fmt.Println(part1(stacksPart1, moves))

	stacksPart2, _ := parse(lines, 9)
	fmt.Println(part2(stacksPart2, moves))
}
