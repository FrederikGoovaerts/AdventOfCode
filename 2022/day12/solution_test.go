package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part1(parse(lines))
	assert.Equal(t, 31, result)
}

func TestPart2Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	board, _, end := parse(lines)
	result := part2(board, end)
	assert.Equal(t, 29, result)
}

func BenchmarkPart1(b *testing.B) {
	lines := util.FileAsLines("input")
	board, start, end := parse(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(board, start, end)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := util.FileAsLines("input")
	board, _, end := parse(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(board, end)
	}
}
