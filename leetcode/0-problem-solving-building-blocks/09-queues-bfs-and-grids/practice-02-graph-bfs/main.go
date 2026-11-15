package main

import (
	"fmt"
)

// BFS performs a breadth-first search on a graph starting from a specific node
func BFS(graph map[int][]int, start int) {
	// 1. Map to keep track of visited nodes (acts like a Set)
	visited := make(map[int]bool)

	// 2. Slice to act as our queue
	queue := []int{start}

	// Mark the starting node as visited
	visited[start] = true

	fmt.Print("BFS Traversal: ")

	// 3. Loop until the queue is empty
	for len(queue) > 0 {
		// Dequeue the front element
		curr := queue[0]
		queue = queue[1:]

		// Process the current node
		fmt.Printf("%d ", curr)

		// 4. Enqueue all unvisited neighbors
		for _, neighbor := range graph[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true // Mark visited as soon as it's queued
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}

func main() {
	// Representing a graph using an Adjacency List
	// 0 --- 1 --- 2
	// |     |
	// 3 --- 4
	graph := map[int][]int{
		0: {1, 3},
		1: {0, 2, 4},
		2: {1},
		3: {0, 4},
		4: {1, 3},
	}

	// Start BFS from node 0
	BFS(graph, 0)
}
