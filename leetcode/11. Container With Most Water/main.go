package main

import "fmt"

// maxArea returns the largest container formed by two vertical lines.
// Time: O(n); additional space: O(1).
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0
	for left < right {
		width := right - left
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * width
		if area > maxWater {
			maxWater = area
		}
		// Discard the shorter side: keeping it with a smaller width
		// cannot improve the best area found so far.
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxWater
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
}
