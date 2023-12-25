package main

import (
	"fmt"
	"strings"
	"time"

	"AdventOfCode/utils"
)

type Vertice string

type Edge struct {
	start  Vertice
	end    Vertice
	weight int
}

type Graph map[Vertice]map[Edge]struct{}

func parseInput(input []string) Graph {
	weight := 1

	graph := Graph{}

	for _, line := range input {
		parts := strings.Split(line, ": ")
		vertice := Vertice(parts[0])
		others := strings.Split(parts[1], " ")

		if _, ok := graph[vertice]; !ok {
			graph[vertice] = map[Edge]struct{}{}
		}

		for _, other := range others {
			otherVertice := Vertice(other)
			if _, ok := graph[otherVertice]; !ok {
				graph[otherVertice] = map[Edge]struct{}{}
			}

			graph[vertice][Edge{vertice, otherVertice, weight}] = struct{}{}
			graph[otherVertice][Edge{otherVertice, vertice, weight}] = struct{}{}
		}
	}

	return graph
}

func breadthFirstSearch(graph Graph, start Vertice, end Vertice) (bool, []Vertice) {
	frontier := []Vertice{start}
	reached := map[Vertice]struct{}{start: {}}
	cameFrom := map[Vertice]Vertice{}

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		if current == end {
			path := reconstructPath(start, end, cameFrom)
			return true, path
		}

		for next := range graph[current] {
			if _, ok := reached[next.end]; !ok {
				frontier = append(frontier, next.end)
				reached[next.end] = struct{}{}
				cameFrom[next.end] = current
			}
		}
	}

	return false, []Vertice{}
}

func reconstructPath(start Vertice, end Vertice, cameFrom map[Vertice]Vertice) []Vertice {
	path := []Vertice{}
	current := end
	for current != start {
		path = append([]Vertice{current}, path...)
		current = cameFrom[current]
	}
	path = append([]Vertice{start}, path...)
	return path
}

func copyGraph(graph Graph) Graph {
	newGraph := Graph{}
	for vertice, edges := range graph {
		newGraph[vertice] = map[Edge]struct{}{}
		for edge := range edges {
			newGraph[vertice][edge] = struct{}{}
		}
	}
	return newGraph
}

func Part1(input []string) int {
	minCut := 3

	graph := parseInput(input)

	var start Vertice
	for vertice := range graph {
		start = vertice
		break
	}

	group1 := []Vertice{start}
	group2 := []Vertice{}

	for end := range graph {
		if start == end {
			continue
		}
		newGraph := copyGraph(graph)
		for i := 0; i < minCut; i++ {
			_, path := breadthFirstSearch(newGraph, start, end)
			for j := 0; j < len(path)-1; j++ {
				edge := Edge{path[j], path[j+1], 1}
				delete(newGraph[path[j]], edge)
			}
		}
		isValid, _ := breadthFirstSearch(newGraph, start, end)
		if isValid {
			group1 = append(group1, end)
		} else {
			group2 = append(group2, end)
		}
	}

	return len(group1) * len(group2)
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))
}
