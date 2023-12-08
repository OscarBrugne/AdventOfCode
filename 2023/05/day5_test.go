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
	result := parseInput(inputTest)
	expected := Almanac{
		seeds: []int{79, 14, 55, 13},
		conversionMaps: [][]RangeOfNumbers{
			{
				{50, 98, 2},
				{52, 50, 48},
			},
			{
				{0, 15, 37},
				{37, 52, 2},
				{39, 0, 15},
			},
			{
				{49, 53, 8},
				{0, 11, 42},
				{42, 0, 7},
				{57, 7, 4},
			},
			{
				{88, 18, 7},
				{18, 25, 70},
			},
			{
				{45, 77, 23},
				{81, 45, 19},
				{68, 64, 13},
			},
			{
				{0, 69, 1},
				{1, 0, 69},
			},
			{
				{60, 56, 37},
				{56, 93, 4},
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestSplitOverlappingIntervals(t *testing.T) {
	souces := []Interval{{0, 10}}
	destinationMap := []RangeOfNumbers{{5, 5, 3}}
	result := splitOverlappingIntervals(souces, destinationMap)
	expected := []Interval{{0, 5}, {5, 8}, {8, 10}}
	if len(result) != len(expected) {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	if !reflect.DeepEqual(utils.NewSet[Interval](result...), utils.NewSet[Interval](expected...)) {
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
