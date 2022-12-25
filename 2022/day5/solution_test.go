package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part1(parse(lines, 3))
	assert.Equal(t, "CMZ", result)
}

func TestPart2Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part2(parse(lines, 3))
	assert.Equal(t, "MCD", result)
}

func BenchmarkPart1(b *testing.B) {
	lines := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stacks, moves := parse(lines, 9)
		part1(stacks, moves)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stacks, moves := parse(lines, 9)
		part1(stacks, moves)
	}
}
