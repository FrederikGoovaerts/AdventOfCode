package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Op int

const (
	On Op = iota
	Off
	Toggle
)

type Instruction struct {
	op     Op
	xStart int
	xEnd   int
	yStart int
	yEnd   int
}

var inputCoordRegex = regexp.MustCompile(" ([0-9]*),([0-9]*) through ([0-9]*),([0-9]*)$")

func parse(lines []string) []Instruction {
	instructions := make([]Instruction, 0, len(lines))

	for _, line := range lines {
		instr := Instruction{}
		match := inputCoordRegex.FindStringSubmatch(line)
		instr.xStart, _ = strconv.Atoi(match[1])
		instr.xEnd, _ = strconv.Atoi(match[3])
		instr.yStart, _ = strconv.Atoi(match[2])
		instr.yEnd, _ = strconv.Atoi(match[4])
		if strings.HasPrefix(line, "turn on") {
			instr.op = On
		} else if strings.HasPrefix(line, "turn off") {
			instr.op = Off
		} else {
			instr.op = Toggle
		}

		instructions = append(instructions, instr)
	}

	return instructions
}

type Row []bool

func (r Row) turnOn(startX, endX int) Row {
	for i := startX; i <= endX; i++ {
		r[i] = true
	}
	return r
}

func (r Row) turnOff(startX, endX int) Row {
	for i := startX; i <= endX; i++ {
		r[i] = false
	}
	return r
}

func (r Row) toggle(startX, endX int) Row {
	for i := startX; i <= endX; i++ {
		r[i] = !r[i]
	}
	return r
}

func (r Row) count() int {
	res := 0
	for i := 0; i < len(r); i++ {
		if r[i] {
			res++
		}
	}
	return res
}

func part1(instructions []Instruction) int {
	lights := make(map[int]Row)
	for i := 0; i < 1000; i++ {
		row := make([]bool, 1000)
		lights[i] = row
	}

	for _, instr := range instructions {
		for y := instr.yStart; y <= instr.yEnd; y++ {
			row := lights[y]
			if instr.op == On {
				lights[y] = row.turnOn(instr.xStart, instr.xEnd)
			} else if instr.op == Off {
				lights[y] = row.turnOff(instr.xStart, instr.xEnd)
			} else if instr.op == Toggle {
				lights[y] = row.toggle(instr.xStart, instr.xEnd)
			}
		}
	}

	result := 0

	for i := 0; i < 1000; i++ {
		result += lights[i].count()
	}

	return result
}

type AdvancedRow []int

func (r AdvancedRow) turnOn(startX, endX int) AdvancedRow {
	for i := startX; i <= endX; i++ {
		r[i] = r[i] + 1
	}
	return r
}

func (r AdvancedRow) turnOff(startX, endX int) AdvancedRow {
	for i := startX; i <= endX; i++ {
		r[i] = util.MaxInt(0, r[i]-1)
	}
	return r
}

func (r AdvancedRow) toggle(startX, endX int) AdvancedRow {
	for i := startX; i <= endX; i++ {
		r[i] = r[i] + 2
	}
	return r
}

func (r AdvancedRow) count() int {
	res := 0
	for i := 0; i < len(r); i++ {
		res += r[i]
	}
	return res
}

func part2(instructions []Instruction) int {
	lights := make(map[int]AdvancedRow)
	for i := 0; i < 1000; i++ {
		row := make([]int, 1000)
		lights[i] = row
	}

	for _, instr := range instructions {
		for y := instr.yStart; y <= instr.yEnd; y++ {
			row := lights[y]
			if instr.op == On {
				lights[y] = row.turnOn(instr.xStart, instr.xEnd)
			} else if instr.op == Off {
				lights[y] = row.turnOff(instr.xStart, instr.xEnd)
			} else if instr.op == Toggle {
				lights[y] = row.toggle(instr.xStart, instr.xEnd)
			}
		}
	}

	result := 0

	for i := 0; i < 1000; i++ {
		result += lights[i].count()
	}

	return result
}

func main() {
	input := util.FileAsLines("input")
	instructions := parse(input)

	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}
