package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// To make life easier, hardcode the amount of stacks
const NB_STACK = 9

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
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	// We build the internal data twice to not resort to weird deep copy logic
	stacksPart1 := []Stack{}
	stacksPart2 := []Stack{}

	moves := []Move{}
	for i := 0; i < NB_STACK; i++ {
		stacksPart1 = append(stacksPart1, Stack{})
		stacksPart2 = append(stacksPart2, Stack{})
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
			for i := 0; i < NB_STACK; i++ {
				crate := crates[i]
				if crate != "---" {
					stacksPart1[i].addToBottom(crate)
					stacksPart2[i].addToBottom(crate)
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

	fmt.Println(part1(stacksPart1, moves))
	fmt.Println(part2(stacksPart2, moves))
}
