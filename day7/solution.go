package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const DISK_SPACE = 70000000
const UPDATE_SIZE = 30000000

type Node struct {
	dir bool
	// always 0 for directories
	size int
	// nil for files
	children *map[string]Node
}

func (n Node) getChildAt(dir []string) Node {
	if len(dir) == 0 {
		return n
	}

	child := (*n.children)[dir[0]]
	return child.getChildAt(dir[1:])
}

func newNodeMapPointer() *map[string]Node {
	result := make(map[string]Node)
	return &result
}

func collect(n Node) (int, int) {
	size := 0
	score := 0
	for _, child := range *n.children {
		if child.dir {
			childSize, childScore := collect(child)
			size += childSize
			score += childScore
		} else {
			size += child.size
		}
	}
	if size <= 100_000 {
		score += size
	}
	return size, score
}

func part1(root Node) int {
	_, score := collect(root)
	return score
}

func checkDelete(node Node, deleteSize int) (int, int) {
	childSizes := make([]int, 0)
	size := 0
	lowest := util.MAX_INT
	for _, child := range *node.children {
		if child.dir {
			childSize, lowestInChild := checkDelete(child, deleteSize)
			childSizes = append(childSizes, childSize)
			size += childSize
			if lowestInChild < lowest {
				lowest = lowestInChild
			}
		} else {
			size += child.size
		}
	}
	for _, childSize := range childSizes {
		if childSize > deleteSize && childSize < lowest {
			lowest = childSize
		}
	}
	if size > deleteSize && size < lowest {
		lowest = size
	}
	return size, lowest
}

func part2(root Node) int {
	size, _ := collect(root)
	deleteSize := UPDATE_SIZE - (DISK_SPACE - size)
	_, lowest := checkDelete(root, deleteSize)
	return lowest
}

func main() {
	lines := util.FileAsLines("input")
	dir := make([]string, 0)
	root := Node{true, 0, newNodeMapPointer()}

	cdRegex := regexp.MustCompile(`^\$ cd (.*)`)
	dirRegex := regexp.MustCompile(`^dir (.*)`)

	for _, line := range lines {
		if line == "$ ls" {
			continue
		}

		if line == "$ cd /" {
			dir = make([]string, 0)
		} else if line == "$ cd .." {
			dir = dir[0 : len(dir)-1]
		} else if cdRegex.MatchString(line) {
			dirName := cdRegex.FindStringSubmatch(line)[1]
			dir = append(dir, dirName)
		} else if dirRegex.MatchString(line) {
			dirName := dirRegex.FindStringSubmatch(line)[1]
			curr := root.getChildAt(dir)
			(*curr.children)[dirName] = Node{true, 0, newNodeMapPointer()}
		} else {
			parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(parts[0])
			curr := root.getChildAt(dir)
			(*curr.children)[parts[1]] = Node{false, size, nil}
		}
	}

	fmt.Println(part1(root))
	fmt.Println(part2(root))
}
