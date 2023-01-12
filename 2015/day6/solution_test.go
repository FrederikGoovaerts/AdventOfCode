package main

import (
	"aoc/util"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	parsed := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(parsed)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	parsed := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(parsed)
	}
}
