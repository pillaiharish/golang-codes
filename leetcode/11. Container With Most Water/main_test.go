package main

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"official example", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"two lines", []int{1, 1}, 1},
		{"equal tallest endpoints", []int{4, 3, 2, 1, 4}, 16},
		{"small peak", []int{1, 2, 1}, 2},
		{"all zero", []int{0, 0, 0}, 0},
		{"all equal", []int{5, 5, 5, 5}, 15},
		{"increasing", []int{1, 2, 3, 4, 5}, 6},
		{"decreasing", []int{5, 4, 3, 2, 1}, 6},
		{"zeros at edges", []int{0, 3, 4, 0}, 3},
		{"nil input outside constraints", nil, 0},
		{"single input outside constraints", []int{8}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.input)
			implementations := []struct {
				name string
				fn   func([]int) int
			}{
				{"Go 1.21+ builtins", maxArea},
				{"legacy comparisons", maxAreaLegacy},
				{"brute force", maxAreaBruteForce},
			}
			for _, impl := range implementations {
				if got := impl.fn(tt.input); got != tt.want {
					t.Fatalf("%s(%v) = %d; want %d", impl.name, tt.input, got, tt.want)
				}
			}
			if !slices.Equal(tt.input, original) {
				t.Fatal("an implementation modified the input")
			}
		})
	}
}

func TestMaxAreaRandomizedAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(11))
	for trial := 0; trial < 1500; trial++ {
		a := make([]int, rng.Intn(40))
		for i := range a {
			a[i] = rng.Intn(100)
		}
		want := maxAreaBruteForce(a)
		for _, impl := range []struct {
			name string
			fn   func([]int) int
		}{
			{"Go 1.21+ builtins", maxArea},
			{"legacy comparisons", maxAreaLegacy},
		} {
			if got := impl.fn(a); got != want {
				t.Fatalf("trial %d, %s, heights=%v: got %d; want %d", trial, impl.name, a, got, want)
			}
		}
	}
}

func TestMaxAreaLongInput(t *testing.T) {
	heights := make([]int, 100000)
	for i := range heights {
		heights[i] = 10000
	}
	want := 10000*(len(heights)-1)
	for _, impl := range []struct {
		name string
		fn   func([]int) int
	}{
		{"Go 1.21+ builtins", maxArea},
		{"legacy comparisons", maxAreaLegacy},
	} {
		if got := impl.fn(heights); got != want {
			t.Fatalf("%s: got %d; want %d", impl.name, got, want)
		}
	}
}
