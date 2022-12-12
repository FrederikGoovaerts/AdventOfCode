package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Forest struct {
	trees [][]int
}

func (f *Forest) isVisibleFromNorth(x int, y int) bool {
	if x == 0 {
		return true
	}
	height := f.trees[x][y]
	for northX := x - 1; northX >= 0; northX-- {
		if f.trees[northX][y] >= height {
			return false
		}
	}
	return true
}

func (f *Forest) getScenicScoreNorth(x int, y int) int {
	if x == 0 {
		return 0
	}
	score := 0
	height := f.trees[x][y]
	for northX := x - 1; northX >= 0; northX-- {
		if f.trees[northX][y] >= height {
			score++
			return score
		} else {
			score++
		}
	}
	return score
}

func (f *Forest) isVisibleFromSouth(x int, y int) bool {
	if x == len(f.trees)-1 {
		return true
	}
	height := f.trees[x][y]
	for southX := x + 1; southX < len(f.trees); southX++ {
		if f.trees[southX][y] >= height {
			return false
		}
	}
	return true
}

func (f *Forest) getScenicScoreSouth(x int, y int) int {
	if x == len(f.trees)-1 {
		return 0
	}
	score := 0
	height := f.trees[x][y]
	for southX := x + 1; southX < len(f.trees); southX++ {
		if f.trees[southX][y] >= height {
			score++
			return score
		} else {
			score++
		}
	}
	return score
}

func (f *Forest) isVisibleFromWest(x int, y int) bool {
	if y == 0 {
		return true
	}
	height := f.trees[x][y]
	for westY := y - 1; westY >= 0; westY-- {
		if f.trees[x][westY] >= height {
			return false
		}
	}
	return true
}

func (f *Forest) getScenicScoreWest(x int, y int) int {
	if y == 0 {
		return 0
	}
	score := 0
	height := f.trees[x][y]
	for westY := y - 1; westY >= 0; westY-- {
		if f.trees[x][westY] >= height {
			score++
			return score
		} else {
			score++
		}
	}
	return score
}

func (f *Forest) isVisibleFromEast(x int, y int) bool {
	if y == len(f.trees[0])-1 {
		return true
	}
	height := f.trees[x][y]
	for eastY := y + 1; eastY < len(f.trees[0]); eastY++ {
		if f.trees[x][eastY] >= height {
			return false
		}
	}
	return true
}

func (f *Forest) getScenicScoreEast(x int, y int) int {
	if y == len(f.trees[0])-1 {
		return 0
	}
	score := 0
	height := f.trees[x][y]
	for eastY := y + 1; eastY < len(f.trees[0]); eastY++ {
		if f.trees[x][eastY] >= height {
			score++
			return score
		} else {
			score++
		}
	}
	return score
}

func (f *Forest) isVisible(x int, y int) bool {
	return f.isVisibleFromNorth(x, y) || f.isVisibleFromEast(x, y) || f.isVisibleFromSouth(x, y) || f.isVisibleFromWest(x, y)
}

func (f *Forest) getScenicScore(x int, y int) int {
	north := f.getScenicScoreNorth(x, y)
	east := f.getScenicScoreEast(x, y)
	south := f.getScenicScoreSouth(x, y)
	west := f.getScenicScoreWest(x, y)
	return north * east * south * west
}

func part1(f Forest) int {
	result := 0
	for x := 0; x < len(f.trees); x++ {
		for y := 0; y < len(f.trees[0]); y++ {
			if f.isVisible(x, y) {
				result++
			}
		}
	}
	return result
}

func part2(f Forest) int {
	result := 0
	for x := 0; x < len(f.trees); x++ {
		for y := 0; y < len(f.trees[0]); y++ {
			score := f.getScenicScore(x, y)
			if score > result {
				result = score
			}
		}
	}
	return result
}

func main() {
	lines := util.FileAsLines("input")
	forest := Forest{make([][]int, 0)}

	for _, line := range lines {
		if line != "" {
			numbers := strings.Split(line, "")
			forestLine := make([]int, len(numbers))
			for index, value := range numbers {
				numberValue, _ := strconv.Atoi(value)
				forestLine[index] = numberValue
			}
			forest.trees = append(forest.trees, forestLine)
		}
	}

	fmt.Println(part1(forest))
	fmt.Println(part2(forest))
}
