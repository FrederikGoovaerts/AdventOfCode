package main

import (
	"aoc/util"
	"testing"
)

func TestPart1Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part1(parse(lines))
	if result != 1651 {
		t.Error("Result not correct")
	}
}

func TestPart2Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part2(parse(lines))
	if result != 1707 {
		t.Error("Result not correct")
	}
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
