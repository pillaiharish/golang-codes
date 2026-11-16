package main

import "testing"

func TestShortestGridPath(t *testing.T) {
	tests := []struct {
		name        string
		grid        [][]int
		start, goal Cell
		want        int
	}{
		{"path", [][]int{{0, 1, 0}, {0, 0, 0}, {1, 1, 0}}, Cell{0, 0}, Cell{2, 2}, 4},
		{"blocked", [][]int{{0, 1}, {1, 0}}, Cell{0, 0}, Cell{1, 1}, -1},
		{"single", [][]int{{0}}, Cell{0, 0}, Cell{0, 0}, 0},
		{"blocked start", [][]int{{1, 0}}, Cell{0, 0}, Cell{0, 1}, -1},
		{"invalid goal", [][]int{{0}}, Cell{0, 0}, Cell{1, 0}, -1},
		{"empty", nil, Cell{0, 0}, Cell{0, 0}, -1},
		{"ragged", [][]int{{0, 0}, {0}, {0, 0}}, Cell{0, 1}, Cell{2, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortestGridPath(tt.grid, tt.start, tt.goal); got != tt.want {
				t.Fatalf("ShortestGridPath = %d, want %d", got, tt.want)
			}
		})
	}
}
