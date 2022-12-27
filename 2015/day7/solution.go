package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Operation interface {
	getValue(OperationMap, ValueMap) uint16
}

func getOrCalc(key string, o OperationMap, v ValueMap) uint16 {
	if v.has(key) {
		return v[key]
	} else {
		val := o[key].getValue(o, v)
		v[key] = val
		return val
	}
}

type EqOp struct {
	subject string
}

func (a *EqOp) getValue(o OperationMap, v ValueMap) uint16 {
	return getOrCalc(a.subject, o, v)
}

type AndOp struct {
	first  string
	second string
}

func (a *AndOp) getValue(o OperationMap, v ValueMap) uint16 {
	first := getOrCalc(a.first, o, v)
	second := getOrCalc(a.second, o, v)
	return first & second
}

type OrOp struct {
	first  string
	second string
}

func (a *OrOp) getValue(o OperationMap, v ValueMap) uint16 {
	first := getOrCalc(a.first, o, v)
	second := getOrCalc(a.second, o, v)
	return first | second
}

type NotOp struct {
	subject string
}

func (a *NotOp) getValue(o OperationMap, v ValueMap) uint16 {
	subject := getOrCalc(a.subject, o, v)
	return ^subject
}

type LShiftOp struct {
	subject string
	amount  uint16
}

func (a *LShiftOp) getValue(o OperationMap, v ValueMap) uint16 {
	subject := getOrCalc(a.subject, o, v)
	return subject << a.amount
}

type RShiftOp struct {
	subject string
	amount  uint16
}

func (a *RShiftOp) getValue(o OperationMap, v ValueMap) uint16 {
	subject := getOrCalc(a.subject, o, v)
	return subject >> a.amount
}

type OperationMap map[string]Operation
type ValueMap map[string]uint16

func (v ValueMap) has(key string) bool {
	_, present := v[key]
	return present
}

var valueRegex = regexp.MustCompile("^([0-9]*) -> ([a-z]*)$")

func parse(lines []string) (OperationMap, ValueMap) {
	o := make(OperationMap)
	v := make(ValueMap)
	v["1"] = 1

	for _, line := range lines {
		valueResult := valueRegex.FindStringSubmatch(line)
		if len(valueResult) > 0 {
			value, _ := strconv.Atoi(valueResult[1])
			v[valueResult[2]] = uint16(value)
		} else {
			outerParts := strings.Split(line, " -> ")
			target := outerParts[1]
			if strings.HasPrefix(line, "NOT ") {
				origin := outerParts[0][4:]
				o[target] = &NotOp{subject: origin}
			} else if strings.Contains(line, " AND ") {
				split := strings.Split(outerParts[0], " AND ")
				o[target] = &AndOp{split[0], split[1]}
			} else if strings.Contains(line, " OR ") {
				split := strings.Split(outerParts[0], " OR ")
				o[target] = &OrOp{split[0], split[1]}
			} else if strings.Contains(line, " LSHIFT ") {
				split := strings.Split(outerParts[0], " LSHIFT ")
				shift, _ := strconv.Atoi(split[1])
				o[target] = &LShiftOp{split[0], uint16(shift)}
			} else if strings.Contains(line, " RSHIFT ") {
				split := strings.Split(outerParts[0], " RSHIFT ")
				shift, _ := strconv.Atoi(split[1])
				o[target] = &RShiftOp{split[0], uint16(shift)}
			} else {
				o[target] = &EqOp{outerParts[0]}
			}
		}
	}

	return o, v
}

func part1(resultWire string, o OperationMap, v ValueMap) int {
	return int(getOrCalc(resultWire, o, v))
}

func part2(resultWire string, o OperationMap, v ValueMap) int {
	v["b"] = 3176
	return int(getOrCalc(resultWire, o, v))
}

func main() {
	// input, resultWire := util.FileAsLines("ex1"), "h"
	input, resultWire := util.FileAsLines("input"), "a"
	operations, values := parse(input)
	fmt.Println(part1(resultWire, operations, values))

	operations2, values2 := parse(input)
	fmt.Println(part2(resultWire, operations2, values2))
}
