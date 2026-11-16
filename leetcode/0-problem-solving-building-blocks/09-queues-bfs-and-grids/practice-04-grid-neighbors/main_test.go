package main

import (
	"reflect"
	"testing"
)

func TestNeighbors(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		at   Cell
		want []Cell
	}{
		{"center", [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, Cell{1, 1}, []Cell{{0, 1}, {2, 1}, {1, 0}, {1, 2}}},
		{"ragged", [][]int{{0, 0, 0}, {0}}, Cell{0, 1}, []Cell{{0, 0}, {0, 2}}},
		{"edge", [][]int{{0, 0}, {0, 0}}, Cell{0, 0}, []Cell{{1, 0}, {0, 1}}},
		{"invalid", [][]int{{0}}, Cell{1, 0}, nil},
		{"empty", nil, Cell{0, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Neighbors(tt.grid, tt.at); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Neighbors = %v, want %v", got, tt.want)
			}
		})
	}
}
