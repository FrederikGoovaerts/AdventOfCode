package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"aoc/utils"
)

type imageTile struct {
	name     string
	contents []string
	frozen   bool
}

type tileSetting struct {
	tile  imageTile
	up    string
	down  string
	left  string
	right string
}

func (t *imageTile) flip() {
	for i, el := range t.contents {
		t.contents[i] = utils.Reverse(el)
	}
}

func (t *imageTile) turn() {
	newContents := make([]string, 0)
	for i := 0; i < len(t.contents); i++ {
		newLine := ""
		for j := 1; j <= len(t.contents); j++ {
			newLine += t.contents[len(t.contents)-j][i : i+1]
		}
		newContents = append(newContents, newLine)
	}
	t.contents = newContents
}

func (t imageTile) size() int {
	return len(t.contents)
}

func (t imageTile) getClockWiseBorder(dir string) string {
	if dir == "U" {
		return t.contents[0]
	} else if dir == "R" {
		newLine := ""
		for i := 0; i < t.size(); i++ {
			newLine += t.contents[i][t.size()-1 : t.size()]
		}
		return newLine
	} else if dir == "L" {
		newLine := ""
		for i := t.size() - 1; i >= 0; i-- {
			newLine += t.contents[i][0:1]
		}
		return newLine
	} else {
		return utils.Reverse(t.contents[t.size()-1])
	}
}

var opposingSide = map[string]string{
	"U": "D",
	"D": "U",
	"R": "L",
	"L": "R",
}

func (t imageTile) matchesSide(dir string, other *imageTile) bool {
	matches := func() bool {
		return t.getClockWiseBorder(dir) == utils.Reverse(other.getClockWiseBorder(opposingSide[dir]))
	}
	if other.frozen {
		return matches()
	}

	for i := 0; i < 4; i++ {
		if matches() {
			return true
		}
		other.turn()
	}
	other.flip()
	for i := 0; i < 4; i++ {
		if matches() {
			return true
		}
		other.turn()
	}

	return false
}

