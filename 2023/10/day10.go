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
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			if v, ok := grid.Data[coord]; ok {
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

func (start Coord) pathFinding(grid Grid) []Coord {
	parents := map[Coord]Coord{}
	alreadySeen := map[Coord]bool{}
	path := []Coord{}

	alreadySeen[start] = true
	path = append(path, start)

	startNeighbors := start.getPossiblesPipes(grid)
	startNeighbor := startNeighbors[0]
	alreadySeen[startNeighbor] = true
	path = append(path, startNeighbor)

	toExplore := []Coord{startNeighbor}
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
					parents[next] = current
					path = append(path, next)
				}
			}
		}
	}

	return path
}

func isolatePath(g Grid, path []Coord, empty byte) Grid {
	newData := map[Coord]byte{}
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			newData[Coord{x, y}] = empty
		}
	}

	for _, c := range path {
		newData[c] = g.Data[c]
	}

	newData[path[0]] = getStartPipe(g)

	newGrid := Grid{
		Width:  g.Width,
		Height: g.Height,
		Data:   newData,
	}
	return newGrid
}

func (c Coord) isInside(g Grid, empty byte) bool {
	if g.Data[c] != empty {
		return false
	}

	startPipe := empty
	numPipeOnLeft := 0
	for x := 0; x < c.X; x++ {
		coord := Coord{x, c.Y}
		v := g.Data[coord]

		switch v {
		case '|':
			numPipeOnLeft++
		case 'L':
			startPipe = 'L'
		case 'F':
			startPipe = 'F'
		case 'J':
			if startPipe == 'F' {
				startPipe = empty
				numPipeOnLeft++
			} else if v == 'L' {
				startPipe = '.'
			}
		case '7':
			if startPipe == 'L' {
				startPipe = '.'
				numPipeOnLeft++
			} else if startPipe == 'F' {
				startPipe = '.'
			}
		}
	}

	return numPipeOnLeft%2 == 1
}

func getStartPipe(g Grid) byte {
	start := findStart(g)
	neighbors := start.getPossiblesPipes(g)

	topCoord := Coord{start.X - 1, start.Y}
	bottomCoord := Coord{start.X + 1, start.Y}
	leftCoord := Coord{start.X, start.Y - 1}
	rightCoord := Coord{start.X, start.Y + 1}

	switch {
	case neighbors[0] == topCoord && neighbors[1] == bottomCoord:
		return '|'
	case neighbors[0] == topCoord && neighbors[1] == rightCoord:
		return 'L'
	case neighbors[0] == topCoord && neighbors[1] == leftCoord:
		return 'J'
	case neighbors[0] == bottomCoord && neighbors[1] == rightCoord:
		return 'F'
	case neighbors[0] == bottomCoord && neighbors[1] == leftCoord:
		return '7'
	case neighbors[0] == rightCoord && neighbors[1] == leftCoord:
		return '-'
	}

	switch {
	case neighbors[1] == topCoord && neighbors[0] == bottomCoord:
		return '|'
	case neighbors[1] == topCoord && neighbors[0] == rightCoord:
		return 'L'
	case neighbors[1] == topCoord && neighbors[0] == leftCoord:
		return 'J'
	case neighbors[1] == bottomCoord && neighbors[0] == rightCoord:
		return 'F'
	case neighbors[1] == bottomCoord && neighbors[0] == leftCoord:
		return '7'
	case neighbors[1] == rightCoord && neighbors[0] == leftCoord:
		return '-'
	}

	panic("Start pipe not find")
}

func Part1(input []string) int {
	grid := buildGrid(input)
	start := findStart(grid)
	path := start.pathFinding(grid)
	numPipesVisited := len(path)
	maxLength := numPipesVisited / 2
	return maxLength
}

func Part2(input []string) int {
	grid := buildGrid(input)
	start := findStart(grid)
	path := start.pathFinding(grid)

	isolateGrid := isolatePath(grid, path, '.')

	cnt := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			c := Coord{X: x, Y: y}
			if c.isInside(isolateGrid, '.') {
				isolateGrid.Data[c] = 'I' // optional
				cnt++
			}
		}
	}

	displayGrid(isolateGrid, ' ') // optional

	return cnt
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
