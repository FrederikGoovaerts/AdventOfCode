package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example1(t *testing.T) {
	result := solve(EX1, 4)
	assert.Equal(t, 7, result)
}

func TestPart1Example2(t *testing.T) {
	result := solve(EX2, 4)
	assert.Equal(t, 5, result)
}

func TestPart1Example4(t *testing.T) {
	result := solve(EX4, 4)
	assert.Equal(t, 10, result)
}

func TestPart2Example1(t *testing.T) {
	result := solve(EX1, 14)
	assert.Equal(t, 19, result)
}

func TestPart2Example2(t *testing.T) {
	result := solve(EX2, 14)
	assert.Equal(t, 23, result)
}

func TestPart2Example4(t *testing.T) {
	result := solve(EX4, 14)
	assert.Equal(t, 29, result)
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(INPUT, 4)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(INPUT, 14)
	}
}
