package main

import (
	"aoc/util"
	"testing"
)

func TestPart1Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part1(parse(lines, 3))
	if result != "CMZ" {
		t.Error("Result not correct")
	}
}

func TestPart2Example(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := part2(parse(lines, 3))
	if result != "MCD" {
		t.Error("Result not correct")
	}
}

func BenchmarkPart1(b *testing.B) {
	lines := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stacks, moves := parse(lines, 9)
		part1(stacks, moves)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stacks, moves := parse(lines, 9)
		part1(stacks, moves)
	}
}
