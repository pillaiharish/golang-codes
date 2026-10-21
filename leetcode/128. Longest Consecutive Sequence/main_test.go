package main

import (
	"slices"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "LeetCode example 1", nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{name: "LeetCode example 2", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
		{name: "LeetCode example 3", nums: []int{1, 0, 1, 2}, want: 3},
		{name: "nil input", nums: nil, want: 0},
		{name: "empty slice", nums: []int{}, want: 0},
		{name: "single number", nums: []int{42}, want: 1},
		{name: "only duplicates", nums: []int{2, 2, 2}, want: 1},
		{name: "already sorted", nums: []int{1, 2, 3, 4}, want: 4},
		{name: "reversed", nums: []int{4, 3, 2, 1}, want: 4},
		{name: "unsorted run", nums: []int{3, 1, 4, 2, 5}, want: 5},
		{name: "negative through zero", nums: []int{-2, 1, -1, 0, 5}, want: 4},
		{name: "disconnected runs", nums: []int{10, 8, 9, 1, 2, 3}, want: 3},
		{name: "gap breaks run", nums: []int{1, 2, 4, 5, 6}, want: 3},
		{name: "valid extreme values", nums: []int{-1000000000, -999999999, 1000000000}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.nums)
			got := longestConsecutive(tt.nums)
			if got != tt.want {
				t.Fatalf("longestConsecutive(%v) = %d; want %d", tt.nums, got, tt.want)
			}
			if !slices.Equal(tt.nums, original) {
				t.Fatalf("input was mutated: got %v, original %v", tt.nums, original)
			}
		})
	}
}

func TestLongestConsecutiveLongRun(t *testing.T) {
	// Exercises linear-growth behavior on the maximum input length.
	const n = 100000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = n - 1 - i
	}
	if got := longestConsecutive(nums); got != n {
		t.Fatalf("longestConsecutive(long run) = %d; want %d", got, n)
	}
}
