package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSymbolAtBasic(t *testing.T) {
	assert.Equal(t, "1", getSymbolAtBasic(-1, -1))
	assert.Equal(t, "2", getSymbolAtBasic(0, -1))
	assert.Equal(t, "3", getSymbolAtBasic(1, -1))
	assert.Equal(t, "6", getSymbolAtBasic(1, 0))
	assert.Equal(t, "9", getSymbolAtBasic(1, 1))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	moves := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(moves)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	moves := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(moves)
	}
}
