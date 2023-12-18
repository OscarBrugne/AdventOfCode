package main

import (
	"fmt"
	"strconv"
	"strings"
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

var (
	North = Coord{0, -1}
	West  = Coord{-1, 0}
	South = Coord{0, 1}
	East  = Coord{1, 0}
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInputPart1(input []string) []Coord {
	const (
		Up    = 'U'
		Left  = 'L'
		Down  = 'D'
		Right = 'R'
	)

	current := Coord{0, 0}
	vertices := []Coord{current}

	for _, line := range input {
		parts := strings.Split(line, " ")
		dirInput := parts[0][0]
		lengthStr := parts[1]
		length := 0
		for i := 0; i < len(lengthStr); i++ {
			length = length*10 + int(lengthStr[i]-'0')
		}

		var dir Coord
		switch dirInput {
		case Up:
			dir = North
		case Left:
			dir = West
		case Down:
			dir = South
		case Right:
			dir = East
		}

		current = current.Add(dir.MultiplyByScalar(length))
		vertices = append(vertices, current)
	}

	return vertices
}

func parseInputPart2(input []string) []Coord {
	const (
		Up    = '3'
		Left  = '2'
		Down  = '1'
		Right = '0'
	)

	current := Coord{0, 0}
	vertices := []Coord{current}

	for _, line := range input {
		parts := strings.Split(line, " ")
		color := parts[2]
		dirInput := color[7]
		lengthStr := color[2:7]
		length := hexStringToInt(lengthStr)

		var dir Coord
		switch dirInput {
		case Up:
			dir = North
		case Left:
			dir = West
		case Down:
			dir = South
		case Right:
			dir = East
		}

		current = current.Add(dir.MultiplyByScalar(length))
		vertices = append(vertices, current)
	}

	return vertices
}

func hexStringToInt(hexStr string) int {
	num, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		panic(err)
	}

	return int(num)
}

func shoelace(vertices []Coord) int {
	n := len(vertices)
	area := 0

	for i := 0; i < n; i++ {
		next := (i + 1) % n
		area += vertices[i].X * vertices[next].Y
		area -= vertices[i].Y * vertices[next].X
	}

	area = Abs(area) / 2
	return area
}

func perimeter(vertices []Coord) int {
	n := len(vertices)
	perim := 0

	for i := 0; i < n; i++ {
		next := (i + 1) % n
		perim += Abs(vertices[i].X-vertices[next].X) + Abs(vertices[i].Y-vertices[next].Y)
	}

	return perim
}

func calculatePolygonArea(vertices []Coord) int {
	// Using Pick's theorem
	return shoelace(vertices) + perimeter(vertices)/2 + 1
}

func Part1(input []string) int {
	vertices := parseInputPart1(input)

	res := calculatePolygonArea(vertices)
	return res
}

func Part2(input []string) int {
	vertices := parseInputPart2(input)

	res := calculatePolygonArea(vertices)
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
