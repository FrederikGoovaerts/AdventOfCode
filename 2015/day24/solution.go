package main

import (
	"aoc/util"
	"fmt"
	"math"
)

func checkOtherGroupsPossible(group []int, remaining []int, index, weight int, remainingDepth int) bool {
	if index == len(remaining) {
		return false
	}
	currentSum := util.Sum(group)
	curr := remaining[index]

	if currentSum+curr == weight {
		if remainingDepth == 0 || checkOtherGroupsPossible([]int{}, util.Remove(remaining, curr), 0, weight, remainingDepth-1) {
			return true
		}
	} else if currentSum+curr < weight {
		if checkOtherGroupsPossible(append(group, curr), util.Remove(remaining, curr), index, weight, remainingDepth) {
			return true
		}
	}

	return checkOtherGroupsPossible(group, remaining, index+1, weight, remainingDepth)
}

func formFirstGroup(group []int, remaining []int, index int, weight int, checkDepth int) (int, int) {
	if index == len(remaining) {
		return math.MaxInt, math.MaxInt
	}

	currentSum := util.Sum(group)
	curr := remaining[index]

	if currentSum+curr > weight {
		// As the first group would be overweight, just skip this element
		return formFirstGroup(group, remaining, index+1, weight, checkDepth)
	} else {
		bestSize := math.MaxInt
		bestQe := math.MaxInt
		if currentSum+curr == weight {
			// If the current element would complete the weight for the first group, check the others
			gr := append(group, curr)
			rem := util.Remove(remaining, curr)
			if checkOtherGroupsPossible([]int{rem[0]}, rem[1:], 0, weight, checkDepth) {
				bestSize = len(gr)
				bestQe = util.Product(gr)
			}
		} else {
			// If the current element still does not complete the weight, add it to the first group and continue
			bestSize, bestQe = formFirstGroup(append(group, curr), util.Remove(remaining, curr), index, weight, checkDepth)
		}

		// Regardless of the case before, also continue without the current element
		skipSize, skipQe := formFirstGroup(group, remaining, index+1, weight, checkDepth)
		if skipSize < bestSize {
			bestSize = skipSize
			bestQe = skipQe
		} else if skipSize == bestSize {
			bestQe = util.MinInt(bestQe, skipQe)
		}
		return bestSize, bestQe
	}
}

func part1(values []int) int {
	groupWeight := util.Sum(values) / 3
	_, bestQe := formFirstGroup([]int{}, values, 0, groupWeight, 0)
	return bestQe
}

func part2(values []int) int {
	groupWeight := util.Sum(values) / 4
	_, bestQe := formFirstGroup([]int{}, values, 0, groupWeight, 1)
	return bestQe
}

func main() {
	// input := util.FileAsNumbers("ex1")
	input := util.FileAsNumbers("input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
