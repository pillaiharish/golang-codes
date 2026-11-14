package main

// maxAreaBruteForce tests every pair for a simple correctness reference.
// Time: O(n^2); additional space: O(1).
func maxAreaBruteForce(height []int) int {
	maxWater := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			h := height[i]
			if height[j] < h {
				h = height[j]
			}
			area := h * (j - i)
			if area > maxWater {
				maxWater = area
			}
		}
	}
	return maxWater
}
