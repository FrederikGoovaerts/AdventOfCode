package main

import (
	"testing"
)

func TestPart1Example1(t *testing.T) {
	result := solve(EX1, 4)
	if result != 7 {
		t.Error("Result not correct")
	}
}

func TestPart1Example2(t *testing.T) {
	result := solve(EX2, 4)
	if result != 5 {
		t.Error("Result not correct")
	}
}

func TestPart1Example4(t *testing.T) {
	result := solve(EX4, 4)
	if result != 10 {
		t.Error("Result not correct")
	}
}

func TestPart2Example1(t *testing.T) {
	result := solve(EX1, 14)
	if result != 19 {
		t.Error("Result not correct")
	}
}

func TestPart2Example2(t *testing.T) {
	result := solve(EX2, 14)
	if result != 23 {
		t.Error("Result not correct")
	}
}

func TestPart2Example4(t *testing.T) {
	result := solve(EX4, 14)
	if result != 29 {
		t.Error("Result not correct")
	}
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
