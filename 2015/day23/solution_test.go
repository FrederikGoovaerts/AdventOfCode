package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	ops := parse(input)
	assert.Equal(t, 2, part1(ops)["a"])
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	ops := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(ops)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	ops := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(ops)
	}
}
