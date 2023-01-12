package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	o, v := parse(input)
	assert.Equal(t, 65412, part1("h", o, v))
	assert.Equal(t, 65079, part1("i", o, v))
	assert.Equal(t, 72, part1("d", o, v))
	assert.Equal(t, 507, part1("e", o, v))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// We need to provide a clean copy
		o, v := parse(input)
		part1("a", o, v)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// We need to provide a clean copy
		o, v := parse(input)
		part2("a", o, v)
	}
}
