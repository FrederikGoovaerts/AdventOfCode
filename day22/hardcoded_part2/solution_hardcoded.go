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

func (m Map) getStartingX() int {
	for x := 0; x <= m.xMax; x++ {
		res, present := m.cont[util.SerializeCoord(x, 0)]
		if present && res == Clear {
			return x
		}
	}
	return -1
}

func (m Map) getNextWithEdges(currX, currY int, dir Direction, e EdgeMap) (int, int, Direction, Path) {
	if e.has(currX, currY, dir) {
		newX, newY, newDir := e.getNext(currX, currY, dir)
		path := m.cont[util.SerializeCoord(newX, newY)]
		return newX, newY, newDir, path
	}

	if dir == North {
		res := m.cont[util.SerializeCoord(currX, currY-1)]
		return currX, currY - 1, dir, res
	} else if dir == East {
		res := m.cont[util.SerializeCoord(currX+1, currY)]
		return currX + 1, currY, dir, res
	} else if dir == South {
		res := m.cont[util.SerializeCoord(currX, currY+1)]
		return currX, currY + 1, dir, res
	} else if dir == West {
		res := m.cont[util.SerializeCoord(currX-1, currY)]
		return currX - 1, currY, dir, res
	}
	panic("Went out of bounds")
}

type EdgeMap map[string]string

func (e *EdgeMap) add(x1, y1, rawDir1, x2, y2, rawDir2 string) {
	dir1 := Direction(rawDir1)
	dir2 := Direction(rawDir2)
	(*e)[util.Serialize(x1, y1, dir1)] = util.Serialize(x2, y2, dir2)
	(*e)[util.Serialize(x2, y2, dir2.doTurn(Left).doTurn(Left))] = util.Serialize(x1, y1, dir1.doTurn(Left).doTurn(Left))
}

func (e EdgeMap) has(x, y int, dir Direction) bool {
	_, present := e[util.Serialize(x, y, dir)]
	return present
}

func (e EdgeMap) getNext(x, y int, dir Direction) (int, int, Direction) {
	res := e[util.Serialize(x, y, dir)]
	parts := strings.Split(res, " ")
	newX, _ := strconv.Atoi(parts[0])
	newY, _ := strconv.Atoi(parts[1])
	return newX, newY, Direction(parts[2])
}

func parse(lines []string) (Map, []Move, EdgeMap) {
	theMap := Map{}
	theMap.cont = make(map[string]Path)
	moves := make([]Move, 0)
	edges := make(EdgeMap)

	readmoves := false
	readedges := false
	for y, line := range lines {
		if line == "" {
			readmoves = true
		} else if readmoves && !readedges {
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
			readedges = true
		} else if readedges {
			parts := strings.Split(line, " ")
			edges.add(parts[0], parts[1], parts[2], parts[3], parts[4], parts[5])
		} else {
			for x, symbol := range line {
				if symbol == '.' {
					theMap.cont[util.SerializeCoord(x, y)] = Clear
					theMap.xMax = util.MaxInt(theMap.xMax, x)
					theMap.yMax = util.MaxInt(theMap.yMax, y)
				} else if symbol == '#' {
					theMap.cont[util.SerializeCoord(x, y)] = Wall
					theMap.xMax = util.MaxInt(theMap.xMax, x)
					theMap.yMax = util.MaxInt(theMap.yMax, y)
				}
			}
		}
	}

	return theMap, moves, edges
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
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")
	theMap, moves, edges := parse(lines)

	fmt.Println(part2(theMap, moves, edges))
}
