package main

import (
	"fmt"
	"sort"
)

type Graph struct {
	V     int
	graph [][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{V: vertices, graph: [][]int{}}
}

func (g *Graph) addEdge(u, v, w int) {
	g.graph = append(g.graph, []int{u, v, w})
}

func (g *Graph) find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return g.find(parent, parent[i])
}

func (g *Graph) applyUnion(parent, rank []int, x, y int) {
	xroot := g.find(parent, x)
	yroot := g.find(parent, y)
	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}
}

func (g *Graph) kruskalAlgo() {
	result := [][]int{}
	i, e := 0, 0
	sort.Slice(g.graph, func(a, b int) bool {
		return g.graph[a][2] < g.graph[b][2]
	})
	parent := make([]int, g.V)
	rank := make([]int, g.V)
	for node := 0; node < g.V; node++ {
		parent[node] = node
		rank[node] = 0
	}
	for e < g.V-1 {
		u, v, w := g.graph[i][0], g.graph[i][1], g.graph[i][2]
		i++
		x := g.find(parent, u)
		y := g.find(parent, v)
		if x != y {
			e++
			result = append(result, []int{u, v, w})
			g.applyUnion(parent, rank, x, y)
		}
	}
	for _, edge := range result {
		fmt.Printf("%d - %d: %d\n", edge[0], edge[1], edge[2])
	}
}

func main() {
	g := NewGraph(6)
	g.addEdge(0, 1, 4)
	g.addEdge(0, 2, 4)
	g.addEdge(1, 2, 2)
	g.addEdge(1, 0, 4)
	g.addEdge(2, 0, 4)
	g.addEdge(2, 1, 2)
	g.addEdge(2, 3, 3)
	g.addEdge(2, 5, 2)
	g.addEdge(2, 4, 4)
	g.addEdge(3, 2, 3)
	g.addEdge(3, 4, 3)
	g.addEdge(4, 2, 4)
	g.addEdge(4, 3, 3)
	g.addEdge(5, 2, 2)
	g.addEdge(5, 4, 3)
	g.kruskalAlgo()
}

// This code is contributed by Ali Barzegari Dahaj