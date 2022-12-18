package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	result := part1(getEx1())
	assert.Equal(t, 10605, result)
}

func TestPart2Example(t *testing.T) {
	result := part2(getEx1())
	assert.Equal(t, 2713310158, result)
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(getInput())
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(getInput())
	}
}
