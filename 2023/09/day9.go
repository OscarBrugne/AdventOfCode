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

func Part1(input []string) int {
	histories := parseInput(input)
	res := 0

	for _, history := range histories {
		extrapolationsSlices := [][]int{}
		extrapolationsSlices = append(extrapolationsSlices, history)

		i := 0
		for !allZeros(extrapolationsSlices[i]) {
			i++
			extrapolationsSlices = append(extrapolationsSlices, []int{})
			previousExtrapolations := extrapolationsSlices[i-1]
			for j := 1; j < len(previousExtrapolations); j++ {
				extrapolation := previousExtrapolations[j] - previousExtrapolations[j-1]
				extrapolationsSlices[i] = append(extrapolationsSlices[i], extrapolation)
			}
		}

		prediction := 0
		for i := 0; i < len(extrapolationsSlices); i++ {
			prediction += extrapolationsSlices[i][len(extrapolationsSlices[i])-1]
		}

		res += prediction
	}

	return res
}

func Part2(input []string) int {
	res := 0
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
