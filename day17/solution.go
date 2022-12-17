package main

import (
	"aoc/util"
	"fmt"
	"strconv"
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

const DESIRED_ROUNDS = 1_000_000_000_000

func getPatternLength(blockHistory []string) int {
	histStart := len(blockHistory) - 1
	for i := 1; i < 3000; i++ {
		if blockHistory[histStart] == blockHistory[histStart-i] {
			matches := true
			for j := 0; j < i; j++ {
				if blockHistory[histStart-j] != blockHistory[histStart-i-j] {
					matches = false
				}
			}
			if matches {
				return i
			}
		}
	}
	return -1
}

func getLoopHeight(blockHistory []string, length int) int {
	histStart := len(blockHistory) - 1
	result := 0
	for i := 0; i < length; i++ {
		parts := strings.Split(blockHistory[histStart-i], " ")
		h, _ := strconv.Atoi(parts[0])
		result += h
	}
	return result
}

func getRemainingHeight(blockHistory []string, length, rounds int) int {
	histStart := len(blockHistory) - 1
	result := 0
	for i := 1; i <= rounds; i++ {
		parts := strings.Split(blockHistory[histStart-length+i], " ")
		h, _ := strconv.Atoi(parts[0])
		result += h
	}
	return result
}

func part2(directions []Direction) int {
	blockHistory := make([]string, 0)

	nextDirCounter := 0
	dir := Left
	highest := -1

	towerBlocks := make(map[string]struct{})

	for i := 0; i < 5000; i++ {
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
		heightDiff := blockHighest - highest
		if heightDiff > 0 {
			highest = blockHighest
			blockHistory = append(blockHistory, fmt.Sprint(heightDiff)+" "+string(blockName)+fmt.Sprint(xOff))
		} else {
			blockHistory = append(blockHistory, "0 "+string(blockName)+fmt.Sprint(xOff))
		}
	}

	roundsLeft := DESIRED_ROUNDS - 5000
	totalHeight := highest

	length := getPatternLength(blockHistory)
	loopHeight := getLoopHeight(blockHistory, length)

	loopsLeft := roundsLeft / length
	roundsLeft = roundsLeft % length
	totalHeight += loopsLeft * loopHeight

	remainingHeight := getRemainingHeight(blockHistory, length, roundsLeft)

	return totalHeight + remainingHeight + 1
}

func main() {
	// line := util.FileAsString("ex1")
	line := util.FileAsString("input")

	directions := parse(line)

	fmt.Println(part1(directions))
	fmt.Println(part2(directions))
}
