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

type Interval struct {
	start int
	end   int
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
		sourceNoOffset := source - rangeOfNumber.sourceStart
		if sourceNoOffset >= 0 && sourceNoOffset < rangeOfNumber.lenght {
			return sourceNoOffset + rangeOfNumber.destinationStart
		}
	}
	return source
}

func splitOverlappingIntervals(sources []Interval, destinationMap []RangeOfNumbers) []Interval {
	overlappingSources := make([]Interval, len(sources))
	copy(overlappingSources, sources)
	nonOverlappingSources := []Interval{}
	for len(overlappingSources) > 0 {
		source := overlappingSources[0]
		overlappingSources = overlappingSources[1:]
		isSplit := false
		i := 0
		for !isSplit && i < len(destinationMap) {
			rangeOfNumber := destinationMap[i]
			rangeSource := Interval{rangeOfNumber.sourceStart, rangeOfNumber.sourceStart + rangeOfNumber.lenght}
			intersectionStart := max(source.start, rangeSource.start)
			intersectionEnd := min(source.end, rangeSource.end)
			isOverlapping := intersectionStart < intersectionEnd
			if isOverlapping {
				nonOverlappingSources = append(nonOverlappingSources, Interval{intersectionStart, intersectionEnd})
				if source.start < rangeSource.start {
					overlappingSources = append(overlappingSources, Interval{source.start, rangeSource.start})
				}
				if source.end > rangeSource.end {
					overlappingSources = append(overlappingSources, Interval{rangeSource.end, source.end})
				}
				isSplit = true
			}
			i++
		}
		if !isSplit {
			nonOverlappingSources = append(nonOverlappingSources, source)
		}
	}
	return nonOverlappingSources
}

func convertSourceToDestination2(source Interval, destinationMap []RangeOfNumbers) Interval {
	for _, rangeOfNumber := range destinationMap {
		rangeOfNumberSourceEnd := rangeOfNumber.sourceStart + rangeOfNumber.lenght
		if source.start >= rangeOfNumber.sourceStart && source.end <= rangeOfNumberSourceEnd {
			shift := rangeOfNumber.destinationStart - rangeOfNumber.sourceStart
			return Interval{
				start: source.start + shift,
				end:   source.end + shift,
			}
		}
	}
	return source
}

func Part1(input []string) int {
	almanac := parseInput(input)
	sources := almanac.seeds
	for _, destinationMap := range almanac.conversionMaps {
		destinations := make([]int, len(sources))
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
	almanac := parseInput(input)
	sources := []Interval{}
	for i := 0; i < len(almanac.seeds); i += 2 {
		sourceStart := almanac.seeds[i]
		sourceLenght := almanac.seeds[i+1]
		sourceEnd := sourceStart + sourceLenght
		source := Interval{sourceStart, sourceEnd}
		sources = append(sources, source)
	}
	for _, destinationMap := range almanac.conversionMaps {
		nonOverlappingSources := splitOverlappingIntervals(sources, destinationMap)
		destinations := []Interval{}
		for _, source := range nonOverlappingSources {
			destination := convertSourceToDestination2(source, destinationMap)
			destinations = append(destinations, destination)
		}
		sources = destinations
	}
	min := sources[0].start
	for _, interval := range sources {
		if interval.start < min {
			min = interval.start
		}
	}
	return min
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println("Answer 2 : ", Part2(input))
}
