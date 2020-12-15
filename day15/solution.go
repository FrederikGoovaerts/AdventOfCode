package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	numStrings := strings.Split(strings.TrimSpace(string(dat)), ",")

	turnMap := make(map[int]int, 0)
	turnCounter := 0
	lastNum := 0

	for i := 0; i < len(numStrings); i++ {
		num, _ := strconv.Atoi(numStrings[i])
		turnCounter++
		lastNum = num
		turnMap[num] = turnCounter
	}
	delete(turnMap, lastNum)

	for turnCounter < 2020 {
		turnCounter++
		lastTurn, seen := turnMap[lastNum]
		turnMap[lastNum] = turnCounter - 1
		if !seen || lastTurn == turnCounter-1 {
			lastNum = 0
		} else {
			lastNum = turnCounter - lastTurn - 1
		}

	}
	fmt.Println(lastNum)

	for turnCounter < 30000000 {
		turnCounter++
		lastTurn, seen := turnMap[lastNum]
		turnMap[lastNum] = turnCounter - 1
		if !seen || lastTurn == turnCounter-1 {
			lastNum = 0
		} else {
			lastNum = turnCounter - lastTurn - 1
		}

	}
	fmt.Println(lastNum)

}
