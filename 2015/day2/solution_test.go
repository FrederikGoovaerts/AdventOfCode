package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWrappingWithSlack(t *testing.T) {
	assert.Equal(t, 58, getWrappingWithSlack([]int{2, 4, 3}))
	assert.Equal(t, 43, getWrappingWithSlack([]int{10, 1, 1}))
}

func TestGetRibbonAndBow(t *testing.T) {
	assert.Equal(t, 34, getRibbonAndBow([]int{2, 4, 3}))
	assert.Equal(t, 14, getRibbonAndBow([]int{10, 1, 1}))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	parsed := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(parsed)
	}
}

func BenchmarkPart1Concurrent(b *testing.B) {
	input := util.FileAsLines("input")
	parsed := parse(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1Concurrent(parsed)
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
