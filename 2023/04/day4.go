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
	parts := strings.Split(line, ": ")
	parts = strings.Split(parts[1], " | ")
	winningNumbers := parseNumbersLine(parts[0])
	scratchedNumbers := parseNumbersLine(parts[1])
	scratchcard := Scratchcard{
		winningNumbers:   winningNumbers,
		scratchedNumbers: scratchedNumbers,
	}
	return scratchcard
}

func parseNumbersLine(numbersLine string) []int {
	numbers := []int{}
	numbersParts := strings.Split(numbersLine, " ")
	for _, numberStr := range numbersParts {
		if numberStr != "" {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}
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
					score += score
				}

			}
		}
	}
	return score
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
	numCard := make([]int, len(input))
	for i := range numCard {
		numCard[i] = 1
	}
	for i, line := range input {
		scratchcard := parseLine(line)
		score := 0
		for _, scratchNumber := range scratchcard.scratchedNumbers {
			for _, winningNumber := range scratchcard.winningNumbers {
				if scratchNumber == winningNumber {
					score++
				}
			}
		}
		for j := 0; j < score; j++ {
			numCard[i+1+j] += numCard[i]
		}
		res += numCard[i]
	}
	return res
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println("Answer 2 : ", Part2(input))
}
