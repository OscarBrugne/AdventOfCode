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
	Data   map[Coord]uint8
}

var (
	North = Coord{0, -1}
	West  = Coord{-1, 0}
	South = Coord{0, 1}
	East  = Coord{1, 0}
)

func (coord Coord) isInBounds(grid Grid) bool {
	return 0 <= coord.X && coord.X < grid.Width && 0 <= coord.Y && coord.Y < grid.Height
}

func buildGrid(input []string) Grid {
	grid := Grid{
		Width:  len(input[0]),
		Height: len(input),
		Data:   make(map[Coord]uint8, len(input)*len(input[0])),
	}

	for y, line := range input {
		for x, char := range line {
			grid.Data[Coord{x, y}] = uint8(char - '0')
		}
	}

	return grid
}

func (grid Grid) toString() string {
	var result string

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			result += string(grid.Data[coord] + '0')
		}
		result += "\n"
	}

	return result
}

func (grid Grid) neighbors4(coord Coord) []Coord {
	neighbors := []Coord{}
	directions := []Coord{North, West, South, East}

	for _, dir := range directions {
		neighbor := coord.Add(dir)
		if neighbor.isInBounds(grid) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func (grid Grid) dijkstra(start Coord, end Coord, maxStraight int) int {
	visited := map[Coord]int{}
	frontier := []Coord{start}

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		currentCost := visited[current]

		for _, neighbor := range grid.neighbors4(current) {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = currentCost + int(grid.Data[neighbor])
				frontier = append(frontier, neighbor)
			}

		}
	}

	return visited[end]
}

func Part1(input []string) int {
	grid := buildGrid(input)
	fmt.Println(grid.toString())

	res := grid.dijkstra(Coord{0, 0}, Coord{grid.Width - 1, grid.Height - 1}, 3)

	return res
}

func Part2(input []string) int {
	res := 0
	return res
}

func main() {
	fileName := "input_test.txt"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
