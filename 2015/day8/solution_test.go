package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDecodeDifference(t *testing.T) {
	assert.Equal(t, 3, getDecodeDifference(`"aaa\"aaa"`))
	assert.Equal(t, 8, getDecodeDifference(`"\x27\"abcdef\\\""`))
}

func TestGetEncodeDifference(t *testing.T) {
	assert.Equal(t, 6, getEncodeDifference(`"aaa\"aaa"`))
	assert.Equal(t, 11, getEncodeDifference(`"\x27\"abcdef\\\""`))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
