package main

import (
	"fmt"
	"slices"
)

// InsertAt inserts at index (including len) and reports false for invalid indices.
// It may reuse the original backing array, so callers with aliases must clone first.
func InsertAt[T any](items []T, index int, value T) ([]T, bool) {
	if index < 0 || index > len(items) {
		return items, false
	}
	var zero T
	items = append(items, zero)
	copy(items[index+1:], items[index:len(items)-1])
	items[index] = value
	return items, true
}

// DeleteAt preserves order and clears the old tail, including pointer-bearing values.
// It may modify aliases to the same backing array.
func DeleteAt[T any](items []T, index int) ([]T, bool) {
	if index < 0 || index >= len(items) {
		return items, false
	}
	copy(items[index:], items[index+1:])
	var zero T
	items[len(items)-1] = zero
	return items[:len(items)-1], true
}

// FilterInPlace retains matching items in original order, then clears unused slots.
func FilterInPlace[T any](items []T, keep func(T) bool) []T {
	write := 0
	for _, item := range items {
		if keep(item) {
			items[write] = item
			write++
		}
	}
	clear(items[write:]) // Go 1.21+
	return items[:write]
}

// DeduplicateStable returns unique values in first-seen order; it does not mutate input.
func DeduplicateStable[T comparable](items []T) []T {
	seen := make(map[T]struct{}, len(items))
	result := make([]T, 0, len(items))
	for _, value := range items {
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

// PopLast implements the stack-style pop operation using the end of a slice.
// The removed slot is cleared so pointer-bearing values are not retained.
func PopLast[T any](items []T) ([]T, T, bool) {
	var zero T
	if len(items) == 0 {
		return items, zero, false
	}
	last := len(items) - 1
	value := items[last]
	items[last] = zero
	return items[:last], value, true
}

// BinarySearch assumes ascending sorted input and returns the index and presence.
func BinarySearch(items []int, target int) (int, bool) {
	return slices.BinarySearch(items, target)
}

// ReverseInPlace mutates the provided slice (and any aliases).
func ReverseInPlace[T any](items []T) {
	slices.Reverse(items)
}

// SortedCopy leaves the input ordering untouched.
func SortedCopy(items []int) []int {
	result := slices.Clone(items)
	slices.Sort(result)
	return result
}

// DeepCloneMatrix copies outer and inner slice storage, but not referenced objects
// that may be inside arbitrary element types (this example's elements are ints).
func DeepCloneMatrix(matrix [][]int) [][]int {
	if matrix == nil {
		return nil
	}
	result := make([][]int, len(matrix))
	for r := range matrix {
		result[r] = slices.Clone(matrix[r])
	}
	return result
}

func main() {
	items := []int{3, 1, 3, 2}
	inserted, _ := InsertAt(slices.Clone(items), 2, 9)
	deleted, _ := DeleteAt(inserted, 2)
	fmt.Println(items, deleted, SortedCopy(items)) // [3 1 3 2] [3 1 3 2] [1 2 3 3]
}
