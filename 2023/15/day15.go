package main

import (
	"fmt"
	"time"

	"AdventOfCode/utils"
)

func Part1(input []string) int {
	line := input[0]

	res := 0
	current := 0
	for _, char := range line {
		if char == ',' {
			res += current
			current = 0
		} else {
			current += int(byte(char))
			current *= 17
			current %= 256
		}
	}
	res += current

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
