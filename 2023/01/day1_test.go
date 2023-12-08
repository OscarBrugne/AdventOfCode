package main

import (
	"testing"

	"AdventOfCode/utils"
)

var fileNameTest1 string = "input1_test.txt"
var inputTest1 []string = utils.ReadFile(fileNameTest1)

var fileNameTest2 string = "input2_test.txt"
var inputTest2 []string = utils.ReadFile(fileNameTest2)

var fileNameDay string = "input.txt"
var inputDay []string = utils.ReadFile(fileNameDay)

func TestPart1(t *testing.T) {
	result := Part1(inputTest1)
	expected := 142
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest2)
	expected := 281
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(inputDay)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(inputDay)
	}
}
