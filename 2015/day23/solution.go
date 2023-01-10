package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Op struct {
	name     string
	register string
	offset   int
}

func (o Op) getNextInstructionOffset(r Registers) int {
	switch o.name {
	case "jmp":
		return o.offset
	case "jie":
		val := r[o.register]
		if val%2 == 0 {
			return o.offset
		}
	case "jio":
		val := r[o.register]
		if val == 1 {
			return o.offset
		}
	}
	return 1
}

type Registers map[string]int

func parse(lines []string) []Op {
	ops := make([]Op, 0, len(lines))

	for _, line := range lines {
		opString, remainder := line[0:3], line[4:]
		// fmt.Println("|" + opString + "|" + remainder + "|")
		op := Op{name: opString}
		switch opString {
		case "hlf":
			op.register = remainder
		case "tpl":
			op.register = remainder
		case "inc":
			op.register = remainder
		case "jmp":
			op.offset = util.StringToInt(remainder)
		case "jie":
			fallthrough
		case "jio":
			parts := strings.Split(remainder, ", ")
			op.register = parts[0]
			op.offset = util.StringToInt(parts[1])
		}
		ops = append(ops, op)
	}

	return ops
}

func part1(ops []Op) Registers {
	reg := Registers{"a": 0, "b": 0}
	instr := 0
	for instr >= 0 && instr < len(ops) {
		op := ops[instr]
		switch op.name {
		case "hlf":
			reg[op.register] = reg[op.register] / 2
		case "tpl":
			reg[op.register] = reg[op.register] * 3
		case "inc":
			reg[op.register] = reg[op.register] + 1
		}
		instr = instr + op.getNextInstructionOffset(reg)
	}
	return reg
}

func part2(ops []Op) Registers {
	reg := Registers{"a": 1, "b": 0}
	instr := 0
	for instr >= 0 && instr < len(ops) {
		op := ops[instr]
		switch op.name {
		case "hlf":
			reg[op.register] = reg[op.register] / 2
		case "tpl":
			reg[op.register] = reg[op.register] * 3
		case "inc":
			reg[op.register] = reg[op.register] + 1
		}
		instr = instr + op.getNextInstructionOffset(reg)
	}
	return reg
}

func main() {
	// input := util.FileAsLines("ex1")
	input := util.FileAsLines("input")
	ops := parse(input)

	// fmt.Println(ops)

	fmt.Println(part1(ops))
	fmt.Println(part2(ops))
}
