// DfS in go
package main

import "fmt"

// Graph represents a graph using an adjacency list.
type Graph struct {
	vertices map[int][]int
}

// NewGraph creates a new Graph.
func NewGraph() *Graph {
	return &Graph{vertices: make(map[int][]int)}
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
}

// DFS performs Depth-First Search starting from a given vertex.
func (g *Graph) DFS(start int) {
	// Create a map to keep track of visited vertices
	visited := make(map[int]bool)
	
	// Call the recursive helper function to perform DFS
	g.dfsUtil(start, visited)
}

// dfsUtil is a recursive function used by DFS to visit vertices.
func (g *Graph) dfsUtil(v int, visited map[int]bool) {
	// Mark the current vertex as visited
	visited[v] = true
	fmt.Printf("%d ", v)

	// Recur for all the vertices adjacent to this vertex
	for _, i := range g.vertices[v] {
		if !visited[i] {
			g.dfsUtil(i, visited)
		}
	}
}

func main() {
	// Create a graph
	g := NewGraph()
	
	// Add edges to the graph
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	
	// Print the DFS traversal starting from vertex 2
	fmt.Println("Depth First Traversal starting from vertex 2:")
	g.DFS(2)
}

// This code is contributed by Erfan Shafiee