package main

import (
	"reflect"
	"testing"

	"AdventOfCode/utils"
)

var fileNameTest string = "input_test.txt"
var inputTest []string = utils.ReadFile(fileNameTest)

var fileNameDay string = "input.txt"
var inputDay []string = utils.ReadFile(fileNameDay)

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := 21
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 0
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestGetContinousGroup(t *testing.T) {
	result := getContinousGroup("#.#.###")
	expected := []int{1, 1, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestCountArrangements(t *testing.T) {
	row := Row{
		Springs:        ".??..??...?##.",
		ContinousGroup: []int{1, 1, 3},
	}
	result := countArrangements(row)
	expected := 4

	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	row = Row{
		Springs:        "?#?#?#?#?#?#?#?",
		ContinousGroup: []int{1, 3, 1, 6},
	}
	result = countArrangements(row)
	expected = 1
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
