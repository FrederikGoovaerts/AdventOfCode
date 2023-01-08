package main

type Boss struct {
	hp    int
	dmg   int
	armor int
}

var boss Boss = Boss{100, 8, 2}

type Item struct {
	cost  int
	dmg   int
	armor int
}

var weapons []Item = []Item{{8, 4, 0}, {10, 5, 0}, {25, 6, 0}, {40, 7, 0}, {74, 8, 0}}
var armor []Item = []Item{{13, 0, 1}, {31, 0, 2}, {53, 0, 3}, {75, 0, 4}, {102, 0, 5}}
var rings []Item = []Item{{25, 1, 0}, {50, 2, 0}, {100, 3, 0}, {20, 0, 1}, {40, 0, 2}, {80, 0, 3}}
