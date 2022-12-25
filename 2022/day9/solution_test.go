package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example1(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := solve(parse(lines), 2)
	assert.Equal(t, 13, result)
}

func TestPart2Example1(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := solve(parse(lines), 10)
	assert.Equal(t, 1, result)
}

func TestPart2Example2(t *testing.T) {
	lines := util.FileAsLines("ex2")
	result := solve(parse(lines), 10)
	assert.Equal(t, 36, result)
}

func BenchmarkPart1(b *testing.B) {
	lines := util.FileAsLines("input")
	parsed := parse(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(parsed, 2)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := util.FileAsLines("input")
	parsed := parse(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(parsed, 10)
	}
}
