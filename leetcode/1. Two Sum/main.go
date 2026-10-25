package main

import "fmt"

// twoSum finds indices of two distinct positions that sum to target.
// Pattern: hash map of values seen BEFORE the current position.
// Expected time O(n), extra space O(n).
func twoSum(nums []int, target int) []int {
	// Step 1: Remember a number's index, not just whether it exists.
	// Why: the answer must contain two original indices.
	indexByValue := make(map[int]int, len(nums))
	for i, value := range nums {
		// Step 2: Derive the only value that can pair with nums[i].
		complement := target - value
		// Step 3: Lookup BEFORE insertion; any match is at an earlier index.
		// Why: using the same index twice is forbidden.
		if previous, ok := indexByValue[complement]; ok {
			return []int{previous, i}
		}
		// Step 4: Save the current index for future elements.
		indexByValue[value] = i
	}
	// Step 5: Only reached when no pair exists (outside the judge guarantee).
	return nil // LeetCode 1 promises one solution; nil makes the fallback explicit.
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
}
