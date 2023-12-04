package main

import (
	"testing"

	"AdventOfCode/utils"
)

var fileName string = "input_test.txt"
var input []string = utils.ReadFile(fileName)

func TestPart1(t *testing.T) {
	result := Part1(input)
	expected := 13
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input)
	expected := 30
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
