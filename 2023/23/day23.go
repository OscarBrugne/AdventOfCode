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

type Pose struct {
	coord Coord
	dir   Coord
}

const (
	Empty       byte = '.'
	Forest      byte = '#'
	NorthSlopes byte = '^'
	SouthSlopes byte = 'v'
	WestSlopes  byte = '<'
	EastSlopes  byte = '>'
)

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

func getValidDirections(grid Grid, coord Coord, isValidNeighborFunc func(grid Grid, pose Pose) bool) []Coord {
	directions := []Coord{North, South, West, East}
	var validDirections []Coord

	for _, dir := range directions {
		neighbor := Pose{coord.Add(dir), dir}
		if isValidNeighborFunc(grid, neighbor) {
			validDirections = append(validDirections, dir)
		}
	}
	return validDirections
}

func isValidNeighborPart1(grid Grid, pose Pose) bool {
	if !isInBounds(grid, pose.coord) {
		return false
	}
	if _, ok := grid.data[pose.coord]; !ok {
		return true
	}

	if grid.data[pose.coord] == Forest {
		return false
	}
	switch pose.dir {
	case North:
		return grid.data[pose.coord] == NorthSlopes
	case South:
		return grid.data[pose.coord] == SouthSlopes
	case West:
		return grid.data[pose.coord] == WestSlopes
	case East:
		return grid.data[pose.coord] == EastSlopes
	}
	panic("Unreachable.")
}

func isValidNeighborPart2(grid Grid, pose Pose) bool {
	if !isInBounds(grid, pose.coord) {
		return false
	}
	if grid.data[pose.coord] == Forest {
		return false
	}

	return true
}

func findPathsRecursive(grid Grid, start Pose, endCoord Coord, isValidNeighborFunc func(grid Grid, pose Pose) bool) [][]Coord {
	nextCoord := start.coord.Add(start.dir)
	if nextCoord == endCoord {
		return [][]Coord{{endCoord}}
	}

	paths := [][]Coord{}
	dirs := getValidDirections(grid, nextCoord, isValidNeighborFunc)

	for _, dir := range dirs {
		if dir == start.dir.opposite() {
			continue
		}

		neighbor := Pose{nextCoord, dir}

		previousPaths := findPathsRecursive(grid, neighbor, endCoord, isValidNeighborFunc)
		for _, path := range previousPaths {
			path = append([]Coord{start.coord}, path...)
			paths = append(paths, path)
		}
	}

	return paths
}

func Part1(input []string) int {
	grid := parseInput(input)

	start := Pose{Coord{1, 0}, South}
	endCoord := Coord{grid.width - 2, grid.height - 1}

	paths := findPathsRecursive(grid, start, endCoord, isValidNeighborPart1)

	maxi := 0
	for _, path := range paths {
		if len(path) > maxi {
			maxi = len(path)
		}
	}

	return maxi
}

func Part2(input []string) int {
	grid := parseInput(input)

	start := Pose{Coord{1, 0}, South}
	endCoord := Coord{grid.width - 2, grid.height - 1}

	paths := findPathsRecursive(grid, start, endCoord, isValidNeighborPart2)

	maxi := 0
	for _, path := range paths {
		if len(path) > maxi {
			maxi = len(path)
		}
	}

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
