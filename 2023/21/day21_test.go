package main

import (
	"testing"

	"AdventOfCode/utils"
)

var fileNameTest string = "input_test.txt"
var inputTest []string = utils.ReadFile(fileNameTest)

var fileNameDay string = "input.txt"
var inputDay []string = utils.ReadFile(fileNameDay)

func TestPart1(t *testing.T) {
	result := Part1(inputTest, 6)
	expected := 16
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest, 10)
	expected := 50
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part2(inputTest, 50)
	expected = 1594
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part2(inputTest, 500)
	expected = 167004
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(inputDay, 64)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(inputDay, 26501365)
	}
}
