package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"AdventOfCode/utils"
)

type Card struct {
	WinningNumbers   []int
	ScratchedNumbers []int
}

func parseLine(line string) Card {
	parts := strings.FieldsFunc(line, func(c rune) bool {
		return c == ':' || c == '|'
	})
	winningNumbers := parseStringToInts(parts[1])
	scratchedNumbers := parseStringToInts(parts[2])

	card := Card{
		WinningNumbers:   winningNumbers,
		ScratchedNumbers: scratchedNumbers,
	}
	return card
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

func countMatches(card Card) int {
	cnt := 0
	for _, scratchNumber := range card.ScratchedNumbers {
		for _, winningNumber := range card.WinningNumbers {
			if scratchNumber == winningNumber {
				cnt++
			}
		}
	}
	return cnt
}

func squareInt(x int) int {
	if x < 0 {
		return 1 << -x
	}
	return 1 << x
}

func calculatePoints(card Card) int {
	num := countMatches(card)
	if num == 0 {
		return 0
	}
	score := squareInt(num - 1)
	return score
}

func Part1(input []string) int {
	res := 0

	for _, line := range input {
		card := parseLine(line)
		score := calculatePoints(card)
		res += score
	}

	return res
}

func Part2(input []string) int {
	numCards := make([]int, len(input))
	for i := range numCards {
		numCards[i] = 1
	}

	for i, line := range input {
		card := parseLine(line)
		numMatches := countMatches(card)
		for j := 0; j < numMatches; j++ {
			numCards[i+1+j] += numCards[i]
		}
	}

	res := 0
	for _, numCard := range numCards {
		res += numCard
	}
	return res
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
