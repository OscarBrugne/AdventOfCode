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
	expected := 46
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 51
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestNextBeam(t *testing.T) {
	grid := buildGrid(inputTest)
	beam := Beam{
		Origin: Coord{X: 4, Y: 6},
		Dir:    North,
	}
	result := nextBeam(grid, beam)
	expected := []Beam{
		{
			Coord{X: 5, Y: 6}, East,
		},
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
