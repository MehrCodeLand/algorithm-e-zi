// BFS algorithm in C#
using System;
using System.Collections.Generic;

class Graph
{
    private int numVertices;
    private List<int>[] adjLists;
    private bool[] visited;

    public Graph(int vertices)
    {
        numVertices = vertices;
        adjLists = new List<int>[vertices];
        for (int i = 0; i < vertices; i++)
            adjLists[i] = new List<int>();
    }

    public void AddEdge(int src, int dest)
    {
        adjLists[src].Add(dest);
        adjLists[dest].Add(src);
    }

    public void BFS(int startVertex)
    {
        visited = new bool[numVertices];
        Queue<int> queue = new Queue<int>();

        visited[startVertex] = true;
        queue.Enqueue(startVertex);

        while (queue.Count != 0)
        {
            int currVertex = queue.Dequeue();
            Console.WriteLine("Visited " + currVertex);

            foreach (int adjVertex in adjLists[currVertex])
            {
                if (!visited[adjVertex])
                {
                    visited[adjVertex] = true;
                    queue.Enqueue(adjVertex);
                }
            }
        }
    }
}

class Program
{
    static void Main(string[] args)
    {
        Graph g = new Graph(4);
        g.AddEdge(0, 1);
        g.AddEdge(0, 2);
        g.AddEdge(1, 2);
        g.AddEdge(2, 0);
        g.AddEdge(2, 3);
        g.AddEdge(3, 3);

        g.BFS(2);
    }
}

// This code is contributed by Ali Barzegari Dahaj
