package main

import (
	"aoc/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsVowels(t *testing.T) {
	assert.Equal(t, false, containsVowels("bbdbfbdbfdb", 3))
	assert.Equal(t, true, containsVowels("abcde", 1))
	assert.Equal(t, true, containsVowels("abcde", 2))
	assert.Equal(t, false, containsVowels("abcde", 3))
	assert.Equal(t, true, containsVowels("riweoghnerogierwgowqgowqg", 3))
}

func TestContainsDouble(t *testing.T) {
	assert.Equal(t, true, containsDouble("bbdfd"))
	assert.Equal(t, false, containsDouble("abcdcba"))
	assert.Equal(t, true, containsDouble("abcddcba"))
	assert.Equal(t, true, containsDouble("abcdcbaa"))
	assert.Equal(t, true, containsDouble("aa"))
	assert.Equal(t, false, containsDouble("d"))
}

func TestIsNice(t *testing.T) {
	assert.Equal(t, true, isNice("ugknbfddgicrmopn"))
	assert.Equal(t, true, isNice("aaa"))
	assert.Equal(t, false, isNice("jchzalrnumimnmhp"))
	assert.Equal(t, false, isNice("haegwjzuvuyypxyu"))
	assert.Equal(t, false, isNice("dvszwmarrgswjxmb"))
}

func TestContainsDuplicate(t *testing.T) {
	precomputeRegexes()
	assert.Equal(t, true, containsDuplicate("xyxy"))
	assert.Equal(t, true, containsDuplicate("rgwerhgxyreqgxy"))
	assert.Equal(t, true, containsDuplicate("aabcdefgaa"))
	assert.Equal(t, false, containsDuplicate("aaa"))
	assert.Equal(t, false, containsDuplicate("xyyx"))
}

func TestIsNewNice(t *testing.T) {
	precomputeRegexes()
	assert.Equal(t, true, isNewNice("qjhvhtzxzqqjkmpb"))
	assert.Equal(t, true, isNewNice("xxyxx"))
	assert.Equal(t, false, isNewNice("uurcxstgmygtbstg"))
	assert.Equal(t, false, isNewNice("ieodomkazucvgmuy"))
}

func BenchmarkPart1(b *testing.B) {
	input := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.FileAsLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
