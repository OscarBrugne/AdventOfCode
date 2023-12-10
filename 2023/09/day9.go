package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"AdventOfCode/utils"
)

func parseInput(input []string) (histories [][]int) {
	histories = [][]int{}
	for _, line := range input {
		numbers := parseStringToInts(line)
		histories = append(histories, numbers)
	}
	return histories
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

func allZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func calculateExtrapolation(history []int) []int {
	extrapolations := []int{}
	for i := 1; i < len(history); i++ {
		extrapolation := history[i] - history[i-1]
		extrapolations = append(extrapolations, extrapolation)
	}
	return extrapolations
}

func calculateExtrapolations(history []int) [][]int {
	extrapolationsSeries := [][]int{}
	extrapolationsSeries = append(extrapolationsSeries, history)

	for i := 1; i < len(history); i++ {
		previousExtrapolations := extrapolationsSeries[i-1]
		if allZeros(previousExtrapolations) {
			return extrapolationsSeries
		}

		extrapolations := calculateExtrapolation(previousExtrapolations)
		extrapolationsSeries = append(extrapolationsSeries, extrapolations)
	}

	return extrapolationsSeries
}

func Part1(input []string) int {
	histories := parseInput(input)
	res := 0

	for _, history := range histories {
		extrapolationsSeries := calculateExtrapolations(history)

		futurePrediction := 0
		for i := len(extrapolationsSeries) - 1; i > -1; i-- {
			futurePrediction = extrapolationsSeries[i][len(extrapolationsSeries[i])-1] + futurePrediction
		}

		res += futurePrediction
	}

	return res
}

func Part2(input []string) int {
	histories := parseInput(input)
	res := 0

	for _, history := range histories {
		extrapolationsSeries := calculateExtrapolations(history)

		pastPrediction := 0
		for i := len(extrapolationsSeries) - 1; i > -1; i-- {
			pastPrediction = extrapolationsSeries[i][0] - pastPrediction
		}

		res += pastPrediction
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
