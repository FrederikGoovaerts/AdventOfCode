package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNext(t *testing.T) {
	assert.Equal(t, "abcdefgi", getNext("abcdefgh"))
	assert.Equal(t, "abcdefgj", getNext("abcdefgi"))
	assert.Equal(t, "abcdefha", getNext("abcdefgz"))
	assert.Equal(t, "abcdegaa", getNext("abcdefzz"))
	assert.Equal(t, "aaaaaaaa", getNext("zzzzzzzz"))
}

func BenchmarkGetNext(b *testing.B) {
	curr := "abcdefgh"
	for i := 0; i < b.N; i++ {
		curr = getNext(curr)
	}
}

func TestIsSecure(t *testing.T) {
	assert.Equal(t, false, isSecure("hijklmmn"))
	assert.Equal(t, false, isSecure("abbceffg"))
	assert.Equal(t, false, isSecure("abbcegjk"))
	assert.Equal(t, true, isSecure("abcdffaa"))
	assert.Equal(t, true, isSecure("ghjaabcc"))
}

func BenchmarkSolve(b *testing.B) {
	curr := input
	for i := 0; i < b.N; i++ {
		curr = solve(curr)
	}
}
