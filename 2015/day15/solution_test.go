package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotal(t *testing.T) {
	ingredients := parse(util.FileAsLines("ex1"))
	assert.Equal(t, 24, getTotal(ingredients, map[string]int{"Butterscotch": 1, "Cinnamon": 2}, false))
	assert.Equal(t, 0, getTotal(ingredients, map[string]int{"Butterscotch": 2, "Cinnamon": 1}, false))
	assert.Equal(t, 50000000, getTotal(ingredients, map[string]int{"Butterscotch": 50, "Cinnamon": 50}, false))
}

func TestGetTotalWithCalories(t *testing.T) {
	ingredients := parse(util.FileAsLines("ex1"))
	assert.Equal(t, 0, getTotal(ingredients, map[string]int{"Butterscotch": 1, "Cinnamon": 2}, true))
	assert.Equal(t, 0, getTotal(ingredients, map[string]int{"Butterscotch": 50, "Cinnamon": 50}, true))
	assert.Equal(t, 5667660, getTotal(ingredients, map[string]int{"Butterscotch": 49, "Cinnamon": 36}, true))
}

func TestPart1Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	ingredients := parse(input)
	assert.Equal(t, 62842880, part1(ingredients))
}

func TestPart2Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	ingredients := parse(input)
	assert.Equal(t, 57600000, part2(ingredients))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	ingredients := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(ingredients)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	ingredients := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(ingredients)
	}
}
