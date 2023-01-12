package main

import (
	"aoc/util"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsString("input")
	moves := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(moves)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsString("input")
	moves := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(moves)
	}
}
