package main

import (
	"fmt"
	"time"

	"AdventOfCode/utils"
)

type Coord struct {
	x float64
	y float64
	z float64
}

type Point struct {
	pos Coord
	vel Coord
}

func parseInput(input []string) []Point {
	points := make([]Point, len(input))
	for i, line := range input {
		point := Point{}
		fmt.Sscanf(
			line, "%f, %f, %f @ %f,  %f, %f",
			&point.pos.x, &point.pos.y, &point.pos.z,
			&point.vel.x, &point.vel.y, &point.vel.z,
		)
		points[i] = point
	}
	return points
}

func isIntersecting2D(p1, p2 Point) (bool, Coord, float64, float64) {
	det := p1.vel.x*p2.vel.y - p2.vel.x*p1.vel.y
	if det == 0 {
		return false, Coord{}, 0, 0
	}
	t1 := (p2.vel.y*(p2.pos.x-p1.pos.x) - p2.vel.x*(p2.pos.y-p1.pos.y)) / det
	t2 := (p1.vel.y*(p2.pos.x-p1.pos.x) - p1.vel.x*(p2.pos.y-p1.pos.y)) / det
	coord := Coord{
		x: p1.pos.x + p1.vel.x*t1,
		y: p1.pos.y + p1.vel.y*t1,
		z: 0,
	}
	return true, coord, t1, t2

}

func Part1(input []string, min float64, max float64) int {
	points := parseInput(input)

	cnt := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < i; j++ {
			isIntersecting, coord, time1, time2 := isIntersecting2D(points[i], points[j])
			isInBound := min <= coord.x && coord.x <= max && min <= coord.y && coord.y <= max
			if isIntersecting && isInBound && time1 >= 0 && time2 >= 0 {
				cnt++
			}
		}
	}
	return cnt
}

func Part2(input []string) int {
	res := 0
	return res
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input, 200000000000000, 400000000000000))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
