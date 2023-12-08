package main

import (
	"reflect"
	"testing"

	"AdventOfCode/utils"
)

var fileName1 string = "input1_test.txt"
var input1 []string = utils.ReadFile(fileName1)

var fileName2 string = "input2_test.txt"
var input2 []string = utils.ReadFile(fileName2)

func TestPart1(t *testing.T) {
	result := Part1(input1)
	expected := 2
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input2)
	expected := 6
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestParseInput(t *testing.T) {
	input := []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (CCC, AAA)",
	}

	result := parseInput(input)
	expected := Network{
		Instructions: "RL",
		Nodes: map[string][2]string{
			"AAA": {"BBB", "CCC"},
			"BBB": {"CCC", "AAA"},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input1)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(input2)
	}
}
