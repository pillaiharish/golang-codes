package main

import "fmt"

// longestConsecutive returns the length of the longest run of consecutive
// integer values (not necessarily adjacent positions in nums).
// Expected time: O(n) with average O(1) map lookups; extra space: O(n).
func longestConsecutive(nums []int) int {
	// A map used as a set: duplicates collapse to one key.
	seen := make(map[int]bool, len(nums))
	for _, num := range nums {
		seen[num] = true
	}

	maxLength := 0
	// Map iteration order is unspecified and does not matter here.
	for num := range seen {
		// If a predecessor exists, num cannot start a sequence.
		if seen[num-1] {
			continue
		}

		current := num
		length := 1
		// Each unique value is expanded only from its sequence start.
		for seen[current+1] {
			current++
			length++
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // 4
}
