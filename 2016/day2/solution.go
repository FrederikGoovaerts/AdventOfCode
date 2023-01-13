package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Move struct {
	xOff int
	yOff int
}

var Up Move = Move{0, -1}
var Down Move = Move{0, 1}
var Right Move = Move{1, 0}
var Left Move = Move{-1, 0}

func getSymbolAtBasic(x int, y int) string {
	return fmt.Sprint(x + 2 + ((y + 1) * 3))
}

func getSymbolAtAdvanced(x int, y int) string {
	if x == -2 {
		return "5"
	} else if x == -1 {
		if y == -1 {
			return "2"
		} else if y == 0 {
			return "6"
		} else {
			return "A"
		}
	} else if x == 0 {
		if y == -2 {
			return "1"
		} else if y == -1 {
			return "3"
		} else if y == 0 {
			return "7"
		} else if y == 1 {
			return "B"
		} else {
			return "D"
		}
	} else if x == 1 {
		if y == -1 {
			return "4"
		} else if y == 0 {
			return "8"
		} else {
			return "C"
		}
	} else {
		return "9"
	}
}

func parse(lines []string) [][]Move {
	moves := make([][]Move, 0)

	for _, line := range lines {
		move := make([]Move, 0)
		for _, dir := range strings.Split(line, "") {
			switch dir {
			case "U":
				move = append(move, Up)
			case "D":
				move = append(move, Down)
			case "L":
				move = append(move, Left)
			case "R":
				move = append(move, Right)
			}
		}
		moves = append(moves, move)
	}

	return moves
}

func part1(moves [][]Move) string {
	result := ""
	x, y := 0, 0
	for _, move := range moves {
		for _, m := range move {
			newX := x + m.xOff
			newY := y + m.yOff
			if util.Abs(newX) <= 1 && util.Abs(newY) <= 1 {
				x = newX
				y = newY
			}
		}
		result += getSymbolAtBasic(x, y)
	}
	return result
}

func part2(moves [][]Move) string {
	result := ""
	x, y := -2, 0
	for _, move := range moves {
		for _, m := range move {
			newX := x + m.xOff
			newY := y + m.yOff
			if util.Abs(newX)+util.Abs(newY) <= 2 {
				x = newX
				y = newY
			}
		}
		result += getSymbolAtAdvanced(x, y)
	}
	return result
}

func main() {
	// input := util.FileAsLines("ex1")
	input := util.FileAsLines("input")

	moves := parse(input)

	fmt.Println(part1(moves))
	fmt.Println(part2(moves))
}
