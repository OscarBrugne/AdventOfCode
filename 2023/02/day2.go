package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"AdventOfCode/utils"
)

type Game struct {
	gameId  int
	subsets []map[string]int
}

func parseInput(lineInput string) Game {
	lineParts := strings.Split(lineInput, ": ")
	gameId := parseGameString(lineParts[0])
	subsets := parseSubsetsString(lineParts[1])
	game := Game{
		gameId:  gameId,
		subsets: subsets,
	}
	return game
}

func parseGameString(gameString string) int {
	gameIdStr := strings.Replace(gameString, "Game ", "", 1)
	gameId, err := strconv.Atoi(gameIdStr)
	if err != nil {
		panic(err)
	}
	return gameId
}

func parseSubsetsString(susbetsString string) []map[string]int {
	subsetsData := strings.Split(susbetsString, "; ")
	subsets := make([]map[string]int, len(subsetsData))
	for i, subsetElement := range subsetsData {
		subsetParts := strings.Split(subsetElement, ", ")
		subsets[i] = parseEachSubset(subsetParts)
	}
	return subsets
}

func parseEachSubset(subsetParts []string) map[string]int {
	subsetMap := map[string]int{}
	for _, cubeSet := range subsetParts {
		cubeSetParts := strings.Split(cubeSet, " ")
		numStr := cubeSetParts[0]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		color := cubeSetParts[1]
		subsetMap[color] = num
	}
	return subsetMap
}

func Part1(input []string) int {
	res := 0
	maxCubesInBag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, line := range input {
		game := parseInput(line)
		isGamePossible := true
		for _, subset := range game.subsets {
			for color, numCubes := range subset {
				if numCubes > maxCubesInBag[color] {
					isGamePossible = false
				}
			}
		}
		if isGamePossible {
			res += game.gameId
		}
	}
	return res
}

func Part2(input []string) int {
	res := 0
	for _, line := range input {
		game := parseInput(line)
		gamePower := 1
		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, subset := range game.subsets {
			for color, numCubes := range subset {
				if numCubes > minCubes[color] {
					minCubes[color] = numCubes
				}
			}
		}
		for _, numCubes := range minCubes {
			gamePower *= numCubes
		}
		res += gamePower
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
