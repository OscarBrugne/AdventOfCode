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

func buildGrid(input []string) Grid {
	width := len(input[0])
	height := len(input)
	data := map[Coord]byte{}
	for j, l := range input {
		for i, c := range l {
			data[Coord{X: i, Y: j}] = byte(c)
		}
	}

	grid := Grid{
		Width:  width,
		Height: height,
		Data:   data,
	}
	return grid
}

func displayGrid(grid Grid, empty byte) {
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if v, ok := grid.Data[Coord{X: y, Y: x}]; ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Printf("%c", empty)
			}
		}
		fmt.Println()
	}
}

func (c Coord) isInBounds(g Grid) bool {
	return 0 <= c.X && c.X < g.Width && 0 <= c.Y && c.Y < g.Height
}

func findStart(grid Grid) Coord {
	for coord, value := range grid.Data {
		if value == 'S' {
			return coord
		}
	}
	return Coord{}
}

func isValidNeighbor(neighbor Coord, g Grid, validPipes []byte) bool {
	if !neighbor.isInBounds(g) {
		return false
	}
	pipe := g.Data[neighbor]
	for _, validPipe := range validPipes {
		if pipe == validPipe {
			return true
		}
	}
	return false
}

func (c Coord) neighborsPipe(g Grid) []Coord {
	isConnectableFrom := map[string][]byte{
		"top":    {'|', 'F', '7'},
		"right":  {'-', 'J', '7'},
		"bottom": {'|', 'L', 'J'},
		"left":   {'-', 'L', 'F'},
	}
	top := Coord{-1, -1}
	right := Coord{-1, -1}
	bottom := Coord{-1, -1}
	left := Coord{-1, -1}

	switch g.Data[c] {
	case '|':
		top = Coord{X: c.X, Y: c.Y - 1}
		bottom = Coord{X: c.X, Y: c.Y + 1}
	case '-':
		left = Coord{X: c.X - 1, Y: c.Y}
		right = Coord{X: c.X + 1, Y: c.Y}
	case 'L':
		top = Coord{X: c.X, Y: c.Y - 1}
		right = Coord{X: c.X + 1, Y: c.Y}
	case 'J':
		top = Coord{X: c.X, Y: c.Y - 1}
		left = Coord{X: c.X - 1, Y: c.Y}
	case '7':
		left = Coord{X: c.X - 1, Y: c.Y}
		bottom = Coord{X: c.X, Y: c.Y + 1}
	case 'F':
		right = Coord{X: c.X + 1, Y: c.Y}
		bottom = Coord{X: c.X, Y: c.Y + 1}
	case 'S':
		top = Coord{X: c.X, Y: c.Y - 1}
		bottom = Coord{X: c.X, Y: c.Y + 1}
		left = Coord{X: c.X - 1, Y: c.Y}
		right = Coord{X: c.X + 1, Y: c.Y}
	default:
		return nil
	}

	neighbors := []Coord{}
	if isValidNeighbor(top, g, isConnectableFrom["top"]) {
		neighbors = append(neighbors, top)
	}
	if isValidNeighbor(right, g, isConnectableFrom["right"]) {
		neighbors = append(neighbors, right)
	}
	if isValidNeighbor(bottom, g, isConnectableFrom["bottom"]) {
		neighbors = append(neighbors, bottom)
	}
	if isValidNeighbor(left, g, isConnectableFrom["left"]) {
		neighbors = append(neighbors, left)
	}
	return neighbors
}

func (c Coord) getPossiblesPipes(g Grid) []Coord {
	possibleDirections := c.neighborsPipe(g)
	validDirections := []Coord{}
	for _, dir := range possibleDirections {
		if dir.isInBounds(g) {
			validDirections = append(validDirections, dir)
		}
	}
	return validDirections
}

func (c Coord) pathFinding(g Grid, length int, alreadySeen map[Coord]bool) int {
	fmt.Println(length)
	newAlreadySeen := make(map[Coord]bool)
	for k, v := range alreadySeen {
		newAlreadySeen[k] = v
	}
	newAlreadySeen[c] = true

	validDirections := c.getPossiblesPipes(g)
	maxLength := length
	for _, dir := range validDirections {
		if !alreadySeen[dir] {
			newLength := dir.pathFinding(g, length+1, newAlreadySeen)
			maxLength = max(maxLength, newLength)
		}
	}
	return maxLength
}

func Part1(input []string) int {
	grid := buildGrid(input)
	start := findStart(grid)
	maxLength := start.pathFinding(grid, 0, map[Coord]bool{})
	lenghtFarthestPoint := (maxLength + 1) / 2
	return lenghtFarthestPoint
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
