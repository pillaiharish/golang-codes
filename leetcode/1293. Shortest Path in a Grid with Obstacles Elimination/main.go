package main

import "fmt"

// shortestPath finds the fewest four-direction moves from the top-left to the
// bottom-right cell while eliminating at most k obstacles.
// The grid is rectangular, contains only 0/1, and has empty endpoints.
// Time: O(m*n*(k+1)); extra space: O(m*n*(k+1)).
func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}

	// A Manhattan-length path has at most m+n-3 intermediate obstacles.
	// With this budget, a shortest route is always possible.
	if k >= m+n-3 {
		return m + n - 2
	}

	// Being at the same cell with different budgets is a different BFS state.
	type state struct {
		row, col, remaining int
	}
	visited := make([][][]bool, m)
	for r := range visited {
		visited[r] = make([][]bool, n)
		for c := range visited[r] {
			visited[r][c] = make([]bool, k+1)
		}
	}

	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := []state{{row: 0, col: 0, remaining: k}}
	visited[0][0][k] = true
	steps := 0

	for len(queue) > 0 {
		// Every item in queue is exactly 'steps' moves from the start.
		nextQueue := make([]state, 0, len(queue))
		for _, current := range queue {
			if current.row == m-1 && current.col == n-1 {
				return steps
			}

			for _, direction := range directions {
				nr, nc := current.row+direction[0], current.col+direction[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					continue
				}

				remaining := current.remaining - grid[nr][nc]
				if remaining < 0 || visited[nr][nc][remaining] {
					continue
				}

				// Mark at enqueue time: no duplicate visit of this exact state.
				visited[nr][nc][remaining] = true
				nextQueue = append(nextQueue, state{nr, nc, remaining})
			}
		}
		queue = nextQueue
		steps++
	}
	return -1
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}
	fmt.Println(shortestPath(grid, 1)) // 6
}
