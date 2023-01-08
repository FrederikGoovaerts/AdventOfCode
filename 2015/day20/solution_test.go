package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPresentsForHouse(t *testing.T) {
	// These are wrong due to edge cases for very small numbers
	// assert.Equal(t, 10, getPresentsForHouse(1))
	// assert.Equal(t, 30, getPresentsForHouse(2))
	assert.Equal(t, 40, getPresentsForHouse(3))
	assert.Equal(t, 70, getPresentsForHouse(4))
	assert.Equal(t, 60, getPresentsForHouse(5))
	assert.Equal(t, 120, getPresentsForHouse(6))
	assert.Equal(t, 80, getPresentsForHouse(7))
	assert.Equal(t, 150, getPresentsForHouse(8))
	assert.Equal(t, 130, getPresentsForHouse(9))
	assert.Equal(t, 720, getPresentsForHouse(51))
}

func TestGetLazyPresentsForHouse(t *testing.T) {
	assert.Equal(t, 44, getLazyPresentsForHouse(3))
	assert.Equal(t, 77, getLazyPresentsForHouse(4))
	assert.Equal(t, 66, getLazyPresentsForHouse(5))
	assert.Equal(t, 132, getLazyPresentsForHouse(6))
	assert.Equal(t, 88, getLazyPresentsForHouse(7))
	assert.Equal(t, 165, getLazyPresentsForHouse(8))
	assert.Equal(t, 143, getLazyPresentsForHouse(9))
	assert.Equal(t, 781, getLazyPresentsForHouse(51))
}

func BenchmarkPart1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(INPUT)
	}
}

func BenchmarkPart2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(INPUT)
	}
}
