package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 0, part1(ex1a))
	assert.Equal(t, 0, part1(ex1b))
	assert.Equal(t, 3, part1(ex2a))
	assert.Equal(t, 3, part1(ex3))
	assert.Equal(t, -1, part1(ex4a))
	assert.Equal(t, -3, part1(ex5a))
	assert.Equal(t, -3, part1(ex5b))
}

func TestPart2Examples(t *testing.T) {
	assert.Equal(t, 1, part2(")"))
	assert.Equal(t, 5, part2("()())"))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsString("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsString("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
