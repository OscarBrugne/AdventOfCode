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

type Direction int

const (
	Undefined Direction = iota
	Top
	Right
	Bottom
	Left
)

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

func (d Direction) opposite() Direction {
	switch d {
	case Top:
		return Bottom
	case Bottom:
		return Top
	case Left:
		return Right
	case Right:
		return Left
	default:
		return Undefined
	}

}

func (from Coord) isConnected(g Grid, to Coord, dir Direction) bool {
	if !from.isInBounds(g) || !to.isInBounds(g) {
		return false
	}

	connectableFrom := map[Direction][]byte{
		Top:    {'|', 'F', '7', 'S'},
		Right:  {'-', 'J', '7', 'S'},
		Bottom: {'|', 'L', 'J', 'S'},
		Left:   {'-', 'L', 'F', 'S'},
	}

	pipeFrom := g.Data[from]
	pipeTo := g.Data[to]

	for _, validPipe := range connectableFrom[dir] {
		if pipeTo == validPipe {
			for _, validPipe := range connectableFrom[dir.opposite()] {
				if pipeFrom == validPipe {
					return true
				}
			}
		}
	}

	return false
}

func (c Coord) neighborsPipe(g Grid) []Coord {
	possibleNeighbors := map[Direction]Coord{
		Top:    {X: c.X, Y: c.Y - 1},
		Right:  {X: c.X + 1, Y: c.Y},
		Bottom: {X: c.X, Y: c.Y + 1},
		Left:   {X: c.X - 1, Y: c.Y},
	}

	neighbors := []Coord{}
	for dir, neighbor := range possibleNeighbors {
		if c.isConnected(g, neighbor, dir) {
			neighbors = append(neighbors, neighbor)
		}
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

func (start Coord) pathFinding(grid Grid) int {
	alreadySeen := make(map[Coord]bool)
	toExplore := []Coord{start}
	alreadySeen[start] = true

	cntPipesVisited := 0
	for len(toExplore) > 0 {
		size := len(toExplore)
		for i := 0; i < size; i++ {
			current := toExplore[0]
			toExplore = toExplore[1:]

			possiblePipes := current.getPossiblesPipes(grid)
			for _, next := range possiblePipes {
				if !alreadySeen[next] {
					toExplore = append(toExplore, next)
					alreadySeen[next] = true
				}
			}
			cntPipesVisited++
		}
	}

	return cntPipesVisited
}

func Part1(input []string) int {
	grid := buildGrid(input)
	start := findStart(grid)
	numPipesVisited := start.pathFinding(grid)
	maxLength := numPipesVisited - 1
	return maxLength
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
