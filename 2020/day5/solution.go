package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type seat struct {
	row int64
	col int64
	id int64
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	seats := make([]seat, 0)
	var highestId int64
	var lowestId int64
	lowestId = math.MaxInt64
	for _, line := range lines {
		rowNumber, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(line[:7], "F", "0"), "B", "1"), 2, 64)
		colNumber, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(line[7:], "L", "0"), "R", "1"), 2, 64)
		id:= rowNumber* 8 + colNumber
		seats = append(seats, seat{rowNumber, colNumber, id})
		if id > highestId {
			highestId = id
		}
		if id < lowestId {
			lowestId = id
		}
	}

	fmt.Println(highestId)

	for i:=lowestId; i < highestId; i++ {
		var found bool
		for _, seat := range seats {
			if seat.id == i {
				found = true
			}
		}
		if !found {
			fmt.Println(i)
		}
	}
}
