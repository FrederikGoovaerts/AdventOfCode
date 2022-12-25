package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part1(parse(lines))

	assert.Equal(t, 24, result)
}

func TestPart2Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part2(parse(lines))
	assert.Equal(t, 93, result)
}

func BenchmarkPart1(b *testing.B) {
	lines := util.FileAsLines("input")
	cave, bedrock := parse(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(cave, bedrock)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := util.FileAsLines("input")
	cave, bedrock := parse(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(cave, bedrock)
	}
}
