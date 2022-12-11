package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	items     []int
	op        (func(int) int)
	check     int
	trueDest  int
	falseDest int
}

func (monkey *Monkey) addItem(item int) {
	monkey.items = append(monkey.items, item)
}

func (monkey *Monkey) pop() int {
	result := monkey.items[0]
	monkey.items = monkey.items[1:]
	return result
}

const ROUNDS = 20

func part1(monkeys []Monkey) int {
	monkeyBusiness := make([]int, len(monkeys))

	for round := 1; round <= ROUNDS; round++ {
		for id, monkey := range monkeys {
			monkeyBusiness[id] += len(monkey.items)

			for len(monkey.items) > 0 {
				item := monkey.pop()
				monkeys[id] = monkey
				updated := monkey.op(item) / 3
				if updated%monkey.check == 0 {
					monkeys[monkey.trueDest].addItem(updated)
				} else {
					monkeys[monkey.falseDest].addItem(updated)
				}
			}
		}
	}

	sort.Ints(monkeyBusiness)
	return monkeyBusiness[len(monkeys)-1] * monkeyBusiness[len(monkeys)-2]
}

func main() {
	monkeys := getInput()

	fmt.Println(part1(monkeys))

}
