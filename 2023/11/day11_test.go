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
	expected := 374
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest, 10)
	expected := 1030
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part2(inputTest, 100)
	expected = 8410
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestExpandGrid(t *testing.T) {
	input := []string{
		"...#.",
		".....",
		"#....",
	}
	grid := buildGrid(input, Empty)
	result := expandGrid(grid, 2)
	expected := buildGrid(
		[]string{
			".....#..",
			"........",
			"........",
			"#.......",
		},
		Empty)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is incorrect, got:\n%s\nwant:\n %s.", result.toString(Empty), expected.toString(Empty))
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(inputDay)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(inputDay, 1000000)
	}
}
