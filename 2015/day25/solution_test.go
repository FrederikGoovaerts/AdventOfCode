package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrdinalFor(t *testing.T) {
	assert.Equal(t, 1, getOrdinalFor(1, 1))
	assert.Equal(t, 5, getOrdinalFor(2, 2))
	assert.Equal(t, 18, getOrdinalFor(4, 3))
}

func BenchmarkPart1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}
