package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoStep(t *testing.T) {
	input := util.FileAsLines("ex1")
	lights, size := parse(input)
	assert.Equal(t, 15, len(lights))
	lights = doStep(lights, size)
	assert.Equal(t, 11, len(lights))
	lights = doStep(lights, size)
	assert.Equal(t, 8, len(lights))
	lights = doStep(lights, size)
	assert.Equal(t, 4, len(lights))
	lights = doStep(lights, size)
	assert.Equal(t, 4, len(lights))
}

func TestPart1Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	lights, size := parse(input)
	assert.Equal(t, 4, part1(lights, size, 4))
}

func TestPart2Example(t *testing.T) {
	input := util.FileAsLines("ex1")
	lights, size := parse(input)
	assert.Equal(t, 17, part2(lights, size, 5))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	lights, size := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(lights, size, 100)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	lights, size := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(lights, size, 100)
	}
}
