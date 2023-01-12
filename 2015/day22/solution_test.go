package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 226, part1(ex1State))
	assert.Equal(t, 641, part1(ex2State))
}

func BenchmarkPart1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(inputState)
	}
}

func BenchmarkPart2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(inputState)
	}
}
