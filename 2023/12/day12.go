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

var cache = map[string]int{}

func countArrangements(row Row, isFirstContiguous bool) int {
	cacheKey := fmt.Sprintf("%v-%v-%v", row.Springs, row.ContinousGroup, isFirstContiguous)
	if val, ok := cache[cacheKey]; ok {
		return val
	}

	if row.Springs == "" {
		if len(row.ContinousGroup) == 0 {
			cache[cacheKey] = 1
			return 1
		}
		if len(row.ContinousGroup) == 1 && row.ContinousGroup[0] == 0 {
			cache[cacheKey] = 1
			return 1
		}
		cache[cacheKey] = 0
		return 0
	}

	switch row.Springs[0] {
	case '#':
		if len(row.ContinousGroup) == 0 || row.ContinousGroup[0] == 0 {
			cache[cacheKey] = 0
			return 0
		}
		row.ContinousGroup[0]--
		result := countArrangements(Row{Springs: row.Springs[1:], ContinousGroup: row.ContinousGroup}, true)
		cache[cacheKey] = result
		return result
	case '.':
		if len(row.ContinousGroup) == 0 {
			result := countArrangements(Row{Springs: row.Springs[1:], ContinousGroup: row.ContinousGroup}, false)
			cache[cacheKey] = result
			return result
		}
		if row.ContinousGroup[0] == 0 {
			result := countArrangements(Row{Springs: row.Springs[1:], ContinousGroup: row.ContinousGroup[1:]}, false)
			cache[cacheKey] = result
			return result
		}
		if isFirstContiguous {
			cache[cacheKey] = 0
			return 0
		}
		result := countArrangements(Row{Springs: row.Springs[1:], ContinousGroup: row.ContinousGroup}, false)
		cache[cacheKey] = result
		return result
	case '?':
		copyContinousGroup := make([]int, len(row.ContinousGroup))
		copy(copyContinousGroup, row.ContinousGroup)

		res1 := countArrangements(Row{Springs: "#" + row.Springs[1:], ContinousGroup: row.ContinousGroup}, true)
		res2 := countArrangements(Row{Springs: "." + row.Springs[1:], ContinousGroup: copyContinousGroup}, isFirstContiguous)
		result := res1 + res2
		cache[cacheKey] = result
		return result
	default:
		panic("Unknown spring")
	}
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
		res += countArrangements(row, false)
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
		res += countArrangements(row, false)
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
