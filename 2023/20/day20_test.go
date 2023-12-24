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
	expected := 32000000
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part1(inputTest2)
	expected = 11687500
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest1)
	expected := 0
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part2(inputTest2)
	expected = 0
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestConjunction(t *testing.T) {
	startPulse := Pulse{
		value:    Low,
		fromName: "button",
		toName:   "broadcaster",
	}
	numCycle := 1

	inputStr := []string{
		"broadcaster -> a, b",
		"%a -> con",
		"%b -> con",
		"&con -> output",
	}
	modules := parseInput(inputStr)

	cntLow, cntHigh := pushButton(modules, startPulse, numCycle)
	expectedCntLow, expectedCntHigh := 4, 3
	if cntLow != expectedCntLow {
		t.Errorf("Number of -low is incorrect, got: %d, want: %d.", cntLow, expectedCntLow)
	}
	if cntHigh != expectedCntHigh {
		t.Errorf("Number of -high is incorrect, got: %d, want: %d.", cntHigh, expectedCntHigh)
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
