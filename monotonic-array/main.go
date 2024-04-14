package main

import "fmt"

func isMonotonic(nums []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			isIncreasing = false
		} else if nums[i] < nums[i+1] {
			isDecreasing = false
		}
	}
	return isDecreasing || isIncreasing
}

func main() {
	fmt.Println(isMonotonic([]int{1, 1, 0}))
	fmt.Println(isMonotonic([]int{2, 2, 2, 1, 4, 5}))
}
