package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

const WALL_X = 7

type Direction int8

const (
	Left  Direction = -1
	Right Direction = 1
)

func getNextDir(dir []Direction, dirCounter int) (Direction, int) {
	return dir[dirCounter], (dirCounter + 1) % (len(dir))
}

func parse(line string) []Direction {
	result := make([]Direction, 0, len(line))
	for _, dir := range strings.Split(line, "") {
		if dir == "<" {
			result = append(result, Left)
		} else {
			result = append(result, Right)
		}
	}
	return result
}

func tryBlow(dir Direction, xOff int, yOff int, block []string, blockName Block, tower map[string]struct{}) int {
	if dir == Left && xOff == 0 {
		return 0
	} else if dir == Right && xOff+blockWidth(blockName) == WALL_X {
		return xOff
	}
	for _, blockPart := range block {
		bX, bY := util.DeserializeCoord(blockPart)
		_, inTower := tower[util.SerializeCoord(bX+int(dir)+xOff, bY+yOff)]
		if inTower {
			return xOff
		}
	}
	return xOff + int(dir)
}

func checkCanDrop(xOff int, yOff int, block []string, tower map[string]struct{}) bool {
	if yOff == 0 {
		return false
	}
	for _, blockPart := range block {
		bX, bY := util.DeserializeCoord(blockPart)
		_, inTower := tower[util.SerializeCoord(bX+xOff, bY+yOff-1)]
		if inTower {
			return false
		}
	}
	return true
}

func part1(directions []Direction) int {
	nextDirCounter := 0
	dir := Left
	highest := -1

	towerBlocks := make(map[string]struct{})

	for i := 0; i < 2022; i++ {
		block, blockName := getNthBlock(i)
		xOff := 2
		yOff := highest + 4

		hit := false

		for !hit {
			// Blow and drop
			dir, nextDirCounter = getNextDir(directions, nextDirCounter)
			xOff = tryBlow(dir, xOff, yOff, block, blockName, towerBlocks)
			canDrop := checkCanDrop(xOff, yOff, block, towerBlocks)
			if canDrop {
				yOff--
			} else {
				hit = true
			}
		}

		for _, blockPart := range block {
			bX, bY := util.DeserializeCoord(blockPart)
			towerBlocks[util.SerializeCoord(bX+xOff, bY+yOff)] = util.EMPTY_STRUCT
		}

		blockHighest := blockHeight(blockName) - 1 + yOff
		if blockHighest > highest {
			highest = blockHighest
		}
	}
	return highest + 1 // 1-indexed answer
}

func main() {
	// line := util.FileAsString("ex1")
	line := util.FileAsString("input")

	directions := parse(line)

	fmt.Println(part1(directions))
}
