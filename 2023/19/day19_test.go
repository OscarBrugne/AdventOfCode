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
	result := Part1(inputTest)
	expected := 19114
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 167409079868000
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestApplyWorkflow(t *testing.T) {
	part := Part{
		X: 787,
		M: 2655,
		A: 1222,
		S: 2876,
	}
	workflows := Workflows{
		"in":  []Rule{{'s', '<', 1351, "px"}, {0, 0, 0, "qqz"}},
		"qqz": []Rule{{'s', '>', 2770, "qs"}, {'m', '>', 1801, "hdj"}, {0, 0, 0, "R"}},
		"qs":  []Rule{{'s', '>', 3448, "A"}, {0, 0, 0, "lnx"}},
		"lnx": []Rule{{'m', '>', 1548, "A"}, {0, 0, 0, "A"}},
		"hdj": []Rule{{0, 0, 0, "R"}},
	}
	workflowName := "in"
	result := applyWorkflow(part, workflows, workflowName)
	expected := true

	if result != expected {
		t.Errorf("Result is incorrect, got: %t, want: %t.", result, expected)
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
