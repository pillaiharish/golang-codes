package main

// maxAreaLegacy preserves the earlier main.go implementation unchanged apart
// from its name, so it can coexist with the new Go 1.21+ solution.
// It uses explicit comparisons rather than the built-in min/max functions.
// Time: O(n); additional space: O(1).
//
// For Go 1.20 and earlier, copy this function into the LeetCode editor and
// rename maxAreaLegacy to maxArea. Do not compile it with the new main.go,
// which requires Go 1.21+.
func maxAreaLegacy(height []int) int {
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
