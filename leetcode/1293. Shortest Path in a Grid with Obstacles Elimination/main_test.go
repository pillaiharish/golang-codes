package main

import (
	"math/rand"
	"reflect"
	"slices"
	"testing"
)

func TestShortestPathAndMemoryOptimized(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{
			name: "LeetCode example 1", k: 1, want: 6,
			grid: [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}},
		},
		{
			name: "LeetCode example 2 unreachable", k: 1, want: -1,
			grid: [][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}},
		},
		{name: "single cell", grid: [][]int{{0}}, k: 0, want: 0},
		{name: "one row no obstacle", grid: [][]int{{0, 0, 0, 0}}, k: 0, want: 3},
		{name: "one row obstacle removed", grid: [][]int{{0, 1, 0}}, k: 1, want: 2},
		{name: "one row obstacle blocked", grid: [][]int{{0, 1, 0}}, k: 0, want: -1},
		{name: "one column obstacle removed", grid: [][]int{{0}, {1}, {0}}, k: 1, want: 2},
		{name: "two distinct states at same cell", grid: [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 0}}, k: 1, want: 4},
		{name: "shortcut exactly enough eliminations", grid: [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}, k: 3, want: 4},
		{name: "shortcut one elimination short", grid: [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}, k: 2, want: -1},
		{name: "blocked two-by-two", grid: [][]int{{0, 1}, {1, 0}}, k: 0, want: -1},
		{name: "two-by-two elimination", grid: [][]int{{0, 1}, {1, 0}}, k: 1, want: 2},
	}

	variants := []struct {
		name string
		fn   func([][]int, int) int
	}{
		{name: "exact-state BFS", fn: shortestPath},
		{name: "dominance BFS", fn: shortestPathMemoryOptimized},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, variant := range variants {
				t.Run(variant.name, func(t *testing.T) {
					before := cloneGrid(tt.grid)
					if got := variant.fn(tt.grid, tt.k); got != tt.want {
						t.Fatalf("%s: got %d; want %d", variant.name, got, tt.want)
					}
					if !reflect.DeepEqual(tt.grid, before) {
						t.Fatal("input grid was mutated")
					}
				})
			}
		})
	}
}

func TestShortestPathVariantsAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(1293))
	for trial := 0; trial < 300; trial++ {
		m, n := 1+rng.Intn(4), 1+rng.Intn(4)
		grid := make([][]int, m)
		for r := range grid {
			grid[r] = make([]int, n)
			for c := range grid[r] {
				grid[r][c] = rng.Intn(2)
			}
		}
		grid[0][0] = 0
		grid[m-1][n-1] = 0
		k := rng.Intn(4)

		want := shortestPathBruteForce(grid, k)
		for _, variant := range []struct {
			name string
			fn   func([][]int, int) int
		}{
			{"exact-state BFS", shortestPath},
			{"dominance BFS", shortestPathMemoryOptimized},
		} {
			if got := variant.fn(grid, k); got != want {
				t.Fatalf("trial %d, %s: grid=%v k=%d, got %d, brute force %d", trial, variant.name, grid, k, got, want)
			}
		}
	}
}

func TestShortestPathLargeOpenGrid(t *testing.T) {
	grid := make([][]int, 40)
	for r := range grid {
		grid[r] = make([]int, 40)
	}
	if got := shortestPath(grid, 0); got != 78 {
		t.Fatalf("shortestPath(40x40 empty) = %d; want 78", got)
	}
	if got := shortestPathMemoryOptimized(grid, 0); got != 78 {
		t.Fatalf("shortestPathMemoryOptimized(40x40 empty) = %d; want 78", got)
	}
}

func cloneGrid(grid [][]int) [][]int {
	copyOf := make([][]int, len(grid))
	for r := range grid {
		copyOf[r] = slices.Clone(grid[r])
	}
	return copyOf
}
