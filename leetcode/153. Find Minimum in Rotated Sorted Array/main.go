package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// If mid element is greater than the rightmost element, the minimum must be in the right half
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// Minimum could be the mid itself, so we include mid in our search
			right = mid
		}
	}
	return nums[left] // The left pointer will point to the minimum element
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("Minimum element is:", findMin(arr)) // Output: 0
}
