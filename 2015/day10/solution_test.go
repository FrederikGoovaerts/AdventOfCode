package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNext(t *testing.T) {
	assert.Equal(t, "11", getNext("1"))
	assert.Equal(t, "21", getNext("11"))
	assert.Equal(t, "1211", getNext("21"))
	assert.Equal(t, "111221", getNext("1211"))
	assert.Equal(t, "312211", getNext("111221"))
}

// About 23 seconds
// func BenchmarkPart1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		part1(input)
// 	}
// }

func BenchmarkPart1Elements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1Elements(input_el)
	}
}

func BenchmarkPart2Elements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2Elements(input_el)
	}
}
