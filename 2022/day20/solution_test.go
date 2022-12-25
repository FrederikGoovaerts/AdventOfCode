package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestShift(t *testing.T) {
	assert.Equal(t, 1, getShortestShift(1, 20))
	assert.Equal(t, 0, getShortestShift(19, 20))
	assert.Equal(t, 0, getShortestShift(38, 20))
	assert.Equal(t, 2, getShortestShift(40, 20))
	assert.Equal(t, -1, getShortestShift(-1, 20))
	assert.Equal(t, 0, getShortestShift(-19, 20))
	assert.Equal(t, 0, getShortestShift(-38, 20))
	assert.Equal(t, -2, getShortestShift(-40, 20))
	assert.Equal(t, -7, getShortestShift(-140, 20))
}

func TestPart1Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part1(parse(lines))
	assert.Equal(t, 3, result)
}

func TestPart2Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part2(parse(lines))
	assert.Equal(t, 1623178306, result)
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
