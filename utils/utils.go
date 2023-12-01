package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func ReadIntFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, ToInt(scanner.Text()))
	}
	return numbers
}

func StringsToIntegers(lines []string) []int {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		integers = append(integers, ToInt(line))
	}
	return integers
}
