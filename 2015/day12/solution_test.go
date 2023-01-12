package main

import (
	"aoc/util"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	assert.Equal(t, 5, count(json.RawMessage("5"), false))
	assert.Equal(t, 17, count(json.RawMessage("[4,15,-2]"), false))
	assert.Equal(t, 26, count(json.RawMessage("[4,[15,7]]"), false))
	assert.Equal(t, 18, count(json.RawMessage(`{"a":2, "b":{}, "c":{"sixteen": 16}}`), false))
	assert.Equal(t, 6, count(json.RawMessage(`[1,{"c":"red","b":2},3]`), false))

	assert.Equal(t, 6, count(json.RawMessage(`[1,2,3]`), true))
	assert.Equal(t, 4, count(json.RawMessage(`[1,{"c":"red","b":2},3]`), true))
	assert.Equal(t, 6, count(json.RawMessage(`[1,"red",5]`), true))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsString("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsString("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
