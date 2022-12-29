package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	names, distances := parse(input)
	assert.Equal(t, 330, part1(names, distances))
}

func TestPart2Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	names, distances := parse(input)
	assert.Equal(t, 286, part2(names, distances))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	names, distances := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(names, distances)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	names, distances := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(names, distances)
	}
}
