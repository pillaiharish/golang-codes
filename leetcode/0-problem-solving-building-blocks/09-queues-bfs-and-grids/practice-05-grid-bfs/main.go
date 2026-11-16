package main

import "fmt"

type Cell struct{ Row, Col int }

var moves = [4]Cell{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// ShortestGridPath returns the fewest moves over zero-valued cells, or -1.
// Each move costs 1. Rows may have different lengths.
func ShortestGridPath(grid [][]int, start, goal Cell) int {
	valid := func(c Cell) bool {
		return c.Row >= 0 && c.Row < len(grid) && c.Col >= 0 &&
			c.Col < len(grid[c.Row]) && grid[c.Row][c.Col] == 0
	}
	if !valid(start) || !valid(goal) {
		return -1
	}
	if start == goal {
		return 0
	}
	visited := make([][]bool, len(grid))
	for r := range grid {
		visited[r] = make([]bool, len(grid[r]))
	}
	visited[start.Row][start.Col] = true
	queue := []Cell{start}
	head, steps := 0, 0

	for head < len(queue) {
		levelEnd := len(queue) // snapshot: appended nodes are next-level states
		for head < levelEnd {
			current := queue[head]
			head++
			for _, move := range moves {
				next := Cell{current.Row + move.Row, current.Col + move.Col}
				if !valid(next) || visited[next.Row][next.Col] {
					continue
				}
				if next == goal {
					return steps + 1
				}
				visited[next.Row][next.Col] = true // visit when enqueuing
				queue = append(queue, next)
			}
		}
		steps++
	}
	return -1
}

func main() {
	grid := [][]int{{0, 1, 0}, {0, 0, 0}, {1, 1, 0}}
	fmt.Println(ShortestGridPath(grid, Cell{0, 0}, Cell{2, 2})) // 4
}
