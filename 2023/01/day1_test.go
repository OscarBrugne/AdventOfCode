package main

import (
	"testing"

	"AdventOfCode/utils"
)

var fileName1 string = "input1_test.txt"
var input1 []string = utils.ReadFile(fileName1)

var fileName2 string = "input2_test.txt"
var input2 []string = utils.ReadFile(fileName2)

func TestPart1(t *testing.T) {
	result := Part1(input1)
	expected := 142
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input2)
	expected := 281
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
