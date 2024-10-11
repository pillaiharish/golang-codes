package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	} // empty grid no islands
	count := 0
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	rows, cols := len(grid), len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				bfs(grid, r, c, rows, cols, directions)
			}
		}
	}
	return count
}

func bfs(grid [][]byte, r, c, rows, cols int, directions [][]int) {
	queue := [][]int{{r, c}}
	grid[r][c] = '0' //visited

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newRow, newCol := cell[0]+dir[0], cell[1]+dir[1]
			// make sure to check the equality of the value as Zero is acceptable new Row/Col index
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == '1' {
				queue = append(queue, []int{newRow, newCol})
				grid[newRow][newCol] = '0' //Visited
			}
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println("Number of islands:", numIslands(grid))
}

// Output:
// Number of islands: 1
