# Depth-First Search (DFS) algorithm in Python

# This code is contributed by Erfan Shafiee# Class to represent a graph
class Graph:
    def __init__(self):
        # Dictionary to store the graph, where keys are nodes and values are lists of adjacent nodes
        self.graph = {}

    def add_edge(self, u, v):
        if u not in self.graph:
            self.graph[u] = []
        self.graph[u].append(v)

    def dfs(self, start):
        # Stack to keep track of the nodes to be visited
        stack = [start]
        
        # List to keep track of visited nodes
        visited = []

        while stack:
            # Pop a node from the stack to visit
            node = stack.pop()

            if node not in visited:
                # Mark the node as visited
                visited.append(node)

                # Get all adjacent nodes of the popped node
                # If an adjacent node hasn't been visited, push it to the stack
                if node in self.graph:
                    for neighbor in reversed(self.graph[node]):
                        if neighbor not in visited:
                            stack.append(neighbor)

        return visited

# Example 
# Create a graph instance
g = Graph()
    
# Add edges to the graph
g.add_edge(0, 1)
g.add_edge(0, 2)
g.add_edge(1, 2)
g.add_edge(2, 0)
g.add_edge(2, 3)
g.add_edge(3, 3)
    
# Perform DFS starting from node 2
result = g.dfs(2)
    
# Print the result of the DFS traversal
print("DFS traversal starting from node 2:", result)
