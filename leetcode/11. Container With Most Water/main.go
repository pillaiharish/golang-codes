package main

import "fmt"

// maxArea returns the maximum water contained by two vertical lines.
// Go 1.21+ is required for the built-in min and max used here.
// For Go 1.20 or earlier, see maxAreaLegacy in variants_pre_go121.go:
// it uses explicit comparisons but the same O(n) two-pointer algorithm.
//
// Invariant: after measuring the current area, the shorter endpoint can
// be discarded because retaining it while reducing width cannot increase
// the area. Moving the shorter side is therefore always safe.
// Time: O(n); additional space: O(1).
func maxArea(height []int) int {
	area := 0
	n := len(height)
	left, right := 0, n-1
	for left < right {
		area = max(area, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
}
