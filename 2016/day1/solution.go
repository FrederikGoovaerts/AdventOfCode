package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Dir int

const (
	N Dir = iota
	E
	S
	W
)

func (d Dir) doTurn(turn rune) Dir {
	if turn == 'R' {
		switch d {
		case N:
			return E
		case E:
			return S
		case S:
			return W
		case W:
			return N
		}
	} else {
		switch d {
		case N:
			return W
		case E:
			return N
		case S:
			return E
		case W:
			return S
		}
	}
	panic("Not a valid turn")
}

type Move struct {
	turn     rune
	distance int
}

func parse(input string) []Move {
	moves := make([]Move, 0)

	for _, m := range strings.Split(input, ", ") {
		moves = append(moves, Move{rune(m[0]), util.StringToInt(m[1:])})
	}

	return moves
}

func part1(moves []Move) int {
	dir := N
	x := 0
	y := 0

	for _, m := range moves {
		dir = dir.doTurn(m.turn)
		switch dir {
		case N:
			x += m.distance
		case E:
			y += m.distance
		case S:
			x -= m.distance
		case W:
			y -= m.distance
		}
	}

	return util.Abs(x) + util.Abs(y)
}

func part2(moves []Move) int {
	visited := make(util.StringSet)
	visited["0 0"] = util.EMPTY_STRUCT
	dir := N
	x := 0
	y := 0

	for _, m := range moves {
		dir = dir.doTurn(m.turn)
		switch dir {
		case N:
			for i := 1; i <= m.distance; i++ {
				newX := x + i
				if visited.Has(util.SerializeCoordRaw(newX, y)) {
					return util.Abs(newX) + util.Abs(y)
				}
				visited.Add(util.SerializeCoordRaw(newX, y))
			}
			x += m.distance
		case E:
			for i := 1; i <= m.distance; i++ {
				newY := y + i
				if visited.Has(util.SerializeCoordRaw(x, newY)) {
					return util.Abs(x) + util.Abs(y)
				}
				visited.Add(util.SerializeCoordRaw(x, newY))
			}
			y += m.distance
		case S:
			for i := 1; i <= m.distance; i++ {
				newX := x - i
				if visited.Has(util.SerializeCoordRaw(newX, y)) {
					return util.Abs(newX) + util.Abs(y)
				}
				visited.Add(util.SerializeCoordRaw(newX, y))
			}
			x -= m.distance
		case W:
			for i := 1; i <= m.distance; i++ {
				newY := y - i
				if visited.Has(util.SerializeCoordRaw(x, newY)) {
					return util.Abs(x) + util.Abs(y)
				}
				visited.Add(util.SerializeCoordRaw(x, newY))
			}
			y -= m.distance
		}
	}

	return -1
}

func main() {
	// input := util.FileAsString("ex2")
	// input := util.FileAsString("ex3")
	input := util.FileAsString("input")

	moves := parse(input)

	fmt.Println(part1(moves))
	fmt.Println(part2(moves))
}
