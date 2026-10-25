package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"official positive", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"official zero and negative", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"two values", []int{5, 7}, []int{7, 5}},
		{"one zero", []int{1, 0, 3, 4}, []int{0, 12, 0, 0}},
		{"two zeroes", []int{0, 2, 0, 4}, []int{0, 0, 0, 0}},
		{"all negative", []int{-1, -2, -3}, []int{6, 3, 2}},
		{"ones", []int{1, 1, 1}, []int{1, 1, 1}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := append([]int(nil), tc.nums...)
			if got := productExceptSelf(tc.nums); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if !reflect.DeepEqual(tc.nums, original) {
				t.Fatalf("input mutated from %v to %v", original, tc.nums)
			}
		})
	}
}
