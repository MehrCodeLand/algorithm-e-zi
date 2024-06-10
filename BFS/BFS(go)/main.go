//BFS algorithm in Go
package main

import (
	"container/list"
	"fmt"
)

// Graph structure
type Graph struct {
	numVertices int
	adjLists    []*list.List
	visited     []bool
}

// Create a new graph with given vertices
func NewGraph(vertices int) *Graph {
	g := &Graph{
		numVertices: vertices,
		adjLists:    make([]*list.List, vertices),
		visited:     make([]bool, vertices),
	}
	for i := range g.adjLists {
		g.adjLists[i] = list.New()
	}
	return g
}

// Add edges to the graph
func (g *Graph) AddEdge(src, dest int) {
	g.adjLists[src].PushBack(dest)
	g.adjLists[dest].PushBack(src)
}

// BFS algorithm
func (g *Graph) BFS(startVertex int) {
	queue := list.New()

	g.visited[startVertex] = true
	queue.PushBack(startVertex)

	for queue.Len() > 0 {
		currVertex := queue.Remove(queue.Front()).(int)
		fmt.Printf("Visited %d ", currVertex)

		for e := g.adjLists[currVertex].Front(); e != nil; e = e.Next() {
			adjVertex := e.Value.(int)
			if !g.visited[adjVertex] {
				g.visited[adjVertex] = true
				queue.PushBack(adjVertex)
			}
		}
	}
}

func main() {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	g.BFS(2)
}

// This code is contributed by Ali Barzegari Dahaj
