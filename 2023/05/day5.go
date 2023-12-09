package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"AdventOfCode/utils"
)

type Interval struct {
	Start int
	End   int
}

var EmptyInterval = Interval{}

func (ab Interval) splitOn(cd Interval) (inter Interval, exter []Interval) {
	a, b := ab.Start, ab.End
	c, d := cd.Start, cd.End

	isDisjoint := b <= c || d <= a
	if isDisjoint {
		exter = append(exter, ab)
		return
	}

	inter = Interval{max(a, c), min(b, d)}

	if a < c {
		before := Interval{a, c}
		exter = append(exter, before)
	}
	if b > d {
		after := Interval{d, b}
		exter = append(exter, after)
	}

	return
}

type ConvRule struct {
	DestStart   int
	SourceStart int
	Length      int
}

type Almanac struct {
	Seeds    []int
	ConvMaps [][]ConvRule
}

func parseInput(input []string) Almanac {
	seedsLine := strings.Split(input[0], ":")[1]
	seeds := parseStringToInts(seedsLine)

	convMaps := [][]ConvRule{}
	i := 3
	convMap := []ConvRule{}
	for i < len(input) {
		line := input[i]
		if line == "" {
			convMaps = append(convMaps, convMap)
			convMap = []ConvRule{}
			i += 2
		} else {
			nums := parseStringToInts(line)
			convRule := ConvRule{
				nums[0],
				nums[1],
				nums[2],
			}
			convMap = append(convMap, convRule)
			i++
		}
	}
	convMaps = append(convMaps, convMap)

	almanac := Almanac{
		Seeds:    seeds,
		ConvMaps: convMaps,
	}
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

func convertSourceToDestination(source int, convMap []ConvRule) int {
	for _, rule := range convMap {
		if (rule.SourceStart <= source) && (source < rule.SourceStart+rule.Length) {
			shift := rule.DestStart - rule.SourceStart
			return source + shift
		}
	}
	return source
}

func convertSourceToDestinations(source Interval, convMap []ConvRule) []Interval {
	sources := []Interval{source}
	destinations := []Interval{}

	for _, rule := range convMap {
		newSources := []Interval{}
		for len(sources) > 0 {
			currentSource := sources[0]
			sources = sources[1:]
			ruleSource := Interval{
				Start: rule.SourceStart,
				End:   rule.SourceStart + rule.Length,
			}

			inter, exter := currentSource.splitOn(ruleSource)

			newSources = append(newSources, exter...)
			if inter != EmptyInterval {
				shift := rule.DestStart - rule.SourceStart
				interShifted := Interval{
					Start: inter.Start + shift,
					End:   inter.End + shift,
				}
				destinations = append(destinations, interShifted)
			}
		}
		sources = append(sources, newSources...)
	}

	destinations = append(destinations, sources...)
	return destinations
}

func Part1(input []string) int {
	almanac := parseInput(input)

	sources := almanac.Seeds
	for _, convMap := range almanac.ConvMaps {
		destinations := make([]int, len(sources))
		for i, source := range sources {
			destination := convertSourceToDestination(source, convMap)
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
	for i := 0; i < len(almanac.Seeds); i += 2 {
		sourceStart := almanac.Seeds[i]
		sourceLength := almanac.Seeds[i+1]
		sourceEnd := sourceStart + sourceLength
		source := Interval{sourceStart, sourceEnd}
		sources = append(sources, source)
	}

	for _, convMap := range almanac.ConvMaps {
		destinations := []Interval{}
		for _, source := range sources {
			convDestinations := convertSourceToDestinations(source, convMap)
			destinations = append(destinations, convDestinations...)
		}
		sources = destinations
	}

	min := sources[0].Start
	for _, interval := range sources {
		if interval.Start < min {
			min = interval.Start
		}
	}
	return min
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
