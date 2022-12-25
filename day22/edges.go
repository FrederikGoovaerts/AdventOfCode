package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

func getScanDirectionOffset(dir Direction) (int, int) {
	switch dir {
	case North:
		return 0, -1
	case East:
		return 1, 0
	case South:
		return 0, 1
	case West:
		return -1, 0
	}
	panic("illegal direction provided")
}

func calcEdges(theMap Map, edgeLength int) EdgeMap {
	x, y := theMap.getStartingX(), 0
	scanDirection := East
	currEdge := FU

	edgeCounter := 0
	// elements are of the form: FU -> South
	// This means that entering the cube on the FU edge makes the arrow direction southward
	edgeInDirection := make(map[Edge]Direction)
	// elements are of the form: FU 1 -> 10 13
	// This means that the first space of edge FU is 10, 13.
	// Spaces are numbered clockwise on their face, which means that corresponding edges have reversed numbering.
	edgeMembers := make(map[string]string)
	for edgeCounter < 24 {
		xOff, yOff := getScanDirectionOffset(scanDirection)
		edgeInDirection[currEdge] = scanDirection.doTurn(Right)
		for i := 0; i < edgeLength; i++ {
			edgeMembers[util.Serialize(currEdge, i)] = util.Serialize(x+i*xOff, y+i*yOff)
		}
		// Go to last members of current edge
		x, y = x+xOff*(edgeLength-1), y+yOff*(edgeLength-1)

		// Check what crossing we encounter
		crossing := Corner
		if theMap.has(x+xOff, y+yOff) {
			// The edge either continues or makes a fold. After figuring out
			// which, already move to start of new edge, depending on crossing
			foldXOff, foldYOff := getScanDirectionOffset(scanDirection.doTurn(Left))
			if theMap.has(x+xOff+foldXOff, y+yOff+foldYOff) {
				crossing = Fold
				x = x + xOff + foldXOff
				y = y + yOff + foldYOff
				scanDirection = scanDirection.doTurn(Left)
			} else {
				crossing = Straight
				x = x + xOff
				y = y + yOff
			}
		} else {
			scanDirection = scanDirection.doTurn(Right)
		}
		currEdge = currEdge.cross(crossing)

		edgeCounter++
	}

	em := make(EdgeMap)

	for k, v := range edgeMembers {
		parts := strings.Split(k, " ")
		edge := Edge(parts[0])
		id, _ := strconv.Atoi(parts[1])
		x1, y1 := util.DeserializeCoordRaw(v)
		dir1 := edgeInDirection[edge].doTurn(Left).doTurn(Left)

		x2, y2 := util.DeserializeCoordRaw(edgeMembers[string(edge.getCorrespondingEdge())+" "+fmt.Sprint(edgeLength-1-id)])
		dir2 := edgeInDirection[edge.getCorrespondingEdge()]

		em.add(x1, y1, dir1, x2, y2, dir2)
	}

	return em
}

type EdgeMap map[string]string

func (e *EdgeMap) add(x1, y1 int, dir1 Direction, x2, y2 int, dir2 Direction) {
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

type EdgeCrossing string

const (
	Straight EdgeCrossing = "s"
	Corner   EdgeCrossing = "c"
	Fold     EdgeCrossing = "f"
)

type Edge string

// The first letter symbolizes the face of the cube for the edge, the second is
// which edge of said face it is. This means that each edge of the cube has two
// versions. This is necessary to combine the correct matches.
// See cube_visualisation.png for a visual representation.
const (
	FU Edge = "FU"
	FR Edge = "FR"
	FD Edge = "FD"
	FL Edge = "FL"
	UF Edge = "UF"
	UL Edge = "UL"
	UB Edge = "UB"
	UR Edge = "UR"
	BU Edge = "BU"
	BR Edge = "BR"
	BL Edge = "BL"
	BD Edge = "BD"
	DF Edge = "DF"
	DB Edge = "DB"
	DR Edge = "DR"
	DL Edge = "DL"
	LF Edge = "LF"
	LU Edge = "LU"
	LB Edge = "LB"
	LD Edge = "LD"
	RF Edge = "RF"
	RU Edge = "RU"
	RB Edge = "RB"
	RD Edge = "RD"
)

func (e Edge) getCorrespondingEdge() Edge {
	parts := strings.Split(string(e), "")
	return Edge(parts[1] + parts[0])
}

func (e Edge) cross(crossing EdgeCrossing) Edge {
	switch crossing {
	case Straight:
		switch e {
		case FU:
			return RU
		case RU:
			return BU
		case BU:
			return LU
		case LU:
			return FU
		case FR:
			return DR
		case DR:
			return BR
		case BR:
			return UR
		case UR:
			return FR
		case FD:
			return LD
		case LD:
			return BD
		case BD:
			return RD
		case RD:
			return FD
		case FL:
			return UL
		case UL:
			return BL
		case BL:
			return DL
		case DL:
			return FL
		case UF:
			return LF
		case LF:
			return DF
		case DF:
			return RF
		case RF:
			return UF
		case UB:
			return RB
		case RB:
			return DB
		case DB:
			return LB
		case LB:
			return UB
		}
	case Corner:
		switch e {
		case FU:
			return FR
		case FR:
			return FD
		case FD:
			return FL
		case FL:
			return FU
		case UF:
			return UL
		case UL:
			return UB
		case UB:
			return UR
		case UR:
			return UF
		case BU:
			return BL
		case BL:
			return BD
		case BD:
			return BR
		case BR:
			return BU
		case DF:
			return DR
		case DR:
			return DB
		case DB:
			return DL
		case DL:
			return DF
		case LU:
			return LF
		case LF:
			return LD
		case LD:
			return LB
		case LB:
			return LU
		case RU:
			return RB
		case RB:
			return RD
		case RD:
			return RF
		case RF:
			return RU
		}
	case Fold:
		switch e {
		case FU:
			return UF
		case FR:
			return RF
		case FD:
			return DF
		case FL:
			return LF
		case UF:
			return FU
		case UL:
			return LU
		case UB:
			return BU
		case UR:
			return RU
		case BU:
			return UB
		case BR:
			return RB
		case BL:
			return LB
		case BD:
			return DB
		case DF:
			return FD
		case DB:
			return BD
		case DR:
			return RD
		case DL:
			return LD
		case LF:
			return FL
		case LU:
			return UL
		case LB:
			return BL
		case LD:
			return DL
		case RF:
			return FR
		case RU:
			return UR
		case RB:
			return BR
		case RD:
			return DR
		}
	}
	panic("illegal combination!")
}
