package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/utils"
)

type Scratchcard struct {
	winningNumbers   []int
	scratchedNumbers []int
}

func parseLine(line string) Scratchcard {
	parts := strings.Split(line, ":")
	parts = strings.Split(parts[1], "|")
	winningNumbers := parseStringToInts(parts[0])
	scratchedNumbers := parseStringToInts(parts[1])
	scratchcard := Scratchcard{
		winningNumbers:   winningNumbers,
		scratchedNumbers: scratchedNumbers,
	}
	return scratchcard
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

func calculatePoints(scratchcard Scratchcard) int {
	score := 0
	for _, scratchNumber := range scratchcard.scratchedNumbers {
		for _, winningNumber := range scratchcard.winningNumbers {
			if scratchNumber == winningNumber {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}
	return score
}

func calculateMatchingNumbers(scratchcard Scratchcard) int {
	numMatchingNumbers := 0
	for _, scratchNumber := range scratchcard.scratchedNumbers {
		for _, winningNumber := range scratchcard.winningNumbers {
			if scratchNumber == winningNumber {
				numMatchingNumbers++
			}
		}
	}
	return numMatchingNumbers
}

func Part1(input []string) int {
	res := 0
	for _, line := range input {
		scratchcard := parseLine(line)
		score := calculatePoints(scratchcard)
		res += score
	}
	return res
}

func Part2(input []string) int {
	res := 0
	numCards := make([]int, len(input))
	for i := range numCards {
		numCards[i] = 1
	}
	for i, line := range input {
		scratchcard := parseLine(line)
		numMatchingNumbers := calculateMatchingNumbers(scratchcard)
		for j := i + 1; j < i+1+numMatchingNumbers; j++ {
			numCards[j] += numCards[i]
		}
	}
	for _, numCard := range numCards {
		res += numCard
	}
	return res
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println("Answer 2 : ", Part2(input))
}
