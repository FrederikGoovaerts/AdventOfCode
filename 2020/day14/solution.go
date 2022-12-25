package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type maskElement struct {
	pos int
	val rune
}

func parseMask(line string) []maskElement {
	maskString := strings.Split(line, "mask = ")[1]
	elements := make([]maskElement, 0)
	for i := 1; i <= len(maskString); i++ {
		valString := []rune(maskString)[len(maskString)-i]
		if valString != 'X' {
			elements = append(elements, maskElement{i - 1, valString})
		}
	}
	return elements
}

func parseMaskV2(line string) [][]maskElement {
	runes := []rune(strings.Split(line, "mask = ")[1])
	masks := make([][]maskElement, 0)
	masks = append(masks, make([]maskElement, 0))
	for i := 1; i <= len(runes); i++ {
		curr := runes[len(runes)-i]
		if curr == '1' {
			newMasks := make([][]maskElement, 0)
			for _, mask := range masks {
				newMask := append(mask, maskElement{i - 1, '1'})
				newMasks = append(newMasks, newMask)
			}
			masks = newMasks
		} else if curr == 'X' {
			newMasks := make([][]maskElement, 0)
			for _, mask := range masks {
				maskVariation1 := make([]maskElement, len(mask))
				copy(maskVariation1, mask)
				maskVariation1 = append(maskVariation1, maskElement{i - 1, '1'})
				newMasks = append(newMasks, maskVariation1)

				maskVariation2 := make([]maskElement, len(mask))
				copy(maskVariation2, mask)
				maskVariation2 = append(maskVariation2, maskElement{i - 1, '0'})
				newMasks = append(newMasks, maskVariation2)
			}
			masks = newMasks
		}
	}
	return masks
}

func parseAssign(line string) (int, int) {
	firstSplit := strings.Split(line, "] = ")
	secondSplit := strings.Split(firstSplit[0], "mem[")
	loc, _ := strconv.Atoi(secondSplit[1])
	val, _ := strconv.Atoi(firstSplit[1])
	return loc, val
}

func applyMask(val string, mask []maskElement) string {
	res := []rune(val)
	for _, el := range mask {
		index := len(val) - 1 - el.pos
		res[index] = el.val
	}
	return string(res)
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	currentMask := make([]maskElement, 0)
	memoryMap := make(map[int]int, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "mask = ") {
			currentMask = parseMask(line)
		} else {
			loc, val := parseAssign(line)
			binVal := "000000000000000000000000000000000000000000000000000" + strconv.FormatInt(int64(val), 2)
			binVal = applyMask(binVal, currentMask)
			decVal, _ := strconv.ParseInt(binVal, 2, 64)
			memoryMap[loc] = int(decVal)
		}
	}

	sum := 0
	for _, v := range memoryMap {
		sum += v
	}
	fmt.Println(sum)

	masks := make([][]maskElement, 0)
	memoryMap = make(map[int]int, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "mask = ") {
			masks = parseMaskV2(line)
		} else {
			loc, val := parseAssign(line)
			binLoc := "0000000000000000000000000000000" + strconv.FormatInt(int64(loc), 2)
			for _, mask := range masks {
				binLoc = applyMask(binLoc, mask)
				decLoc, _ := strconv.ParseInt(binLoc, 2, 64)
				memoryMap[int(decLoc)] = int(val)
			}
		}
	}

	sum = 0
	for _, v := range memoryMap {
		sum += v
	}
	fmt.Println(sum)
}
