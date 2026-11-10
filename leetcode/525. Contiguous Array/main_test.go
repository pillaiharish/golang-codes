package main

import (
	"math/rand"
	"slices"
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "LeetCode example 1", nums: []int{0, 1}, want: 2},
		{name: "LeetCode example 2", nums: []int{0, 1, 0}, want: 2},
		{name: "repeat earliest balance", nums: []int{0, 1, 0, 1}, want: 4},
		{name: "nil input", nums: nil, want: 0},
		{name: "empty input", nums: []int{}, want: 0},
		{name: "single zero", nums: []int{0}, want: 0},
		{name: "single one", nums: []int{1}, want: 0},
		{name: "all zeros", nums: []int{0, 0, 0, 0}, want: 0},
		{name: "all ones", nums: []int{1, 1, 1}, want: 0},
		{name: "best subarray starts after index zero", nums: []int{1, 1, 1, 0, 0}, want: 4},
		{name: "valid subarray at end", nums: []int{1, 1, 0, 0, 1}, want: 4},
		{name: "balanced whole array", nums: []int{0, 0, 1, 0, 1, 1}, want: 6},
		{name: "one repeated balance without equality", nums: []int{1, 0, 0}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.nums)
			got := findMaxLength(tt.nums)
			if got != tt.want {
				t.Fatalf("findMaxLength(%v) = %d; want %d", tt.nums, got, tt.want)
			}
			if !slices.Equal(tt.nums, original) {
				t.Fatalf("input mutated: got %v; before %v", tt.nums, original)
			}
		})
	}
}

func TestFindMaxLengthAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(525))
	for trial := 0; trial < 300; trial++ {
		nums := make([]int, rng.Intn(31))
		for i := range nums {
			nums[i] = rng.Intn(2)
		}
		if got, want := findMaxLength(nums), bruteForceMaxLength(nums); got != want {
			t.Fatalf("trial %d, nums=%v: got %d; want %d", trial, nums, got, want)
		}
	}
}

func TestFindMaxLengthLargeAlternating(t *testing.T) {
	const n = 100000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i % 2
	}
	if got := findMaxLength(nums); got != n {
		t.Fatalf("findMaxLength(100000 alternating values) = %d; want %d", got, n)
	}
}

// An intentionally simple oracle for testing, not the submitted algorithm.
func bruteForceMaxLength(nums []int) int {
	best := 0
	for start := range nums {
		balance := 0
		for end := start; end < len(nums); end++ {
			if nums[end] == 1 {
				balance++
			} else {
				balance--
			}
			if balance == 0 && end-start+1 > best {
				best = end - start + 1
			}
		}
	}
	return best
}
