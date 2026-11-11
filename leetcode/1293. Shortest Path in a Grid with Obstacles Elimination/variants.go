package main

// shortestPathBruteForce is a learning/reference variant for SMALL grids only.
// It enumerates simple paths and has exponential worst-case running time.
func shortestPathBruteForce(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}

	visited := make([][]bool, m)
	for r := range visited {
		visited[r] = make([]bool, n)
	}
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	// A shortest feasible route can always be made simple (at most m*n-1
	// moves), so m*n+1 is a safe sentinel for "not found".
	best := m*n + 1

	var dfs func(row, col, steps, remaining int)
	dfs = func(row, col, steps, remaining int) {
		if row == m-1 && col == n-1 {
			if steps < best {
				best = steps
			}
			return
		}
		// No completion can beat best if even the unobstructed direct path
		// would take at least best steps.
		if steps+(m-1-row)+(n-1-col) >= best {
			return
		}
		for _, direction := range directions {
			nr, nc := row+direction[0], col+direction[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || visited[nr][nc] {
				continue
			}
			budget := remaining - grid[nr][nc]
			if budget < 0 {
				continue
			}
			visited[nr][nc] = true
			dfs(nr, nc, steps+1, budget)
			visited[nr][nc] = false
		}
	}

	visited[0][0] = true
	dfs(0, 0, 0, k)
	if best == m*n+1 {
		return -1
	}
	return best
}

// shortestPathMemoryOptimized stores the best budget per cell rather than
// a 3D visited array. The tracking table is O(m*n), but the queued states
// still have O(m*n*(k+1)) worst-case total/peak bounds.
func shortestPathMemoryOptimized(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}
	if k >= m+n-3 {
		return m + n - 2
	}

	type state struct {
		row, col, remaining int
	}
	maxRemaining := make([][]int, m)
	for r := range maxRemaining {
		maxRemaining[r] = make([]int, n)
		for c := range maxRemaining[r] {
			maxRemaining[r][c] = -1
		}
	}

	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := []state{{row: 0, col: 0, remaining: k}}
	maxRemaining[0][0] = k
	steps := 0

	for len(queue) > 0 {
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
				if remaining < 0 || remaining <= maxRemaining[nr][nc] {
					continue
				}
				// All prior enqueues have no greater distance (BFS order).
				// More budget at the same cell dominates fewer eliminations.
				maxRemaining[nr][nc] = remaining
				nextQueue = append(nextQueue, state{nr, nc, remaining})
			}
		}
		queue = nextQueue
		steps++
	}
	return -1
}
