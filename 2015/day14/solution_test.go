package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	reindeer := parse(input)
	assert.Equal(t, 1120, part1(reindeer, 1000))
}

func TestPart2Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	reindeer := parse(input)
	assert.Equal(t, 689, part2(reindeer, 1000))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	reindeer := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(reindeer, 2503)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	reindeer := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(reindeer, 2503)
	}
}
