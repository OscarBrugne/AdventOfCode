package main

import (
	"reflect"
	"testing"

	"AdventOfCode/utils"
)

var fileNameTest1 string = "input1_test.txt"
var inputTest1 []string = utils.ReadFile(fileNameTest1)
var fileNameTest2 string = "input2_test.txt"
var inputTest2 []string = utils.ReadFile(fileNameTest2)

var fileNameTest3 string = "input3_test.txt"
var inputTest3 []string = utils.ReadFile(fileNameTest3)
var fileNameTest4 string = "input4_test.txt"
var inputTest4 []string = utils.ReadFile(fileNameTest4)

var fileNameDay string = "input.txt"
var inputDay []string = utils.ReadFile(fileNameDay)

func TestPart1(t *testing.T) {
	result := Part1(inputTest1)
	expected := 4
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part1(inputTest2)
	expected = 8
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest3)
	expected := 8
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part2(inputTest4)
	expected = 10
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestBuildGrid(t *testing.T) {
	input := []string{
		"123",
		"456",
	}
	result := buildGrid(input)
	expected := Grid{
		Width:  3,
		Height: 2,
		Data: map[Coord]byte{
			{X: 0, Y: 0}: '1',
			{X: 1, Y: 0}: '2',
			{X: 2, Y: 0}: '3',
			{X: 0, Y: 1}: '4',
			{X: 1, Y: 1}: '5',
			{X: 2, Y: 1}: '6',
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
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
