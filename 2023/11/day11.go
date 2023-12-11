package main

import (
	"fmt"
	"strings"
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

var Empty byte = '.'

func buildGrid(input []string, empty byte) Grid {
	grid := Grid{
		Width:  len(input[0]),
		Height: len(input),
		Data:   map[Coord]byte{},
	}

	for y, line := range input {
		for x, char := range line {
			if byte(char) != empty {
				grid.Data[Coord{x, y}] = byte(char)
			}
		}
	}

	return grid
}

func (grid Grid) toString(empty byte) string {
	var result strings.Builder

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			if v, ok := grid.Data[coord]; ok {
				result.WriteByte(v)
			} else {
				result.WriteByte(empty)
			}
		}
		result.WriteByte('\n')
	}

	return result.String()
}

func (grid Grid) getEmptyRows() []int {
	emptyRows := []int{}
	for y := 0; y < grid.Height; y++ {
		isEmpty := true

		x := 0
		for x < grid.Width {
			if _, ok := grid.Data[Coord{x, y}]; ok {
				isEmpty = false
			}
			x++
		}

		if isEmpty {
			emptyRows = append(emptyRows, y)
		}
	}
	return emptyRows
}

func (grid Grid) getEmptyCols() []int {
	emptyCols := []int{}
	for x := 0; x < grid.Width; x++ {
		isEmpty := true

		y := 0
		for y < grid.Height {
			if _, ok := grid.Data[Coord{x, y}]; ok {
				isEmpty = false
			}
			y++
		}

		if isEmpty {
			emptyCols = append(emptyCols, x)
		}
	}
	return emptyCols
}

func calculateOffsets(emptyIndexes []int, bound int) []int {
	offsets := make([]int, bound)
	for _, idx := range emptyIndexes {
		for i := idx + 1; i < len(offsets); i++ {
			offsets[i]++
		}
	}
	return offsets
}

func expandGrid(grid Grid, expansionFactor int) Grid {
	emptyCols := grid.getEmptyCols()
	emptyRows := grid.getEmptyRows()
	numLinesToAdd := expansionFactor - 1

	newGrid := Grid{
		Width:  grid.Width + len(emptyCols)*numLinesToAdd,
		Height: grid.Height + len(emptyRows)*numLinesToAdd,
		Data:   make(map[Coord]byte, len(grid.Data)),
	}

	dXs := calculateOffsets(emptyCols, grid.Width)
	dYs := calculateOffsets(emptyRows, grid.Height)

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			if _, ok := grid.Data[coord]; ok {
				newCoord := Coord{x + dXs[x]*numLinesToAdd, y + dYs[y]*numLinesToAdd}
				newGrid.Data[newCoord] = grid.Data[coord]
			}
		}
	}

	return newGrid
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateLength(grid Grid, c1, c2 Coord) int {
	dX := abs(c2.X - c1.X)
	dY := abs(c2.Y - c1.Y)
	return dX + dY
}

func Part1(input []string) int {
	grid := buildGrid(input, Empty)

	expandedGrid := expandGrid(grid, 2)

	res := 0
	alreadySeen := map[Coord]struct{}{}
	for coord1 := range expandedGrid.Data {
		for coord2 := range alreadySeen {
			length := calculateLength(expandedGrid, coord1, coord2)
			res += length
		}
		alreadySeen[coord1] = struct{}{}
	}

	return res
}

func Part2(input []string, expansionFactor int) int {
	grid := buildGrid(input, Empty)

	expandedGrid := expandGrid(grid, expansionFactor)

	res := 0
	alreadySeen := map[Coord]bool{}
	for coord1 := range expandedGrid.Data {
		for coord2 := range alreadySeen {
			length := calculateLength(expandedGrid, coord1, coord2)
			res += length
		}
		alreadySeen[coord1] = true
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
	fmt.Println("Answer 2 : ", Part2(input, 1000000))
	fmt.Println(time.Since(start2))
}
