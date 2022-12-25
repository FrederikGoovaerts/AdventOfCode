package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	line := util.FileAsLines("ex1")
	sensors, beacons, sX, lX := parse(line)
	result := part1(sensors, beacons, sX, lX, 10)
	assert.Equal(t, 26, result)
}

func TestPart2Example(t *testing.T) {
	line := util.FileAsLines("ex1")
	sensors, _, _, _ := parse(line)
	result := part2(sensors, 20)
	assert.Equal(t, 56000011, result)
}

func BenchmarkPart1(b *testing.B) {
	line := util.FileAsLines("input")
	sensors, beacons, sX, lX := parse(line)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(sensors, beacons, sX, lX, 2_000_000)
	}
}

func BenchmarkPart2(b *testing.B) {
	line := util.FileAsLines("input")
	sensors, _, _, _ := parse(line)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(sensors, 4_000_000)
	}
}
