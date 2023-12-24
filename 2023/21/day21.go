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

func (c Coord) MultiplyByScalar(s int) Coord {
	return Coord{c.X * s, c.Y * s}
}

type Grid struct {
	Width  int
	Height int
	Data   map[Coord]byte
}

var (
	North = Coord{0, -1}
	West  = Coord{-1, 0}
	South = Coord{0, 1}
	East  = Coord{1, 0}
)

const (
	Empty byte = '.'
	Rock  byte = '#'
	Start byte = 'S'
)

func (grid Grid) toString() string {
	res := ""
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if v, ok := grid.Data[Coord{x, y}]; ok {
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
	return 0 <= coord.X && coord.X < grid.Width && 0 <= coord.Y && coord.Y < grid.Height
}

func parseInput(input []string) Grid {
	grid := Grid{
		Width:  len(input[0]),
		Height: len(input),
		Data:   make(map[Coord]byte),
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

func findStart(grid Grid) Coord {
	for coord, char := range grid.Data {
		if char == Start {
			return coord
		}
	}
	panic("No start found.")
}

func neighbors4(grid Grid, coord Coord) []Coord {
	neighbors := []Coord{
		coord.Add(North),
		coord.Add(South),
		coord.Add(East),
		coord.Add(West),
	}

	validNeighbors := []Coord{}

	for _, neighbor := range neighbors {
		if isInBounds(grid, neighbor) && grid.Data[neighbor] != Rock {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}

	return validNeighbors
}

func breadthFirstSearch(grid Grid, start Coord, neighborFunc func(Grid, Coord) []Coord) map[Coord]int {
	frontier := []Coord{start}
	reached := map[Coord]struct{}{start: {}}
	cameFrom := map[Coord]Coord{start: start}
	distances := map[Coord]int{start: 0}

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		for _, next := range neighborFunc(grid, current) {
			if _, ok := reached[next]; !ok {
				frontier = append(frontier, next)
				reached[next] = struct{}{}
				cameFrom[next] = current
				distances[next] = distances[current] + 1
			}
		}
	}

	return distances
}

func distancesFromExtremities(grid Grid) map[Coord]map[Coord]int {
	distances := make(map[Coord]map[Coord]int, 8)

	extremities := []Coord{
		{0, 0},
		{grid.Width / 2, 0},
		{grid.Width, 0},
		{grid.Width, grid.Height / 2},
		{grid.Width, grid.Height},
		{grid.Width / 2, grid.Height},
		{0, grid.Height},
		{0, grid.Height / 2},
	}

	for _, start := range extremities {
		distances[start] = breadthFirstSearch(grid, start, neighbors4)
	}

	return distances
}

func neighbors8(grid Grid, coord Coord) []Coord {
	neighbors := []Coord{
		coord.Add(North),
		coord.Add(South),
		coord.Add(East),
		coord.Add(West),
		coord.Add(North).Add(East),
		coord.Add(North).Add(West),
		coord.Add(South).Add(East),
		coord.Add(South).Add(West),
	}

	return neighbors
}

func Part1(input []string, numSteps int) int {
	grid := parseInput(input)

	start := findStart(grid)
	distances := breadthFirstSearch(grid, start, neighbors4)

	cnt := 0
	for _, dist := range distances {
		if dist <= numSteps && dist%2 == 0 {
			cnt++
		}
	}
	return cnt
}

func Part2(input []string, numSteps int) int {
	grid := parseInput(input)

	start := findStart(grid)
	distancesFromStart := breadthFirstSearch(grid, start, neighbors4)
	distancesFromExtremities := distancesFromExtremities(grid)

	_, _ = distancesFromStart, distancesFromExtremities
	// La map étant infinie, on change de côté quand on arrive à un bord
	// Ex: Si on est en (0, 0) et qu'on va au nord, on arrive en (0, grid.Height) qui est en bas de la map

	// La map a 2 particularités :
	// 1) Il n'y a pas d'obstacle sur la même ligne sur la même colonne que le départ.
	// 2) Il n'y a pas d'obstacle sur le bord.

	// Pour calculer le nombre de case situés à numSteps du départ sans reparcourir toutes les cartes :
	// 1) On parcourt l'ensemble de la carte depuis S et on compte le nombre de cases situés à un nombre de pas pair/impair.
	// 2) On va sur les 4 bords en ligne droite depuis S, respectivement edgeN, edgeS, edgeE, edgeW.
	// 3) On change de côté, on arrive respectivement à startS, startN, startW, startE.
	// 4) On va en ligne droite perpendiculairement pour atteindre les cartes en diagonales.
	// 5) On compte le nombre de cases situés à un nombre de pas pair/impair pour chacune de ses cartes.
	// 6) On répète ce processur avec un algorithme de BFS où chaque case correspond à une carte.

	return 0
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input, 64))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input, 26501365))
	fmt.Println(time.Since(start2))
}
