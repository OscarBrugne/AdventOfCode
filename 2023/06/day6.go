package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"AdventOfCode/utils"
)

func parseInput1(input []string) (times, distances []int) {
	linesInts := [][]int{}
	for _, line := range input {
		parts := strings.Split(line, ":")
		numbers := parseStringToInts(parts[1])
		linesInts = append(linesInts, numbers)
	}
	return linesInts[0], linesInts[1]
}

func parseInput2(input []string) (t, d int) {
	ints := []int{}
	for _, line := range input {
		parts := strings.Split(line, ":")
		numberStr := strings.Replace(parts[1], " ", "", -1)
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		ints = append(ints, number)
	}
	return ints[0], ints[1]
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

func Part1(input []string) int {
	res := 1
	times, distances := parseInput1(input)
	for i := 0; i < len(times); i++ {
		num := 0
		t := times[i]
		d := distances[i]
		for timeHold := 0; timeHold < t; timeHold++ {
			if (timeHold * (t - timeHold)) > d {
				num++
			}
		}
		res *= num
	}
	return res
}

func Part2(input []string) int {
	res := 0
	t, d := parseInput2(input)
	for timeHold := 0; timeHold < t; timeHold++ {
		if (timeHold * (t - timeHold)) > d {
			res++
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
