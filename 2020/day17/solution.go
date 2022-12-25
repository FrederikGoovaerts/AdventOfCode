package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cube struct {
	x int
	y int
	z int
}

type hypercube struct {
	x int
	y int
	z int
	w int
}

type change struct {
	x int
	y int
	z int
}

type hyperchange struct {
	x int
	y int
	z int
	w int
}

func neighbors(c cube) []cube {
	neighbors := make([]cube, 0)
	changes := make([]change, 0)
	increments := []int{-1, 0, 1}
	for _, x := range increments {
		for _, y := range increments {
			for _, z := range increments {
				if !(x == 0 && y == 0 && z == 0) {
					changes = append(changes, change{x, y, z})
				}
			}
		}
	}

	for _, dir := range changes {
		neighbors = append(neighbors, cube{c.x + dir.x, c.y + dir.y, c.z + dir.z})
	}
	return neighbors
}

func hyperneighbors(h hypercube) []hypercube {
	neighbors := make([]hypercube, 0)
	changes := make([]hyperchange, 0)
	increments := []int{-1, 0, 1}
	for _, x := range increments {
		for _, y := range increments {
			for _, z := range increments {
				for _, w := range increments {
					if !(x == 0 && y == 0 && z == 0 && w == 0) {
						changes = append(changes, hyperchange{x, y, z, w})
					}
				}
			}
		}
	}

	for _, dir := range changes {
		neighbors = append(neighbors, hypercube{h.x + dir.x, h.y + dir.y, h.z + dir.z, h.w + dir.w})
	}
	return neighbors
}

func performIteration(cubes map[cube]struct{}) map[cube]struct{} {
	cubesToCheck := make(map[cube]struct{}, 0)
	for cube := range cubes {
		cubesToCheck[cube] = struct{}{}
		for _, neigh := range neighbors(cube) {
			cubesToCheck[neigh] = struct{}{}
		}
	}

	newCubes := make(map[cube]struct{}, 0)
	for cube := range cubesToCheck {
		neighborsActive := 0
		_, cubeActive := cubes[cube]
		for _, neighbor := range neighbors(cube) {
			_, neighborActive := cubes[neighbor]
			if neighborActive {
				neighborsActive++
			}
		}
		if (cubeActive && (neighborsActive == 2 || neighborsActive == 3)) || (!cubeActive && neighborsActive == 3) {
			newCubes[cube] = struct{}{}
		}
	}
	return newCubes
}

func performHyperiteration(hypercubes map[hypercube]struct{}) map[hypercube]struct{} {
	cubesToCheck := make(map[hypercube]struct{}, 0)
	for cube := range hypercubes {
		cubesToCheck[cube] = struct{}{}
		for _, neigh := range hyperneighbors(cube) {
			cubesToCheck[neigh] = struct{}{}
		}
	}

	newCubes := make(map[hypercube]struct{}, 0)
	for cube := range cubesToCheck {
		neighborsActive := 0
		_, cubeActive := hypercubes[cube]
		for _, neighbor := range hyperneighbors(cube) {
			_, neighborActive := hypercubes[neighbor]
			if neighborActive {
				neighborsActive++
			}
		}
		if (cubeActive && (neighborsActive == 2 || neighborsActive == 3)) || (!cubeActive && neighborsActive == 3) {
			newCubes[cube] = struct{}{}
		}
	}
	return newCubes
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	cubes := make(map[cube]struct{}, 0)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				cube := cube{x, y, 0}
				cubes[cube] = struct{}{}
			}
		}
	}

	for i := 0; i < 6; i++ {
		cubes = performIteration(cubes)

	}
	fmt.Println(len(cubes))

	hypercubes := make(map[hypercube]struct{}, 0)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				cube := hypercube{x, y, 0, 0}
				hypercubes[cube] = struct{}{}
			}
		}
	}

	for i := 0; i < 6; i++ {
		hypercubes = performHyperiteration(hypercubes)

	}
	fmt.Println(len(hypercubes))

}
