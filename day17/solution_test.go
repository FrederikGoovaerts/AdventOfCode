package main

import (
	"aoc/util"
	"testing"
)

func TestPart1Example(t *testing.T) {
	line := util.FileAsString("ex1")
	result := part1(parse(line))
	if result != 3068 {
		t.Error("Result not correct")
	}
}

func TestPart2Example(t *testing.T) {
	line := util.FileAsString("ex1")
	result := part2(parse(line))
	if result != 1514285714288 {
		t.Error("Result not correct")
	}
}

func BenchmarkPart1(b *testing.B) {
	line := util.FileAsString("input")
	parsed := parse(line)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(parsed)
	}
}

func BenchmarkPart2(b *testing.B) {
	line := util.FileAsString("input")
	parsed := parse(line)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(parsed)
	}
}
