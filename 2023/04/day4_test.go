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
	expected := 13
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 30
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestParseLine(t *testing.T) {
	inputLine := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	result := parseLine(inputLine)
	expected := Scratchcard{
		winningNumbers:   []int{41, 48, 83, 86, 17},
		scratchedNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
	}
	if !reflect.DeepEqual(result, expected) {
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
