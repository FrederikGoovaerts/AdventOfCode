package main

import (
	"aoc/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 2, part1([]string{">"}))
	assert.Equal(t, 4, part1([]string{"^", ">", "v", "<"}))
	assert.Equal(t, 2, part1([]string{"^", "v", "^", "v", "^", "v", "^", "v", "^", "v"}))
}
func TestPart2Examples(t *testing.T) {
	assert.Equal(t, 3, part2([]string{"^", "v"}))
	assert.Equal(t, 3, part2([]string{"^", ">", "v", "<"}))
	assert.Equal(t, 11, part2([]string{"^", "v", "^", "v", "^", "v", "^", "v", "^", "v"}))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsString("input")
	parsed := strings.Split(input, "")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(parsed)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsString("input")
	parsed := strings.Split(input, "")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(parsed)
	}
}
