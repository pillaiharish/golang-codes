package main

import "fmt"

type Cell struct{ Row, Col int }

var directions = [4]Cell{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// Neighbors returns valid four-direction positions, including on ragged grids.
func Neighbors(grid [][]int, at Cell) []Cell {
	if at.Row < 0 || at.Row >= len(grid) || at.Col < 0 || at.Col >= len(grid[at.Row]) {
		return nil
	}
	result := make([]Cell, 0, 4)
	for _, direction := range directions {
		next := Cell{at.Row + direction.Row, at.Col + direction.Col}
		if next.Row < 0 || next.Row >= len(grid) {
			continue
		}
		if next.Col < 0 || next.Col >= len(grid[next.Row]) {
			continue
		}
		result = append(result, next)
	}
	return result
}

func main() {
	fmt.Println(Neighbors([][]int{{0, 0, 0}, {0, 0}}, Cell{0, 1}))
}
