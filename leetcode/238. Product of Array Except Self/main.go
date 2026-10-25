package main

import "fmt"

// productExceptSelf returns the product of every other element for each index.
// It uses no division; its output slice is not counted as auxiliary space.
// Time O(n), additional space O(1) beyond the returned slice.
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	// Step 1: Sweep left to right, storing products strictly BEFORE index i.
	// Why: Multiplying nums[i] only after writing ans[i] excludes itself.
	prefix := 1 // Multiplicative identity: the empty left side has product 1.
	for i := 0; i < n; i++ {
		ans[i] = prefix
		prefix *= nums[i]
	}

	// Step 2: Sweep right to left, multiplying by products strictly AFTER i.
	// Why: ans[i] already contains the left product, so this completes it.
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suffix // Exclude nums[i] again before updating suffix.
		suffix *= nums[i]
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4})) // [24 12 8 6]
}
