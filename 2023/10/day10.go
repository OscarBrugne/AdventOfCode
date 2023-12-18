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

type Pipe byte

type Grid struct {
	Width  int
	Height int
	Data   map[Coord]Pipe
}

type Direction int

const (
	Undefined Direction = iota
	Top
	Right
	Bottom
	Left
)

const (
	Empty             Pipe = '.'
	Start             Pipe = 'S'
	Vertical          Pipe = '|'
	Horizontal        Pipe = '-'
	TopLeftCorner     Pipe = 'J'
	TopRightCorner    Pipe = 'L'
	BottomLeftCorner  Pipe = '7'
	BottomRightCorner Pipe = 'F'
	Enclosed          Pipe = 'X'
)

func buildGrid(input []string) Grid {
	grid := Grid{
		Width:  len(input[0]),
		Height: len(input),
		Data:   map[Coord]Pipe{},
	}

	for y, line := range input {
		for x, char := range line {
			if Pipe(char) != Empty {
				grid.Data[Coord{x, y}] = Pipe(char)
			}
		}
	}

	return grid
}

func (grid Grid) toString() string {
	pipesRepres := map[Pipe]string{
		Empty:             " ",
		Start:             "S",
		Vertical:          "║",
		Horizontal:        "═",
		TopLeftCorner:     "╝",
		TopRightCorner:    "╚",
		BottomLeftCorner:  "╗",
		BottomRightCorner: "╔",
		Enclosed:          "X",
	}

	var res string

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			if v, ok := grid.Data[coord]; ok {
				res += pipesRepres[v]
			} else {
				res += pipesRepres[Empty]
			}
		}
		res += "\n"
	}

	return res
}

func (c Coord) isInBounds(g Grid) bool {
	return 0 <= c.X && c.X < g.Width && 0 <= c.Y && c.Y < g.Height
}

func findStart(grid Grid) Coord {
	for coord, value := range grid.Data {
		if value == Start {
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

	connectableFrom := map[Direction][]Pipe{
		Top:    {Vertical, BottomRightCorner, BottomLeftCorner, Start},
		Right:  {Horizontal, TopLeftCorner, BottomLeftCorner, Start},
		Bottom: {Vertical, TopLeftCorner, TopRightCorner, Start},
		Left:   {Horizontal, TopRightCorner, BottomRightCorner, Start},
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

func isolatePath(g Grid, path []Coord, empty Pipe) Grid {
	newGrid := Grid{
		Width:  g.Width,
		Height: g.Height,
		Data:   map[Coord]Pipe{},
	}

	for _, coord := range path {
		newGrid.Data[coord] = g.Data[coord]
	}

	newGrid.Data[path[0]] = getStartPipe(g)

	return newGrid
}

func (c Coord) isInside(g Grid, empty Pipe) bool {
	if _, ok := g.Data[c]; ok {
		return false
	}

	startPipe := empty
	numPipeOnLeft := 0
	for x := 0; x < c.X; x++ {
		coord := Coord{x, c.Y}
		v := g.Data[coord]

		switch v {
		case Vertical:
			numPipeOnLeft++
		case TopRightCorner:
			startPipe = TopRightCorner
		case BottomRightCorner:
			startPipe = BottomRightCorner
		case TopLeftCorner:
			if startPipe == BottomRightCorner {
				startPipe = empty
				numPipeOnLeft++
			} else if v == TopRightCorner {
				startPipe = Empty
			}
		case BottomLeftCorner:
			if startPipe == TopRightCorner {
				startPipe = Empty
				numPipeOnLeft++
			} else if startPipe == BottomRightCorner {
				startPipe = Empty
			}
		}
	}

	return numPipeOnLeft%2 == 1
}

func getStartPipe(g Grid) Pipe {
	start := findStart(g)
	neighbors := start.getPossiblesPipes(g)

	topCoord := Coord{start.X, start.Y - 1}
	bottomCoord := Coord{start.X, start.Y + 1}
	leftCoord := Coord{start.X - 1, start.Y}
	rightCoord := Coord{start.X + 1, start.Y}

	switch {
	case neighbors[0] == topCoord && neighbors[1] == bottomCoord:
		return Vertical
	case neighbors[0] == topCoord && neighbors[1] == rightCoord:
		return TopRightCorner
	case neighbors[0] == topCoord && neighbors[1] == leftCoord:
		return TopLeftCorner
	case neighbors[0] == bottomCoord && neighbors[1] == rightCoord:
		return BottomRightCorner
	case neighbors[0] == bottomCoord && neighbors[1] == leftCoord:
		return BottomLeftCorner
	case neighbors[0] == rightCoord && neighbors[1] == leftCoord:
		return Horizontal
	}

	switch {
	case neighbors[1] == topCoord && neighbors[0] == bottomCoord:
		return Vertical
	case neighbors[1] == topCoord && neighbors[0] == rightCoord:
		return TopRightCorner
	case neighbors[1] == topCoord && neighbors[0] == leftCoord:
		return TopLeftCorner
	case neighbors[1] == bottomCoord && neighbors[0] == rightCoord:
		return BottomRightCorner
	case neighbors[1] == bottomCoord && neighbors[0] == leftCoord:
		return BottomLeftCorner
	case neighbors[1] == rightCoord && neighbors[0] == leftCoord:
		return Horizontal
	}

	panic("Start pipe not found")
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

	isolateGrid := isolatePath(grid, path, Empty)

	cnt := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			c := Coord{X: x, Y: y}
			if c.isInside(isolateGrid, Empty) {
				cnt++
			}
		}
	}

	return cnt
}

func Part2WithPrint(input []string) int {
	grid := buildGrid(input)

	start := findStart(grid)
	path := start.pathFinding(grid)

	isolateGrid := isolatePath(grid, path, Empty)

	cnt := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			c := Coord{X: x, Y: y}
			if c.isInside(isolateGrid, Empty) {
				isolateGrid.Data[c] = Enclosed
				cnt++
			}
		}
	}

	fmt.Println(isolateGrid.toString())
	return cnt
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)

	Part2WithPrint(input)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
