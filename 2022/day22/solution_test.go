package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	ma, mo, _ := parse(lines, 4)
	result := part1(ma, mo)
	assert.Equal(t, 6032, result)
}

func TestPart2Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part2(parse(lines, 4))
	assert.Equal(t, 5031, result)
}

// There's quite some logic in the parsing so we add it to the benchmark

func BenchmarkPart1(b *testing.B) {
	lines := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ma, mo, _ := parse(lines, 50)
		part1(ma, mo)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(parse(lines, 50))
	}
}
