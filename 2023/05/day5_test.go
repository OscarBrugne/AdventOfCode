package main

import (
	"reflect"
	"testing"

	"AdventOfCode/utils"
)

var fileName string = "input_test.txt"
var input []string = utils.ReadFile(fileName)

func TestPart1(t *testing.T) {
	result := Part1(input)
	expected := 35
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input)
	expected := 46
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestParseInput(t *testing.T) {
	result := parseInput(input)
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
