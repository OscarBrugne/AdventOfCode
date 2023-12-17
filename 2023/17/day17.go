package main

import (
	"container/heap"
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

func (c1 Coord) Substract(c2 Coord) Coord {
	return Coord{c1.X - c2.X, c1.Y - c2.Y}
}

func (c1 Coord) opposite() Coord {
	return Coord{-c1.X, -c1.Y}
}

type Grid struct {
	Width  int
	Height int
	Data   map[Coord]int
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
		Data:   make(map[Coord]int, len(input)*len(input[0])),
	}

	for y, line := range input {
		for x, char := range line {
			grid.Data[Coord{x, y}] = int(char - '0')
		}
	}

	return grid
}

func (grid Grid) toString() string {
	var result string

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			result += string(rune(grid.Data[coord] + '0'))
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

func (grid Grid) dijkstra(start Coord, goal Coord) (map[Coord]Coord, map[Coord]int) {
	frontier := &utils.CostQueue{}
	heap.Init(frontier)
	heap.Push(frontier, utils.CostQueueItem{Item: start, Cost: 0})

	cameFrom := make(map[Coord]Coord)
	costSoFar := make(map[Coord]int)
	cameFrom[start] = start
	costSoFar[start] = 0

	for frontier.Len() > 0 {
		minItem := heap.Pop(frontier).(utils.CostQueueItem)
		current := minItem.Item.(Coord)
		currentCost := minItem.Cost

		if current == goal {
			break
		}

		for _, next := range grid.neighbors4(current) {
			newCost := currentCost + grid.Data[next]
			if cost, isFound := costSoFar[next]; !isFound || newCost < cost {
				costSoFar[next] = newCost
				priority := newCost
				heap.Push(frontier, utils.CostQueueItem{Item: next, Cost: priority})
				cameFrom[next] = current
			}
		}
	}

	return cameFrom, costSoFar
}

func (grid Grid) dijkstraConstrained(start Coord, goal Coord, minStraight int, maxStraight int) int {
	type Info struct {
		coord       Coord
		dir         Coord
		numStraight int
	}
	startInfo := Info{coord: start, dir: Coord{}, numStraight: 0}

	frontier := &utils.CostQueue{}
	heap.Init(frontier)
	heap.Push(frontier, utils.CostQueueItem{Item: startInfo, Cost: 0})

	costSoFar := make(map[Info]int)
	costSoFar[startInfo] = 0

	for frontier.Len() > 0 {
		minItem := heap.Pop(frontier).(utils.CostQueueItem)
		current := minItem.Item.(Info)
		currentCost := minItem.Cost

		if current.coord == goal {
			return currentCost
		}

		for _, next := range grid.neighbors4(current.coord) {
			newDir := next.Substract(current.coord)
			newNumStraight := 1
			if newDir == current.dir {
				newNumStraight += current.numStraight
			}

			nextInfo := Info{
				coord:       next,
				dir:         newDir,
				numStraight: newNumStraight,
			}
			newCost := currentCost + grid.Data[next]
			actualCost, isFound := costSoFar[nextInfo]

			isValid := (!isFound || newCost < actualCost) &&
				(current.numStraight >= minStraight || newDir == current.dir || current.coord == start) &&
				(newNumStraight <= maxStraight) &&
				(newDir != current.dir.opposite())
			if isValid {
				costSoFar[nextInfo] = newCost

				queueItem := utils.CostQueueItem{Item: nextInfo, Cost: newCost}
				heap.Push(frontier, queueItem)
			}
		}
	}

	return -1
}

func reconstructPath(cameFrom map[Coord]Coord, start Coord, goal Coord) []Coord {
	path := []Coord{goal}

	current := goal
	for current != start {
		fmt.Println(current)
		current = cameFrom[current]
		path = append([]Coord{current}, path...)
	}

	return path
}

func buildPathGrid(grid Grid, path []Coord) Grid {
	pathGrid := Grid{
		Width:  grid.Width,
		Height: grid.Height,
		Data:   grid.Data,
	}

	for i := 0; i < len(path)-1; i++ {
		prevCoord := path[i]
		currCoord := path[i+1]
		dir := currCoord.Substract(prevCoord)
		switch dir {
		case North:
			pathGrid.Data[currCoord] = '^' - '0'
		case East:
			pathGrid.Data[currCoord] = '>' - '0'
		case South:
			pathGrid.Data[currCoord] = 'v' - '0'
		case West:
			pathGrid.Data[currCoord] = '<' - '0'
		}
	}

	return pathGrid
}

func Part1(input []string) int {
	grid := buildGrid(input)

	start := Coord{0, 0}
	goal := Coord{grid.Width - 1, grid.Height - 1}
	res := grid.dijkstraConstrained(start, goal, 0, 3)

	return res
}

func Part2(input []string) int {
	grid := buildGrid(input)

	start := Coord{0, 0}
	goal := Coord{grid.Width - 1, grid.Height - 1}
	res := grid.dijkstraConstrained(start, goal, 4, 10)

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
