package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	total, containers := parse(input)
	assert.Equal(t, 4, part1(total, containers))
}

func TestPart2Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	total, containers := parse(input)
	assert.Equal(t, 3, part2(total, containers))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	total, containers := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(total, containers)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	total, containers := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(total, containers)
	}
}
