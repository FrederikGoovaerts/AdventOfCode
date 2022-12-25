package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coord struct {
	x int
	y int
}

var offsetMap = map[string]coord{
	"e":  {1, 0},
	"w":  {-1, 0},
	"se": {0, -1},
	"sw": {-1, -1},
	"ne": {1, 1},
	"nw": {0, 1},
}

func getNeighbors(c coord) []coord {
	return []coord{
		{c.x + 1, c.y},
		{c.x - 1, c.y},
		{c.x, c.y - 1},
		{c.x - 1, c.y - 1},
		{c.x + 1, c.y + 1},
		{c.x, c.y + 1},
	}
}

func splitLine(line string) []string {
	res := make([]string, 0)
	curr := line
	for len(curr) > 0 {
		symbol := string(curr[0])
		curr = curr[1:]
		if symbol == "s" || symbol == "n" {
			symbol += string(curr[0])
			curr = curr[1:]
		}
		res = append(res, symbol)
	}
	return res
}

func getTile(changeList []string) coord {
	curr := coord{0, 0}
	for _, el := range changeList {
		offset := offsetMap[el]
		curr = coord{curr.x + offset.x, curr.y + offset.y}
	}
	return curr
}

func dailyFlip(prev map[coord]struct{}) map[coord]struct{} {
	toCheck := make(map[coord]struct{}, 0)
	nextBlackTiles := make(map[coord]struct{})
	for k := range prev {
		toCheck[k] = struct{}{}
		for _, n := range getNeighbors(k) {
			toCheck[n] = struct{}{}
		}
	}
	for k := range toCheck {
		_, blackTile := prev[k]
		blackAdjTiles := 0
		for _, n := range getNeighbors(k) {
			_, blackAdjTile := prev[n]
			if blackAdjTile {
				blackAdjTiles++
			}
		}
		if (blackTile && (blackAdjTiles == 1 || blackAdjTiles == 2)) || (!blackTile && blackAdjTiles == 2) {
			nextBlackTiles[k] = struct{}{}
		}
	}
	return nextBlackTiles
}

func main() {
	blackTiles := make(map[coord]struct{}, 0)

	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	for _, line := range lines {
		tile := getTile(splitLine(line))
		_, black := blackTiles[tile]
		if black {
			delete(blackTiles, tile)
		} else {
			blackTiles[tile] = struct{}{}
		}
	}
	fmt.Println(len(blackTiles))
	for i := 0; i < 100; i++ {
		blackTiles = dailyFlip(blackTiles)
	}
	fmt.Println(len(blackTiles))

}
