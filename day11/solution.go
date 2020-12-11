package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type seat struct {
	x int
	y int
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func intifySeat(s seat) int {
	return s.x*100000 + s.y
}

func stringifySeating(seats map[int]struct{}) string {
	result := ""
	list := make([]int, 0)
	for k := range seats {
		list = append(list, k)
	}
	sort.Ints(list)
	for _, v := range list {
		result += fmt.Sprint(v) + ","
	}
	return result
}

func adjacentSeatNeighbors(s seat, seats map[seat]struct{}, dist int) []seat {
	return []seat{
		{s.x - 1, s.y - 1},
		{s.x - 1, s.y},
		{s.x - 1, s.y + 1},
		{s.x, s.y - 1},
		{s.x, s.y + 1},
		{s.x + 1, s.y - 1},
		{s.x + 1, s.y},
		{s.x + 1, s.y + 1},
	}
}

type change struct {
	x int
	y int
}

func viewSeatNeighbors(s seat, seats map[seat]struct{}, dist int) []seat {
	neighbors := make([]seat, 0)
	changes := []change{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, dir := range changes {
		curr := seat{s.x + dir.x, s.y + dir.y}
		found := false
		for curr.x > -1 && curr.y > -1 && curr.x < dist+1 && curr.y < dist+1 && !found {
			_, found = seats[curr]
			if found {
				neighbors = append(neighbors, curr)
			} else {
				curr = seat{curr.x + dir.x, curr.y + dir.y}
			}
		}
	}
	return neighbors
}

func performIteration(seats map[seat]struct{}, seatTaken map[int]struct{}, maxDist int, neighborThreshold int, seatNeighbors func(seat, map[seat]struct{}, int) []seat) map[int]struct{} {
	newSeats := make(map[int]struct{}, 0)
	for seat := range seats {
		_, taken := seatTaken[intifySeat(seat)]
		neighborsTaken := 0
		for _, neighbor := range seatNeighbors(seat, seats, maxDist) {
			_, neighborTaken := seatTaken[intifySeat(neighbor)]
			if neighborTaken {
				neighborsTaken++
			}
		}
		if (taken && neighborsTaken < neighborThreshold) || !taken && neighborsTaken == 0 {
			newSeats[intifySeat(seat)] = struct{}{}
		}
	}
	return newSeats
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	maxDist := max(len(lines), len(lines[0]))

	seats := make(map[seat]struct{}, 0)
	originalSeatTaken := make(map[int]struct{}, 0)

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				seat := seat{x, y}
				seats[seat] = struct{}{}
				if char == '#' {
					originalSeatTaken[intifySeat(seat)] = struct{}{}
				}
			}
		}
	}
	seatTaken := make(map[int]struct{}, 0)
	for k, v := range originalSeatTaken {
		seatTaken[k] = v
	}

	visited := make(map[string]struct{})
	visited[stringifySeating(seatTaken)] = struct{}{}
	iteration := 0
	looped := false
	for !looped {
		seatTaken = performIteration(seats, seatTaken, maxDist, 4, adjacentSeatNeighbors)
		iteration++
		seatString := stringifySeating(seatTaken)
		_, found := visited[seatString]
		if found {
			looped = true
		} else {
			visited[seatString] = struct{}{}
		}
	}

	fmt.Println(len(seatTaken))

	seatTaken = make(map[int]struct{}, 0)
	for k, v := range originalSeatTaken {
		seatTaken[k] = v
	}

	visited = make(map[string]struct{})
	visited[stringifySeating(seatTaken)] = struct{}{}
	iteration = 0
	looped = false
	for !looped {
		seatTaken = performIteration(seats, seatTaken, maxDist, 5, viewSeatNeighbors)
		iteration++
		seatString := stringifySeating(seatTaken)
		_, found := visited[seatString]
		if found {
			looped = true
		} else {
			visited[seatString] = struct{}{}
		}
	}

	fmt.Println(len(seatTaken))

}
