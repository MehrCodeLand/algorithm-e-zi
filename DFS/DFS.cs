// DFS algorithm in C#
using System;
using System.Collections.Generic;

class Graph
{
    private int numVertices;
    private List<int>[] adjLists;
    private bool[] visited;

    public Graph(int V)
    {
        numVertices = V;
        adjLists = new List<int>[V];
        for (int i = 0; i < V; ++i)
            adjLists[i] = new List<int>();
        visited = new bool[V];
    }

    public void AddEdge(int src, int dest)
    {
        adjLists[src].Insert(0, dest);
    }

    public void DFS(int vertex)
    {
        visited[vertex] = true;
        Console.Write(vertex + " ");

        foreach (int i in adjLists[vertex])
        {
            if (!visited[i])
                DFS(i);
        }
    }

    static void Main(string[] args)
    {
        Graph g = new Graph(4);
        g.AddEdge(0, 1);
        g.AddEdge(0, 2);
        g.AddEdge(1, 2);
        g.AddEdge(2, 3);

        g.DFS(2);
    }
}

// This code is contributed by Ali Barzegari Dahaj