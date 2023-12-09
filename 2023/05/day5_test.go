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
	expected := 35
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 46
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestParseInput(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
	}

	result := parseInput(input)
	expected := Almanac{
		Seeds: []int{79, 14, 55, 13},
		ConvMaps: [][]ConvRule{
			{
				{50, 98, 2},
				{52, 50, 48},
			},
			{
				{0, 15, 37},
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestSplitOn(t *testing.T) {
	ab := Interval{2, 5}
	cd := Interval{5, 10}
	inter, exter := ab.splitOn(cd)
	expectedInter := Interval{}
	expectedExter := []Interval{{2, 5}}

	if inter != expectedInter {
		t.Errorf("Intersection is incorrect, got: %v, want: %v.", inter, expectedInter)
	}

	if !reflect.DeepEqual(exter, expectedExter) {
		t.Errorf("External part is incorrect, got: %v, want: %v.", exter, expectedExter)
	}

	ab = Interval{2, 10}
	cd = Interval{5, 7}
	inter, exter = ab.splitOn(cd)
	expectedInter = Interval{5, 7}
	expectedExter = []Interval{{2, 5}, {7, 10}}

	if inter != expectedInter {
		t.Errorf("Intersection is incorrect, got: %v, want: %v.", inter, expectedInter)
	}

	if !reflect.DeepEqual(exter, expectedExter) {
		t.Errorf("External part is incorrect, got: %v, want: %v.", exter, expectedExter)
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
