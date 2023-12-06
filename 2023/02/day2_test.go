package main

import (
	"testing"

	"AdventOfCode/utils"
)

var fileName string = "input_test.txt"
var input []string = utils.ReadFile(fileName)

func TestPart1(t *testing.T) {
	result := Part1(input)
	expected := 8
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input)
	expected := 2286
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}
