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

func (c1 Coord) Add(c2 Coord) Coord {
	return Coord{c1.X + c2.X, c1.Y + c2.Y}
}

type Grid struct {
	Width  int
	Height int
	Data   map[Coord]byte
}

func (coord Coord) isInBounds(grid Grid) bool {
	return 0 <= coord.X && coord.X < grid.Width && 0 <= coord.Y && coord.Y < grid.Height
}

const (
	Empty     byte = '.'
	CubicRock byte = '#'
	RoundRock byte = 'O'
)

var (
	North = Coord{0, -1}
	West  = Coord{-1, 0}
	South = Coord{0, 1}
	East  = Coord{1, 0}
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

func shiftSingleRock(grid Grid, coord Coord, dir Coord) {
	if grid.Data[coord] == RoundRock {
		current := coord
		before := coord.Add(dir)

		_, ok := grid.Data[before]
		for !ok && before.isInBounds(grid) {
			grid.Data[before] = RoundRock
			delete(grid.Data, current)

			current = before
			before = before.Add(dir)
			_, ok = grid.Data[before]
		}
	}
}

func shiftRocks(grid Grid, dir Coord) {
	switch dir {
	case North, West:
		for x := 0; x < grid.Width; x++ {
			for y := 0; y < grid.Height; y++ {
				shiftSingleRock(grid, Coord{x, y}, dir)
			}
		}

	case South, East:
		for x := grid.Width - 1; x >= 0; x-- {
			for y := grid.Height - 1; y >= 0; y-- {
				shiftSingleRock(grid, Coord{x, y}, dir)
			}
		}
	}
}

func cycleRocks(grid Grid) {
	shiftRocks(grid, North)
	shiftRocks(grid, West)
	shiftRocks(grid, South)
	shiftRocks(grid, East)
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
	shiftRocks(grid, North)

	return calculateLoad(grid)
}

func Part2(input []string) int {
	numCycles := 1000000000

	grid := buildGrids(input)
	cache := make(map[string]int)

	for i := 0; i < numCycles; i++ {
		gridStr := grid.toString()

		if iStartCycle, ok := cache[gridStr]; ok {
			remainingCycles := (numCycles - iStartCycle) % (i - iStartCycle)
			for j := 0; j < remainingCycles; j++ {
				cycleRocks(grid)
			}
			return calculateLoad(grid)
		}

		cache[gridStr] = i
		cycleRocks(grid)
	}

	return calculateLoad(grid)
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
