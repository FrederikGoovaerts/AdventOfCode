package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Op int

const (
	Plus Op = iota
	Minus
	Mult
	Div
)

type Monkey struct {
	name      string
	isValue   bool
	value     int
	firstRef  string
	secondRef string
	operation Op
}

func getOperation(in string) Op {
	switch in {
	case "+":
		return Plus
	case "-":
		return Minus
	case "*":
		return Mult
	case "/":
		return Div
	}
	panic("Not an Op!")
}

func parse(lines []string) []Monkey {
	monkeys := make([]Monkey, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		result, err := strconv.Atoi(parts[1])
		if err != nil {
			// Monkey shouts operation
			operationParts := strings.Split(parts[1], " ")
			monkeys = append(monkeys, Monkey{parts[0], false, 0, operationParts[0], operationParts[2], getOperation(operationParts[1])})
		} else {
			// Monkey shouts value
			monkeys = append(monkeys, Monkey{parts[0], true, result, "", "", Plus})
		}

	}

	return monkeys
}

func getValue(name string, monkeys *map[string]Monkey, values *map[string]int) int {
	val, present := (*values)[name]
	if present {
		return val
	}
	m := (*monkeys)[name]
	calculatedValue := 0
	switch m.operation {
	case Plus:
		calculatedValue = getValue(m.firstRef, monkeys, values) + getValue(m.secondRef, monkeys, values)
	case Minus:
		calculatedValue = getValue(m.firstRef, monkeys, values) - getValue(m.secondRef, monkeys, values)
	case Mult:
		calculatedValue = getValue(m.firstRef, monkeys, values) * getValue(m.secondRef, monkeys, values)
	case Div:
		calculatedValue = getValue(m.firstRef, monkeys, values) / getValue(m.secondRef, monkeys, values)
	}
	(*values)[name] = calculatedValue

	return calculatedValue
}

func part1(monkeys []Monkey) int {
	values := make(map[string]int)
	monkeyMap := make(map[string]Monkey)
	for _, m := range monkeys {
		monkeyMap[m.name] = m
		if m.isValue {
			values[m.name] = m.value
		}
	}

	return getValue("root", &monkeyMap, &values)
}

func dependsOnHumn(name string, monkeys map[string]Monkey) bool {
	m := monkeys[name]
	if m.isValue {
		return m.name == "humn"
	} else {
		return dependsOnHumn(m.firstRef, monkeys) || dependsOnHumn(m.secondRef, monkeys)
	}
}

func getEqualizer(value int, monkey Monkey, monkeys *map[string]Monkey, values *map[string]int) int {
	if monkey.name == "humn" {
		return value
	}
	if dependsOnHumn(monkey.firstRef, *monkeys) {
		secondVal := getValue(monkey.secondRef, monkeys, values)
		switch monkey.operation {
		case Plus:
			return getEqualizer(value-secondVal, (*monkeys)[monkey.firstRef], monkeys, values)
		case Minus:
			return getEqualizer(value+secondVal, (*monkeys)[monkey.firstRef], monkeys, values)
		case Div:
			return getEqualizer(value*secondVal, (*monkeys)[monkey.firstRef], monkeys, values)
		case Mult:
			return getEqualizer(value/secondVal, (*monkeys)[monkey.firstRef], monkeys, values)
		}
	} else {
		firstVal := getValue(monkey.firstRef, monkeys, values)
		switch monkey.operation {
		case Plus:
			return getEqualizer(value-firstVal, (*monkeys)[monkey.secondRef], monkeys, values)
		case Minus:
			return getEqualizer(firstVal-value, (*monkeys)[monkey.secondRef], monkeys, values)
		case Div:
			return getEqualizer(firstVal/value, (*monkeys)[monkey.secondRef], monkeys, values)
		case Mult:
			return getEqualizer(value/firstVal, (*monkeys)[monkey.secondRef], monkeys, values)
		}
	}
	panic("uh")
}

func part2(monkeys []Monkey) int {
	values := make(map[string]int)
	monkeyMap := make(map[string]Monkey)
	for _, m := range monkeys {
		monkeyMap[m.name] = m
		if m.isValue {
			values[m.name] = m.value
		}
	}
	rootFirst := monkeyMap["root"].firstRef
	rootSecond := monkeyMap["root"].secondRef
	if dependsOnHumn(rootFirst, monkeyMap) {
		secondVal := getValue(rootSecond, &monkeyMap, &values)
		return getEqualizer(secondVal, monkeyMap[rootFirst], &monkeyMap, &values)
	} else {
		firstVal := getValue(rootFirst, &monkeyMap, &values)
		return getEqualizer(firstVal, monkeyMap[rootSecond], &monkeyMap, &values)

	}
}

func main() {
	// lines := util.FileAsLines("ex1")
	lines := util.FileAsLines("input")
	monkeys := parse(lines)

	fmt.Println(part1(monkeys))
	fmt.Println(part2(monkeys))
}
