package main

import (
	"fmt"
	"math"
)

var vertices = [][]int{
	{0, 0, 1, 1, 0, 0, 0},
	{0, 0, 1, 0, 0, 1, 0},
	{1, 1, 0, 1, 1, 0, 0},
	{1, 0, 1, 0, 0, 0, 1},
	{0, 0, 1, 0, 0, 1, 0},
	{0, 1, 0, 0, 1, 0, 1},
	{0, 0, 0, 1, 0, 1, 0},
}

var edges = [][]int{
	{0, 0, 1, 2, 0, 0, 0},
	{0, 0, 2, 0, 0, 3, 0},
	{1, 2, 0, 1, 3, 0, 0},
	{2, 0, 1, 0, 0, 0, 1},
	{0, 0, 3, 0, 0, 2, 0},
	{0, 3, 0, 0, 2, 0, 1},
	{0, 0, 0, 1, 0, 1, 0},
}

var visitedAndDistance [][]int

func toBeVisited(numOfVertices int) int {
	v := -10
	for index := 0; index < numOfVertices; index++ {
		if visitedAndDistance[index][0] == 0 && (v < 0 || visitedAndDistance[index][1] <= visitedAndDistance[v][1]) {
			v = index
		}
	}
	return v
}

func main() {
	numOfVertices := len(vertices[0])
	visitedAndDistance = make([][]int, numOfVertices)
	visitedAndDistance[0] = []int{0, 0}
	for i := 1; i < numOfVertices; i++ {
		visitedAndDistance[i] = []int{0, math.MaxInt32}
	}

	for vertex := 0; vertex < numOfVertices; vertex++ {
		toVisit := toBeVisited(numOfVertices)
		for neighborIndex := 0; neighborIndex < numOfVertices; neighborIndex++ {
			if vertices[toVisit][neighborIndex] == 1 && visitedAndDistance[neighborIndex][0] == 0 {
				newDistance := visitedAndDistance[toVisit][1] + edges[toVisit][neighborIndex]
				if visitedAndDistance[neighborIndex][1] > newDistance {
					visitedAndDistance[neighborIndex][1] = newDistance
				}
			}
		}
		visitedAndDistance[toVisit][0] = 1
	}

	for i, distance := range visitedAndDistance {
		fmt.Printf("Distance of %c from source vertex: %d\n", 'a'+i, distance[1])
	}
}

// This code is contributed by Ali Barzegari Dahaj