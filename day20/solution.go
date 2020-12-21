package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var tileBorders map[string][]string

func extractBorders(input []string) []string {
	borders := make([]string, 0)
	// Extract all "top" borders at different orientations

	// Top side
	borders = append(borders, input[0])
	// Right and left side
	right := make([]rune, len(input))
	left := make([]rune, len(input))
	horizontalLength := len(input[0])
	for i := 0; i < len(input); i++ {
		right[i] = []rune(input[i])[horizontalLength-1]
		left[i] = []rune(input[horizontalLength-1-i])[0]
	}
	borders = append(borders, string(right))
	borders = append(borders, string(left))
	// Bottom side
	borders = append(borders, simpleReverse(input[len(input)-1]))
	return borders
}

func simpleReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func simpleMatchBorders(id string) []string {
	res := make([]string, 0)
	borders := tileBorders[id]
	for k, v := range tileBorders {
		if k != id {
			for _, ownBorder := range borders {
				for _, otherBorder := range v {
					if ownBorder == otherBorder || ownBorder == simpleReverse(otherBorder) {
						res = append(res, k)
					}
				}
			}
		}
	}
	return res
}

func main() {
	tileBorders = make(map[string][]string, 0)
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	tiles := strings.Split(strings.TrimSpace(string(dat)), "\n\n")

	for _, tile := range tiles {
		splitTile := strings.Split(tile, "\n")
		tileID := splitTile[0]
		tileData := splitTile[1:]
		tileBorders[tileID] = extractBorders(tileData)
		fmt.Println(tileID, tileBorders[tileID])
	}

	mult := 1
	for k := range tileBorders {
		matches := simpleMatchBorders(k)
		if len(matches) == 2 {
			val, _ := strconv.Atoi(k[5:9])
			mult *= val
		}
	}
	fmt.Println(mult)

}
