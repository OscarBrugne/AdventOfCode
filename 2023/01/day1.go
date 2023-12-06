package main

import (
	"fmt"
	"time"

	"AdventOfCode/utils"
)

var spelledDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findDigit(line string, start int, end int, step int, isSpelledDigitAllowed bool) int {
	i := start
	for (step > 0 && i < end) || (step < 0 && i > end) {
		char := line[i]
		isDigit := ('0' <= char) && (char <= '9')
		if isDigit {
			return int(char - '0')
		}
		if isSpelledDigitAllowed {
			for word, value := range spelledDigits {
				isSpelledDigit := (i+len(word) <= len(line)) && (line[i:i+len(word)] == word)
				if isSpelledDigit {
					return value
				}
			}
		}
		i += step
	}
	return 0
}

func firstDigit(line string, isPart2 bool) int {
	return findDigit(line, 0, len(line), 1, isPart2)
}

func lastDigit(line string, isPart2 bool) int {
	return findDigit(line, len(line)-1, -1, -1, isPart2)
}

func calculateResult(input []string, isPart2 bool) int {
	res := 0
	for _, line := range input {
		value := 10*firstDigit(line, isPart2) + lastDigit(line, isPart2)
		res += value
	}
	return res
}

func Part1(input []string) int {
	return calculateResult(input, false)
}

func Part2(input []string) int {
	return calculateResult(input, true)
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
