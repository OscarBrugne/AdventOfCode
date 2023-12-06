package main

import (
	"fmt"
	"strconv"
	"time"

	"AdventOfCode/utils"
)

type Symbol struct {
	iLine int
	iCol  int
}

type Number struct {
	iLine int
	iCol  int
	value int
}

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

func isSymbol(input []string, iLine int, iCol int) bool {
	if (iLine < 0) || (iLine >= len(input)) {
		return false
	}
	line := input[iLine]
	if (iCol < 0) || (iCol >= len(line)) {
		return false
	}
	isDigit, _ := checkIfIsDigit(line, iCol)
	isSymbol := (line[iCol] != '.') && !isDigit
	return isSymbol
}

func isNumberAdjacentToASymbol(input []string, iLine int, iNumberStart int, iNumberEnd int) bool {
	for iLineSymbol := iLine - 1; iLineSymbol <= iLine+1; iLineSymbol++ {
		for iColSymbol := iNumberStart - 1; iColSymbol < iNumberEnd+1; iColSymbol++ {
			if isSymbol(input, iLineSymbol, iColSymbol) {
				return true
			}
		}
	}
	return false
}

func isAsterisk(input []string, iLine int, iCol int) bool {
	if (iLine < 0) || (iLine >= len(input)) {
		return false
	}
	line := input[iLine]
	if (iCol < 0) || (iCol >= len(line)) {
		return false
	}
	return line[iCol] == '*'
}

func checkIfIsAsteriskNumber(input []string, iLine int, iNumberStart int, iNumberEnd int) (bool, int, int) {
	for iLineSymbol := iLine - 1; iLineSymbol <= iLine+1; iLineSymbol++ {
		for iColSymbol := iNumberStart - 1; iColSymbol < iNumberEnd+1; iColSymbol++ {
			if isAsterisk(input, iLineSymbol, iColSymbol) {
				return true, iLineSymbol, iColSymbol
			}
		}
	}
	return false, 0, 0
}

func Part1(input []string) int {
	res := 0
	for iLine, line := range input {
		for iCol := 0; iCol < len(line); iCol++ {
			isNumber, number := checkIfIsNumber(line, iCol)
			if isNumber {
				numberLen := len(strconv.Itoa(number))
				iNumberStart := iCol
				iNumberEnd := iCol + numberLen
				isAdjacentToASymbol := isNumberAdjacentToASymbol(input, iLine, iNumberStart, iNumberEnd)
				if isAdjacentToASymbol {
					res += number
				}
				iCol += numberLen
			}
		}
	}
	return res
}

func Part2(input []string) int {
	res := 0
	asteriskNumbers := map[Symbol][]Number{}
	for iLine, line := range input {
		for iCol := 0; iCol < len(line); iCol++ {
			isNumber, number := checkIfIsNumber(line, iCol)
			if isNumber {
				NumberLen := len(strconv.Itoa(number))
				iNumberStart := iCol
				iNumberEnd := iCol + NumberLen
				isAdjacentToAnAsterisk, iLineSymbol, iColSymbol := checkIfIsAsteriskNumber(input, iLine, iNumberStart, iNumberEnd)
				if isAdjacentToAnAsterisk {
					key := Symbol{iLineSymbol, iColSymbol}
					value := Number{iLine, iCol, number}
					asteriskNumbers[key] = append(asteriskNumbers[key], value)
				}
				iCol += NumberLen
			}
		}
	}
	for _, values := range asteriskNumbers {
		if len(values) == 2 {
			res += values[0].value * values[1].value
		}
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
