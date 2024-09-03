package main

import (
	"fmt"
)

const INF = 9999999

func main() {
	// number of vertices in graph
	V := 5
	// create a 2d array of size 5x5 for adjacency matrix to represent graph
	G := [][]int{
		{0, 9, 75, 0, 0},
		{9, 0, 95, 19, 42},
		{75, 95, 0, 51, 66},
		{0, 19, 51, 0, 31},
		{0, 42, 66, 31, 0},
	}
	// create a slice to track selected vertex
	selected := make([]bool, V)
	// set number of edge to 0
	noEdge := 0
	// choose 0th vertex and make it true
	selected[0] = true
	// print for edge and weight
	fmt.Println("Edge : Weight")

	for noEdge < V-1 {
		// For every vertex in the set S, find all adjacent vertices
		// calculate the distance from the vertex selected at step 1.
		// if the vertex is already in the set S, discard it otherwise
		// choose another vertex nearest to selected vertex at step 1.
		minimum := INF
		x, y := 0, 0
		for i := 0; i < V; i++ {
			if selected[i] {
				for j := 0; j < V; j++ {
					if !selected[j] && G[i][j] != 0 {
						// not in selected and there is an edge
						if minimum > G[i][j] {
							minimum = G[i][j]
							x = i
							y = j
						}
					}
				}
			}
		}
		fmt.Printf("%d-%d: %d\n", x, y, G[x][y])
		selected[y] = true
		noEdge++
	}
}

// This code is contributed by Ali Barzegari Dahaj