package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalCalories(t *testing.T) {
	elf := Elf{[]int{1, 2, 3, 4, 5, 20}}
	assert.Equal(t, 35, elf.totalCalories())
}

func TestPart1Example(t *testing.T) {
	line := util.FileAsString("ex1")
	result := part1(parse(line))
	assert.Equal(t, 24000, result)
}

func TestPart2Example(t *testing.T) {
	line := util.FileAsString("ex1")
	result := part2(parse(line))
	assert.Equal(t, 45000, result)
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
