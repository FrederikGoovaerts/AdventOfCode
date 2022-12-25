package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Turn string

const (
	Left  Turn = "L"
	Right Turn = "R"
)

type Move struct {
	isTurn bool
	turn   Turn
	dist   int
}

type Path int

const (
	Clear Path = iota
	Wall
)

type Map struct {
	xMax int
	yMax int
	cont map[string]Path
}

func (m Map) visualize() {
	for y := 0; y <= m.yMax; y++ {
		for x := 0; x <= m.xMax; x++ {
			res, present := m.cont[util.SerializeCoordRaw(x, y)]
			if present {
				if res == Clear {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (m Map) getStartingX() int {
	for x := 0; x <= m.xMax; x++ {
		res, present := m.cont[util.SerializeCoordRaw(x, 0)]
		if present && res == Clear {
			return x
		}
	}
	return -1
}

func (m Map) has(x, y int) bool {
	_, present := m.cont[util.Serialize(x, y)]
	return present
}

func (m Map) get(x, y int) Path {
	return m.cont[util.Serialize(x, y)]
}

func (m Map) getNext(currX, currY int, dir Direction) (int, int, Path) {
	if dir == North {
		res, present := m.cont[util.SerializeCoordRaw(currX, currY-1)]
		if present {
			return currX, currY - 1, res
		}
		for y := m.yMax; y > currY; y-- {
			res, present := m.cont[util.SerializeCoordRaw(currX, y)]
			if present {
				return currX, y, res
			}
		}
	} else if dir == East {
		res, present := m.cont[util.SerializeCoordRaw(currX+1, currY)]
		if present {

			return currX + 1, currY, res
		}
		for x := 0; x < currX; x++ {
			res, present := m.cont[util.SerializeCoordRaw(x, currY)]
			if present {
				return x, currY, res
			}
		}
	} else if dir == South {
		res, present := m.cont[util.SerializeCoordRaw(currX, currY+1)]
		if present {
			return currX, currY + 1, res
		}
		for y := 0; y < currY; y++ {
			res, present := m.cont[util.SerializeCoordRaw(currX, y)]
			if present {
				return currX, y, res
			}
		}
	} else if dir == West {
		res, present := m.cont[util.SerializeCoordRaw(currX-1, currY)]
		if present {
			return currX - 1, currY, res
		}
		for x := m.xMax; x > currX; x-- {
			res, present := m.cont[util.SerializeCoordRaw(x, currY)]
			if present {
				return x, currY, res
			}
		}
	}
	panic("Went out of bounds")
}

func (m Map) getNextWithEdges(currX, currY int, dir Direction, e EdgeMap) (int, int, Direction, Path) {
	if e.has(currX, currY, dir) {
		newX, newY, newDir := e.getNext(currX, currY, dir)
		path := m.cont[util.SerializeCoordRaw(newX, newY)]
		return newX, newY, newDir, path
	}

	if dir == North {
		res := m.cont[util.SerializeCoordRaw(currX, currY-1)]
		return currX, currY - 1, dir, res
	} else if dir == East {
		res := m.cont[util.SerializeCoordRaw(currX+1, currY)]
		return currX + 1, currY, dir, res
	} else if dir == South {
		res := m.cont[util.SerializeCoordRaw(currX, currY+1)]
		return currX, currY + 1, dir, res
	} else if dir == West {
		res := m.cont[util.SerializeCoordRaw(currX-1, currY)]
		return currX - 1, currY, dir, res
	}
	panic("Went out of bounds")
}

func parse(lines []string, edgeLength int) (Map, []Move, EdgeMap) {
	theMap := Map{}
	theMap.cont = make(map[string]Path)
	moves := make([]Move, 0)

	readmoves := false
	for y, line := range lines {
		if line == "" {
			readmoves = true
		} else if readmoves {
			tokens := strings.Split(line, "")
			curr := ""
			for _, token := range tokens {
				if token == "L" || token == "R" {
					if curr != "" {
						res, _ := strconv.Atoi(curr)
						moves = append(moves, Move{false, Right, res})
						curr = ""
					}
					if token == "L" {
						moves = append(moves, Move{true, Left, 0})
					} else {
						moves = append(moves, Move{true, Right, 0})
					}
				} else {
					curr += token
				}
			}
			if curr != "" {
				res, _ := strconv.Atoi(curr)
				moves = append(moves, Move{false, Right, res})
			}
		} else {
			for x, symbol := range line {
				if symbol == '.' {
					theMap.cont[util.SerializeCoordRaw(x, y)] = Clear
					theMap.xMax = util.MaxInt(theMap.xMax, x)
					theMap.yMax = util.MaxInt(theMap.yMax, y)
				} else if symbol == '#' {
					theMap.cont[util.SerializeCoordRaw(x, y)] = Wall
					theMap.xMax = util.MaxInt(theMap.xMax, x)
					theMap.yMax = util.MaxInt(theMap.yMax, y)
				}
			}
		}
	}

	edgeMap := calcEdges(theMap, edgeLength)

	return theMap, moves, edgeMap
}

type Direction string

const (
	North Direction = "n"
	East  Direction = "e"
	South Direction = "s"
	West  Direction = "w"
)

func (d Direction) getValue() int {
	switch d {
	case North:
		return 3
	case East:
		return 0
	case South:
		return 1
	case West:
		return 2
	}
	panic("Not a valid direction")
}

func (d Direction) doTurn(t Turn) Direction {
	if t == Right {
		switch d {
		case North:
			return East
		case East:
			return South
		case South:
			return West
		case West:
			return North
		}
	} else {
		switch d {
		case North:
			return West
		case East:
			return North
		case South:
			return East
		case West:
			return South
		}
	}
	panic("Not a valid direction or turn")
}

func part1(theMap Map, moves []Move) int {
	x := theMap.getStartingX()
	y := 0
	dir := East

	// fmt.Println(x, y, dir)
	for _, move := range moves {
		if move.isTurn {
			dir = dir.doTurn(move.turn)
			// fmt.Println(x, y, dir)
		} else {
			for step := 0; step < move.dist; step++ {
				newX, newY, path := theMap.getNext(x, y, dir)
				if path == Clear {
					x = newX
					y = newY
				}
				// fmt.Println(x, y, dir)
			}
		}
	}

	return 1000*(y+1) + 4*(x+1) + dir.getValue()
}

func part2(theMap Map, moves []Move, edges EdgeMap) int {
	x := theMap.getStartingX()
	y := 0
	dir := East

	for _, move := range moves {
		if move.isTurn {
			dir = dir.doTurn(move.turn)
		} else {
			for step := 0; step < move.dist; step++ {
				newX, newY, newDir, path := theMap.getNextWithEdges(x, y, dir, edges)
				if path == Clear {
					x = newX
					y = newY
					dir = newDir
				}
			}
		}
	}

	return 1000*(y+1) + 4*(x+1) + dir.getValue()
}

func main() {
	// lines, edgeLength := util.FileAsLines("ex1"), 4
	lines, edgeLength := util.FileAsLines("input"), 50
	theMap, moves, edgeMap := parse(lines, edgeLength)

	fmt.Println(part1(theMap, moves))
	fmt.Println(part2(theMap, moves, edgeMap))
}
