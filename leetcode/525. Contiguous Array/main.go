package main

import "fmt"

// findMaxLength returns the length of the longest contiguous subarray with
// equal numbers of zeros and ones. nums must contain only 0 and 1.
// Expected time: O(n) with average O(1) map lookups; extra space: O(n).
func findMaxLength(nums []int) int {
	// A balance of zero exists just before the first element. This lets us
	// measure a qualifying subarray beginning at index 0.
	firstIndex := map[int]int{0: -1}
	balance := 0
	best := 0

	for i, value := range nums {
		// Treat 0 as -1 and 1 as +1.
		if value == 1 {
			balance++
		} else {
			balance--
		}

		// Equal prefix balances enclose a zero-sum subarray.
		if start, seen := firstIndex[balance]; seen {
			if length := i - start; length > best {
				best = length
			}
		} else {
			// Keep the earliest index to maximize every later distance.
			firstIndex[balance] = i
		}
	}

	return best
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1, 0, 1})) // 4
}
