package main

import (
	"fmt"
	"time"

	"AdventOfCode/utils"
)

type Coord struct {
	ILine int
	ICol  int
}

type Number struct {
	ILine int
	ICol  int
	Value int
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

func isInBounds(input []string, iLine, iCol int) bool {
	if iLine < 0 || iLine >= len(input) {
		return false
	}
	line := input[iLine]
	return iCol >= 0 && iCol < len(line)
}

func isSymbol(input []string, iLine, iCol int) bool {
	if !isInBounds(input, iLine, iCol) {
		return false
	}

	line := input[iLine]
	isDigit, _ := checkIfIsDigit(line, iCol)
	return line[iCol] != '.' && !isDigit
}

func isAsterisk(input []string, iLine int, iCol int) bool {
	if !isInBounds(input, iLine, iCol) {
		return false
	}

	line := input[iLine]
	return line[iCol] == '*'
}

func isNumberAdjacentTo(
	input []string,
	iLine, iStartNum, iEndNum int,
	checkAdjacent func(input []string, i, j int) bool,
) (bool, []Coord) {
	isAdjacent := false
	coord := []Coord{}
	for iLineSymbol := iLine - 1; iLineSymbol <= iLine+1; iLineSymbol++ {
		for iColSymbol := iStartNum - 1; iColSymbol < iEndNum+1; iColSymbol++ {
			if checkAdjacent(input, iLineSymbol, iColSymbol) {
				isAdjacent = true
				coord = append(coord, Coord{iLineSymbol, iColSymbol})
			}
		}
	}

	return isAdjacent, coord
}

func isNumberAdjacentToASymbol(input []string, iLine int, iStartNum int, iEndNum int) bool {
	res, _ := isNumberAdjacentTo(input, iLine, iStartNum, iEndNum, isSymbol)
	return res
}

func isAnAsteriskNumber(input []string, iLine int, iStartNum int, iEndNum int) (bool, []Coord) {
	return isNumberAdjacentTo(input, iLine, iStartNum, iEndNum, isAsterisk)
}

func lenInt(num int) int {
	if num == 0 {
		return 1
	}
	cnt := 0
	for num != 0 {
		num /= 10
		cnt++
	}
	return cnt
}

func Part1(input []string) int {
	res := 0

	for iLine, line := range input {
		iCol := 0
		for iCol < len(line) {
			isNumber, value := checkIfIsNumber(line, iCol)
			if isNumber {
				numberLen := lenInt(value)
				iStartNum := iCol
				iEndNum := iCol + numberLen

				isAdjacentToASymbol := isNumberAdjacentToASymbol(input, iLine, iStartNum, iEndNum)
				if isAdjacentToASymbol {
					res += value
				}

				iCol += numberLen
			} else {
				iCol++
			}
		}
	}

	return res
}

func Part2(input []string) int {
	asteriskNumbers := map[Coord][]Number{}

	for iLine, line := range input {
		iCol := 0
		for iCol < len(line) {
			isNumber, value := checkIfIsNumber(line, iCol)
			if isNumber {
				NumberLen := lenInt(value)
				iStartNum := iCol
				iEndNum := iCol + NumberLen

				isAdjacentToAnAsterisk, asteriks := isAnAsteriskNumber(input, iLine, iStartNum, iEndNum)
				if isAdjacentToAnAsterisk {
					for _, asterik := range asteriks {
						key := asterik
						number := Number{iLine, iCol, value}
						asteriskNumbers[key] = append(asteriskNumbers[key], number)
					}
				}

				iCol += NumberLen
			} else {
				iCol++
			}
		}
	}

	res := 0
	for _, values := range asteriskNumbers {
		if len(values) == 2 {
			res += values[0].Value * values[1].Value
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
