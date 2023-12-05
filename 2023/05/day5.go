package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/utils"
)

type RangeOfNumbers struct {
	destinationStart int
	sourceStart      int
	lenght           int
}

type Almanac struct {
	seeds          []int
	conversionMaps [][]RangeOfNumbers
}

func parseInput(input []string) Almanac {
	seeds := parseSeedsToInts(input[0])
	conversionMaps := [][]RangeOfNumbers{}
	i := 3
	conversionMap := []RangeOfNumbers{}
	for i < len(input) {
		line := input[i]
		if line == "" {
			conversionMaps = append(conversionMaps, conversionMap)
			conversionMap = []RangeOfNumbers{}
			i += 2
		} else {
			rangeNumbers := parseStringToInts(line)
			rangeOfNumbers := RangeOfNumbers{
				rangeNumbers[0],
				rangeNumbers[1],
				rangeNumbers[2],
			}
			conversionMap = append(conversionMap, rangeOfNumbers)
			i++
		}
	}
	conversionMaps = append(conversionMaps, conversionMap)
	almanac := Almanac{seeds, conversionMaps}
	return almanac
}

func parseStringToInts(numbersLine string) []int {
	numbers := []int{}
	numbersParts := strings.Fields(numbersLine)
	for _, numberStr := range numbersParts {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func parseSeedsToInts(seedsLine string) []int {
	numbersLine := strings.Split(seedsLine, ":")[1]
	seeds := parseStringToInts(numbersLine)
	return seeds
}

func convertSourceToDestination(source int, destinationMap []RangeOfNumbers) int {
	for _, rangeOfNumber := range destinationMap {
		sourceStart := rangeOfNumber.sourceStart
		destinationStart := rangeOfNumber.destinationStart
		lenght := rangeOfNumber.lenght
		sourceNoOffset := source - sourceStart
		if sourceNoOffset >= 0 && sourceNoOffset < lenght {
			return sourceNoOffset + destinationStart
		}
	}
	return source
}

func Part1(input []string) int {
	almanac := parseInput(input)
	sources := almanac.seeds
	destinations := make([]int, len(sources))
	for _, destinationMap := range almanac.conversionMaps {
		for i, source := range sources {
			destination := convertSourceToDestination(source, destinationMap)
			destinations[i] = destination
		}
		sources = destinations
	}
	min := sources[0]
	for _, value := range sources {
		if value < min {
			min = value
		}
	}
	return min
}

func Part2(input []string) int {
	return 0
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println("Answer 2 : ", Part2(input))
}
