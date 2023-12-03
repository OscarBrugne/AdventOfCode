package main

import (
	"fmt"
	"strconv"

	"AdventOfCode/utils"
)

func checkIfIsDigit(line string, index int) (bool, int) {
	if index >= len(line) {
		return false, 0
	}
	isDigit := ('0' <= line[index]) && (line[index] <= '9')
	if isDigit {
		return true, int(line[index] - '0')
	}
	return false, 0
}

func checkIfIsNumber(line string, index int) (bool, int) {
	isDigit, digit := checkIfIsDigit(line, index)
	if !isDigit {
		return false, 0
	}
	number := 0
	for isDigit {
		number = number*10 + digit
		index++
		isDigit, digit = checkIfIsDigit(line, index)
	}
	return true, number
}

func isAdjacentSymbol(input []string, iInput int, iLine int) bool {
	if (iInput < 0) || (iInput >= len(input)) {
		return false
	}
	line := input[iInput]
	if (iLine < 0) || (iLine >= len(line)) {
		return false
	}
	isDigit, _ := checkIfIsDigit(line, iLine)
	isAdjacent := (line[iLine] != '.') && !isDigit
	return isAdjacent
}

func isNumberAdjacentToASymbol(input []string, iInput int, iNumberStart int, iNumberEnd int) bool {
	for iInputSymbol := iInput - 1; iInputSymbol <= iInput+1; iInputSymbol++ {
		for iLineSymbol := iNumberStart - 1; iLineSymbol < iNumberEnd+1; iLineSymbol++ {
			if isAdjacentSymbol(input, iInputSymbol, iLineSymbol) {
				return true
			}
		}
	}
	return false
}

func isAsterisk(input []string, iInput int, iLine int) bool {
	if (iInput < 0) || (iInput >= len(input)) {
		return false
	}
	line := input[iInput]
	if (iLine < 0) || (iLine >= len(line)) {
		return false
	}
	return line[iLine] == '*'
}

func checkIfIsAsteriskNumber(input []string, iInput int, iNumberStart int, iNumberEnd int) (bool, int, int) {
	for iInputSymbol := iInput - 1; iInputSymbol <= iInput+1; iInputSymbol++ {
		for iLineSymbol := iNumberStart - 1; iLineSymbol < iNumberEnd+1; iLineSymbol++ {
			if isAsterisk(input, iInputSymbol, iLineSymbol) {
				return true, iInputSymbol, iLineSymbol
			}
		}
	}
	return false, 0, 0
}

func Part1(input []string) int {
	res := 0
	for iInput, line := range input {
		for iLine := 0; iLine < len(line); iLine++ {
			isNumber, number := checkIfIsNumber(line, iLine)
			if isNumber {
				NumberLen := len(strconv.Itoa(number))
				iNumberStart := iLine
				iNumberEnd := iLine + NumberLen
				isAdjacentToASymbol := isNumberAdjacentToASymbol(input, iInput, iNumberStart, iNumberEnd)
				if isAdjacentToASymbol {
					res += number
				}
				iLine += NumberLen
			}

		}
	}
	return res
}

func Part2(input []string) int {
	res := 0
	asteriskNumbers := map[[3]int][2]int{}
	for iInput, line := range input {
		for iLine := 0; iLine < len(line); iLine++ {
			isNumber, number := checkIfIsNumber(line, iLine)
			if isNumber {
				NumberLen := len(strconv.Itoa(number))
				iNumberStart := iLine
				iNumberEnd := iLine + NumberLen
				isAdjacentToAnAsterisk, iInputSymbol, iLineSymbol := checkIfIsAsteriskNumber(input, iInput, iNumberStart, iNumberEnd)
				if isAdjacentToAnAsterisk {
					key := [3]int{iInput, iNumberStart, number}
					value := [2]int{iInputSymbol, iLineSymbol}
					asteriskNumbers[key] = value
				}
				iLine += NumberLen
			}

		}
	}
	for key, value := range asteriskNumbers {
		for key2, value2 := range asteriskNumbers {
			if key != key2 && value == value2 {
				res += key[2] * key2[2]
			}
		}
	}
	return res / 2
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println("Answer 2 : ", Part2(input))
}
