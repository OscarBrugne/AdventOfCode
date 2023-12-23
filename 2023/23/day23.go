package main

import (
	"fmt"
	"time"

	"AdventOfCode/utils"
)

type Coord struct {
	x int
	y int
}

func (c1 Coord) Add(c2 Coord) Coord {
	return Coord{c1.x + c2.x, c1.y + c2.y}
}

func (c Coord) opposite() Coord {
	return Coord{-c.x, -c.y}
}

type Grid struct {
	width  int
	height int
	data   map[Coord]byte
}

var (
	North = Coord{0, -1}
	South = Coord{0, 1}
	West  = Coord{-1, 0}
	East  = Coord{1, 0}
)

const (
	Empty       byte = '.'
	Forest      byte = '#'
	NorthSlopes byte = '^'
	SouthSlopes byte = 'v'
	WestSlopes  byte = '<'
	EastSlopes  byte = '>'
)

var SlopeToDir = map[byte]Coord{
	NorthSlopes: North,
	SouthSlopes: South,
	WestSlopes:  West,
	EastSlopes:  East,
}

func (grid Grid) toString() string {
	res := ""
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if v, ok := grid.data[Coord{x, y}]; ok {
				res += string(v)
			} else {
				res += string(Empty)
			}
		}
		res += "\n"
	}
	return res
}

func isInBounds(grid Grid, coord Coord) bool {
	return 0 <= coord.x && coord.x < grid.width && 0 <= coord.y && coord.y < grid.height
}

func parseInput(input []string) Grid {
	grid := Grid{
		width:  len(input[0]),
		height: len(input),
		data:   make(map[Coord]byte),
	}

	for y, line := range input {
		for x, char := range line {
			if byte(char) != Empty {
				grid.data[Coord{x, y}] = byte(char)
			}
		}
	}

	return grid
}

func isValidNeighbor(grid Grid, coord, dir Coord) bool {
	if !isInBounds(grid, coord) {
		return false
	}
	if _, ok := grid.data[coord]; !ok {
		return true
	}
	if grid.data[coord] == Forest {
		return false
	}
	return SlopeToDir[grid.data[coord]] == dir
}

func neighbors4(grid Grid, coord Coord) []Coord {
	directions := []Coord{North, South, West, East}
	validNeighbors := []Coord{}

	for _, dir := range directions {
		neighbor := coord.Add(dir)
		if isValidNeighbor(grid, neighbor, dir) {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}

	return validNeighbors
}

func depthFirstSearch(grid Grid, current Coord, end Coord, seen map[Coord]struct{}) int {
	if current == end {
		return 0
	}

	maxi := -grid.width * grid.height

	seen[current] = struct{}{}
	for _, neighbor := range neighbors4(grid, current) {
		if _, ok := seen[neighbor]; !ok {
			maxi = max(maxi, depthFirstSearch(grid, neighbor, end, seen)+1)
		}
	}
	delete(seen, current)

	return maxi
}

func Part1(input []string) int {
	grid := parseInput(input)

	start := Coord{1, 0}
	end := Coord{grid.width - 2, grid.height - 1}

	maxi := depthFirstSearch(grid, start, end, map[Coord]struct{}{})
	return maxi
}

func Part2(input []string) int {
	grid := parseInput(input)
	for coord, value := range grid.data {
		if _, ok := SlopeToDir[value]; ok {
			delete(grid.data, coord)
		}
	}

	start := Coord{1, 0}
	end := Coord{grid.width - 2, grid.height - 1}

	maxi := depthFirstSearch(grid, start, end, map[Coord]struct{}{})
	return maxi
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
