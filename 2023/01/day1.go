package main

import (
	"fmt"
	"strings"
	"unicode"

	"AdventOfCode/utils"
)

func findDigit(line string, start int, end int, step int) int {
	i := start
	for (step > 0 && i < end) || (step < 0 && i > end) {
		var char rune = rune(line[i])
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
		i += step
	}
	return 0
}

func firstDigit(line string) int {
	return findDigit(line, 0, len(line), 1)
}

func lastDigit(line string) int {
	return findDigit(line, len(line)-1, -1, -1)
}

func replaceLetterDigits(line string) string {
	letterDigits := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	replacedLine := line
	for key, element := range letterDigits {
		replacedLine = strings.Replace(replacedLine, key, element, -1)
	}
	return replacedLine
}

func replaceLetterDigitsInInput(input []string) []string {
	replacedInput := []string{}
	for _, line := range input {
		replacedInput = append(replacedInput, replaceLetterDigits(line))
	}
	return replacedInput
}

func Part1(input []string) int {
	res := 0
	for _, line := range input {
		value := 10*firstDigit(line) + lastDigit(line)
		res += value
	}
	return res
}

func Part2(input []string) int {
	newInput := replaceLetterDigitsInInput(input)
	return Part1(newInput)
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println("Answer 2 : ", Part2(input))
}
