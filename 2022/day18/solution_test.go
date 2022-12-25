package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundsSurface(t *testing.T) {
	bounds := Bounds{0, 1, 2, 3, 6, 8}
	assert.Equal(t, 188, bounds.surface())
}

func TestPart1Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part1(parse(lines))
	assert.Equal(t, 64, result)
}

func TestPart2Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part2(parse(lines))
	assert.Equal(t, 58, result)
}

func BenchmarkPart1(b *testing.B) {
	lines := util.FileAsLines("input")
	parsed := parse(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(parsed)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := util.FileAsLines("input")
	parsed := parse(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(parsed)
	}
}