func (t imageTile) getTileNumber() int {
	val, _ := strconv.Atoi(t.name[5:9])
	return val
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	tileStrings := strings.Split(strings.TrimSpace(string(dat)), "\n\n")

	origin := tileSetting{}
	tiles := make(map[string]imageTile, 0)
	remainingTiles := make(map[string]imageTile, 0)
	doneTiles := make(map[string]*tileSetting, 0)
	for _, tileString := range tileStrings {
		splitTile := strings.Split(tileString, "\n")
		tileID := splitTile[0]
		tileData := splitTile[1:]
		tile := imageTile{tileID, tileData, false}
		tiles[tileID] = tile
		remainingTiles[tileID] = tile

		if origin.tile.name == "" {
			origin = tileSetting{tile, "", "", "", ""}
		}
	}

	delete(remainingTiles, origin.tile.name)
	toCheck := make([]*tileSetting, 0)
	toCheck = append(toCheck, &origin)
	origin.tile.frozen = true

	handleSettingSide := func(setting *tileSetting, dir string) *tileSetting {
		tile := setting.tile
		if (dir == "U" && setting.up != "") || (dir == "D" && setting.down != "") || (dir == "L" && setting.left != "") || (dir == "R" && setting.right != "") {
			return setting
		}
		for k, v := range remainingTiles {
			if tile.matchesSide(dir, &v) {
				var newSetting tileSetting
				if dir == "U" {
					newSetting = tileSetting{}
					newSetting.tile = v
					newSetting.down = setting.tile.name
					setting.up = k
				} else if dir == "D" {
					newSetting = tileSetting{}
					newSetting.tile = v
					newSetting.up = setting.tile.name
					setting.down = k
				} else if dir == "L" {
					newSetting = tileSetting{}
					newSetting.tile = v
					newSetting.right = setting.tile.name
					setting.left = k
				} else if dir == "R" {
					newSetting = tileSetting{}
					newSetting.tile = v
					newSetting.left = setting.tile.name
					setting.right = k
				}
				newSetting.tile.frozen = true
				delete(remainingTiles, k)
				doneTiles[k] = &newSetting
				toCheck = append(toCheck, &newSetting)
				return setting
			}
		}
		for k, v := range doneTiles {
			if tile.matchesSide(dir, &v.tile) {
				if dir == "U" {
					v.down = setting.tile.name
					setting.up = k
				} else if dir == "D" {
					v.up = setting.tile.name
					setting.down = k
				} else if dir == "L" {
					v.right = setting.tile.name
					setting.left = k
				} else if dir == "R" {
					v.left = setting.tile.name
					setting.right = k
				}
				return setting
			}
		}
		return setting
	}

	for len(toCheck) > 0 {
		var checking *tileSetting
		checking, toCheck = toCheck[0], toCheck[1:]
		checking = handleSettingSide(checking, "U")
		checking = handleSettingSide(checking, "R")
		checking = handleSettingSide(checking, "D")
		checking = handleSettingSide(checking, "L")
		doneTiles[checking.tile.name] = checking
	}
	topLeft := origin
	for topLeft.up != "" || topLeft.left != "" {
		if topLeft.up != "" {
			topLeft = *doneTiles[topLeft.up]
		} else {
			topLeft = *doneTiles[topLeft.left]
		}
	}
	curr := topLeft
	currRow := make([]imageTile, 0)
	currRowStart := topLeft
	orderedTiles := make([][]imageTile, 0)
	for curr.right != "" {
		currRow = append(currRow, curr.tile)
		curr = *doneTiles[curr.right]
		if curr.right == "" {
			currRow = append(currRow, curr.tile)
			orderedTiles = append(orderedTiles, currRow)
			currRow = make([]imageTile, 0)
			if currRowStart.down != "" {
				currRowStart = *doneTiles[currRowStart.down]
				curr = currRowStart
			}
		}
	}
	mult := 1
	for i, a := range orderedTiles {
		for j, b := range a {
			if (i == 0 || i == len(orderedTiles)-1) && (j == 0 || j == len(a)-1) {
				mult *= b.getTileNumber()
			}
		}
	}
	fmt.Println(mult)

	acc := ""
	length := 10 // Eh, we'll just take it

	// Top row
	for row := 0; row < len(orderedTiles); row++ {
		for i := 1; i <= length-2; i++ {
			for _, tile := range orderedTiles[row] {
				start := 1
				end := 9
				acc += tile.contents[i][start:end]

			}
			acc += "\n"
		}
	}
	imageLines := strings.Fields(acc)
	masterTile := imageTile{"master", imageLines, true}

	line1Regex := regexp.MustCompile("..................#.")
	line2Regex := regexp.MustCompile("#....##....##....###")
	line3Regex := regexp.MustCompile(".#..#..#..#..#..#...")

	max := 0

	checkHits := func() {
		image := masterTile.contents
		count := 0
		monsters := false
		countHits := func(line string) {
			for _, r := range line {
				if r == '#' {
					count++
				}
			}
		}
		for i := 0; i < len(image)-2; i++ {
			for j := 0; j < len(image[0])-19; j++ {
				firstLine := image[i][j : j+20]
				firstHit := line1Regex.MatchString(firstLine)
				secondLine := image[i+1][j : j+20]
				secondHit := line2Regex.MatchString(secondLine)
				thirdLine := image[i+2][j : j+20]
				thirdHit := line3Regex.MatchString(thirdLine)
				if firstHit && secondHit && thirdHit {
					monsters = true
					count -= 15
				}
			}
		}
		for _, el := range image {
			countHits(el)
		}
		if monsters && count > max {
			max = count
		}
	}

	for i := 0; i < 4; i++ {
		checkHits()
		masterTile.turn()
	}
	masterTile.flip()
	for i := 0; i < 4; i++ {
		checkHits()
		masterTile.turn()
	}
	fmt.Println(max)

}
