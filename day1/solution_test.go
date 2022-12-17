package main

import (
	"aoc/util"
	"testing"
)

func TestTotalCalories(t *testing.T) {
	elf := Elf{[]int{1, 2, 3, 4, 5, 20}}
	result := elf.totalCalories()
	if result != 35 {
		t.Errorf("Total calories not correct, expected 35, got %d", result)
	}
}

func TestPart1Example(t *testing.T) {
	line := util.FileAsString("ex1")
	result := part1(parse(line))
	if result != 24000 {
		t.Error("Result not correct")
	}
}

func TestPart2Example(t *testing.T) {
	line := util.FileAsString("ex1")
	result := part2(parse(line))
	if result != 45000 {
		t.Error("Result not correct")
	}
}

func BenchmarkPart1(b *testing.B) {
	line := util.FileAsString("input")
	elves := parse(line)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(elves)
	}
}

func BenchmarkPart2(b *testing.B) {
	line := util.FileAsString("input")
	elves := parse(line)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(elves)
	}
}
