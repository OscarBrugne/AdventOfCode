package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"AdventOfCode/utils"
)

type Row struct {
	Springs        string
	ContinousGroup []int
}

func parseInput(input []string) (rows []Row) {
	rows = []Row{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		springs := parts[0]
		ints := parseStringToInts(parts[1])

		row := Row{
			Springs:        springs,
			ContinousGroup: ints,
		}
		rows = append(rows, row)
	}
	return rows
}

func parseStringToInts(numbersLine string) []int {
	numbers := []int{}
	numbersParts := strings.Split(numbersLine, ",")
	for _, numberStr := range numbersParts {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func getFirstIndex(str string, char rune) int {
	for i, c := range str {
		if c == char {
			return i
		}
	}
	return -1
}

func getContinousGroup(springs string) []int {
	group := []int{}
	count := 0
	for _, c := range springs {
		if c == '?' {
			return nil
		} else if c == '#' {
			count++
		} else {
			if count > 0 {
				group = append(group, count)
				count = 0
			}
		}
	}
	if count > 0 {
		group = append(group, count)
	}
	return group
}

func isValidArrangement(row Row) bool {
	group := getContinousGroup(row.Springs)

	if len(group) != len(row.ContinousGroup) {
		return false
	}

	for i, g := range group {
		if g != row.ContinousGroup[i] {
			return false
		}
	}
	return true
}

func countArrangements(row Row) int {
	i := getFirstIndex(row.Springs, '?')

	if i == -1 {
		if isValidArrangement(row) {
			return 1
		}
		return 0
	}

	newSprings1 := row.Springs[:i] + "." + row.Springs[i+1:]
	newSprings2 := row.Springs[:i] + "#" + row.Springs[i+1:]

	res1 := countArrangements(Row{Springs: newSprings1, ContinousGroup: row.ContinousGroup})
	res2 := countArrangements(Row{Springs: newSprings2, ContinousGroup: row.ContinousGroup})
	return res1 + res2
}

func unfoldRow(row Row, unfoldingFactor int) Row {
	newRow := Row{
		Springs:        row.Springs,
		ContinousGroup: row.ContinousGroup,
	}

	for i := 1; i < unfoldingFactor; i++ {
		newRow.Springs += "?" + row.Springs
		newRow.ContinousGroup = append(newRow.ContinousGroup, row.ContinousGroup...)
	}

	return newRow
}

func Part1(input []string) int {
	rows := parseInput(input)

	res := 0
	for _, row := range rows {
		res += countArrangements(row)
	}

	return res
}

func Part2(input []string) int {
	rows := parseInput(input)

	unfoldedRows := []Row{}
	for _, row := range rows {
		unfoldedRows = append(unfoldedRows, unfoldRow(row, 5))
	}

	res := 0
	for _, row := range unfoldedRows {
		res += countArrangements(row)
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
