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
	expected := 1320
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 145
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestGetBoxes(t *testing.T) {
	input := []string{"rn=1", "cm-", "qp=3", "cm=2", "qp-", "pc=4", "ot=9", "ab=5", "pc-", "pc=6", "ot=7"}
	result := getBoxes(input)
	expected := map[int][]map[string]int{
		0: {
			{"rn": 1},
			{"cm": 2},
		},
		3: {
			{"ot": 7},
			{"ab": 5},
			{"pc": 6},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is incorrect, got:\n%s\nwant:\n%s.", toStringBoxes(result), toStringBoxes(expected))
	}
}

func TestCalculatePower(t *testing.T) {
	boxes := map[int][]map[string]int{
		0: {
			{"rn": 1},
			{"cm": 2},
		},
		3: {
			{"ot": 7},
			{"ab": 5},
			{"pc": 6},
		},
	}

	result := calculatePower(boxes)
	expected := 145

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
