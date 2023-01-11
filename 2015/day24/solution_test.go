package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	input := util.FileAsNumbers("ex1")
	assert.Equal(t, 99, part1(input))
}

func TestPart2Example(t *testing.T) {
	input := util.FileAsNumbers("ex1")
	assert.Equal(t, 44, part2(input))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsNumbers("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsNumbers("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
