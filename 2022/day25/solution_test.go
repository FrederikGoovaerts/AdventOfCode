package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeSnafu(t *testing.T) {
	assert.Equal(t, 1, decodeSnafu("1"))
	assert.Equal(t, 2, decodeSnafu("2"))
	assert.Equal(t, 3, decodeSnafu("1="))
	assert.Equal(t, 4, decodeSnafu("1-"))
	assert.Equal(t, 5, decodeSnafu("10"))
	assert.Equal(t, 6, decodeSnafu("11"))
	assert.Equal(t, 7, decodeSnafu("12"))
	assert.Equal(t, 8, decodeSnafu("2="))
	assert.Equal(t, 9, decodeSnafu("2-"))
	assert.Equal(t, 10, decodeSnafu("20"))
	assert.Equal(t, 15, decodeSnafu("1=0"))
	assert.Equal(t, 20, decodeSnafu("1-0"))
	assert.Equal(t, 2022, decodeSnafu("1=11-2"))
}
func TestEncodeSnafu(t *testing.T) {
	assert.Equal(t, "1", encodeSnafu(1))
	assert.Equal(t, "2", encodeSnafu(2))
	assert.Equal(t, "1=", encodeSnafu(3))
	assert.Equal(t, "1-", encodeSnafu(4))
	assert.Equal(t, "10", encodeSnafu(5))
	assert.Equal(t, "11", encodeSnafu(6))
	assert.Equal(t, "12", encodeSnafu(7))
	assert.Equal(t, "2=", encodeSnafu(8))
	assert.Equal(t, "2-", encodeSnafu(9))
	assert.Equal(t, "20", encodeSnafu(10))
	assert.Equal(t, "1=0", encodeSnafu(15))
	assert.Equal(t, "1-0", encodeSnafu(20))
	assert.Equal(t, "1=11-2", encodeSnafu(2022))
	assert.Equal(t, "1-0---0", encodeSnafu(12345))
	assert.Equal(t, "1121-1110-1=0", encodeSnafu(314159265))
}

func TestSolveExample(t *testing.T) {
	lines := util.FileAsLines("ex1")
	result := solve(lines)
	assert.Equal(t, "2=-1=0", result)
}

func BenchmarkSolve(b *testing.B) {
	lines := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(lines)
	}
}
