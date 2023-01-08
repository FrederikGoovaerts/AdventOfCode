package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerWins(t *testing.T) {
	assert.True(t, playerWins(Boss{12, 7, 2}, []Item{{0, 2, 3}, {0, 3, 2}}, 8))
}

func BenchmarkPart1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(boss)
	}
}

func BenchmarkPart2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(boss)
	}
}
