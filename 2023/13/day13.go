package main

import (
	"fmt"
	"time"

	"AdventOfCode/utils"
)

type Mirror struct {
	Rows []int
	Cols []int
}

func parseInput(input []string) []Mirror {
	mirrors := []Mirror{}

	mirrorStr := []string{}
	for _, line := range input {
		if line == "" {
			mirrors = append(mirrors, parseMirror(mirrorStr))
			mirrorStr = []string{}
		} else {
			mirrorStr = append(mirrorStr, line)
		}
	}
	mirrors = append(mirrors, parseMirror(mirrorStr))

	return mirrors
}

func parseMirror(mirrorStr []string) Mirror {
	mirror := Mirror{
		Rows: make([]int, len(mirrorStr)),
		Cols: make([]int, len(mirrorStr[0])),
	}

	for y, line := range mirrorStr {
		for x, char := range line {
			mirror.Rows[y] <<= 1
			mirror.Cols[x] <<= 1
			if char == '#' {
				mirror.Rows[y]++
				mirror.Cols[x]++
			}
		}
	}

	return mirror
}

func getMirrorAxis(lines []int) int {
	for i := 1; i < len(lines); i++ {
		isMirror := true
		for j := 0; isMirror && j < min(i, len(lines)-i); j++ {
			if lines[i-1-j] != lines[i+j] {
				isMirror = false
			}
		}
		if isMirror {
			return i
		}
	}
	return 0
}

func Part1(input []string) int {
	mirrors := parseInput(input)
	res := 0
	for _, mirror := range mirrors {
		res += getMirrorAxis(mirror.Cols)
		res += getMirrorAxis(mirror.Rows) * 100
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
