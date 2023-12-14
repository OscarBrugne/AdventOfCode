package main

import (
	"fmt"
	"time"

	"AdventOfCode/utils"
)

type Coord struct {
	X int
	Y int
}

type Grid struct {
	Width  int
	Height int
	Data   map[Coord]byte
}

const (
	Empty     byte = '.'
	CubicRock byte = '#'
	RoundRock byte = 'O'
)

func buildGrids(input []string) Grid {
	grid := Grid{
		Width:  len(input[0]),
		Height: len(input),
		Data:   map[Coord]byte{},
	}

	for y, line := range input {
		for x, char := range line {
			if byte(char) != Empty {
				grid.Data[Coord{x, y}] = byte(char)
			}
		}
	}

	return grid
}

func (grid Grid) toString() string {
	var result string

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			if v, ok := grid.Data[coord]; ok {
				result += string(v)
			} else {
				result += string(Empty)
			}
		}
		result += "\n"
	}

	return result
}

func shiftRocksNorth(grid Grid) {
	for x := 0; x < grid.Width; x++ {
		for y := 1; y < grid.Height; y++ {
			if grid.Data[Coord{x, y}] == RoundRock {
				newY := y
				_, ok := grid.Data[Coord{x, newY - 1}]
				for !ok && newY >= 1 {
					grid.Data[Coord{x, newY - 1}] = RoundRock
					delete(grid.Data, Coord{x, newY})
					newY--
					_, ok = grid.Data[Coord{x, newY - 1}]
				}
			}
		}
	}
}

func calculateLoad(grid Grid) int {
	load := 0

	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			coord := Coord{x, y}
			if grid.Data[coord] == RoundRock {
				load += grid.Height - y
			}
		}
	}

	return load
}

func Part1(input []string) int {
	grid := buildGrids(input)
	shiftRocksNorth(grid)

	res := calculateLoad(grid)
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
