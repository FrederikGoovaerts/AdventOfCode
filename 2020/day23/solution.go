package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	cups    map[int]*cup
	current *cup
}

type cup struct {
	label int
	next  *cup
}

func (c cup) nextTimes(t int) *cup {
	curr := &c
	for i := 0; i < t; i++ {
		curr = curr.next
	}
	return curr
}

func (g game) nbCups() int {
	return len(g.cups)
}

func (g game) getCup(label int) *cup {
	return g.cups[label]
}

func (g game) printState() {
	fmt.Print("(" + fmt.Sprint(g.current.label) + ") ")
	curr := g.current.next
	for curr != g.current {
		fmt.Print(fmt.Sprint(curr.label) + " ")
		curr = curr.next
	}
	fmt.Println()
}

func (g *game) move() {
	curr := g.current
	pickStart := curr.next
	pickMid := curr.nextTimes(2)
	pickEnd := curr.nextTimes(3)
	impactedLabels := []int{pickStart.label, pickMid.label, pickEnd.label}
	nextCurr := pickEnd.next
	dest := curr.label
	nextDest := func() {
		dest--
		if dest < 1 {
			dest += g.nbCups()
		}
	}
	nextDest()
	for utils.ContainsInt(impactedLabels, dest) {
		nextDest()
	}
	destCup := g.getCup(dest)
	destFollowing := destCup.next

	// Do actual move
	curr.next = nextCurr
	destCup.next = pickStart
	pickEnd.next = destFollowing
	g.current = nextCurr
}

func (g game) getOrder() string {
	one := g.getCup(1)
	curr := one.next
	res := ""
	for curr != one {
		res += fmt.Sprint(curr.label)
		curr = curr.next
	}
	return res
}

func main() {
	// input := "389125467" // ex1
	input := "872495136" // input
	cups := make(map[int]*cup, 0)

	labels := strings.Split(input, "")
	firstLabelString, labels := labels[0], labels[1:]
	firstLabel, _ := strconv.Atoi(firstLabelString)
	firstCup := cup{firstLabel, nil}
	cups[firstLabel] = &firstCup
	prev := &firstCup
	for _, labelString := range labels {
		label, _ := strconv.Atoi(labelString)
		cup := cup{label, nil}
		prev.next = &cup
		prev = &cup
		cups[label] = &cup
	}
	prev.next = &firstCup

	g := game{cups, &firstCup}
	for i := 0; i < 100; i++ {
		g.move()
	}
	fmt.Println(g.getOrder())

	// Redo setup for part 2

	cups = make(map[int]*cup, 0)
	firstCup = cup{firstLabel, nil}
	cups[firstLabel] = &firstCup
	prev = &firstCup
	for _, labelString := range labels {
		label, _ := strconv.Atoi(labelString)
		cup := cup{label, nil}
		prev.next = &cup
		prev = &cup
		cups[label] = &cup
	}
	for i := len(cups) + 1; i <= 1000000; i++ {
		cup := cup{i, nil}
		prev.next = &cup
		prev = &cup
		cups[i] = &cup
	}
	prev.next = &firstCup
	g = game{cups, &firstCup}

	for i := 0; i < 10000000; i++ {
		g.move()
	}
	fmt.Println(g.getCup(1).next.label * g.getCup(1).nextTimes(2).label)
}
