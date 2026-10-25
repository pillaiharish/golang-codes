package main

import "fmt"

// longestConsecutive returns the length of the longest run of consecutive
// integer values (not necessarily adjacent positions in nums).
// Expected time: O(n) with average O(1) map lookups; extra space: O(n).
func longestConsecutive(nums []int) int {
	// Step 1: A map used as a set; duplicates collapse to one key.
	// Why: consecutive SEQUENCE length counts distinct numeric values.
	seen := make(map[int]bool, len(nums))
	for _, num := range nums {
		seen[num] = true
	}

	// Step 2: Empty input has no consecutive run.
	maxLength := 0
	// Map iteration order is unspecified and does not matter here.
	for num := range seen {
		// Step 3: If a predecessor exists, num cannot start a sequence.
		// Why: scanning only starts prevents repeated work and O(n^2) time.
		if seen[num-1] {
			continue
		}

		// Step 4: Expand once from the unique start of this run.
		current := num
		length := 1
		// Count each successor; no visited flags or input mutation needed.
		for seen[current+1] {
			current++
			length++
		}
		// Step 5: Save the longest completed run (outside the inner loop).
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
