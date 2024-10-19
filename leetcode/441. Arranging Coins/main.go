package main

import "fmt"

func arrangeCoins(n int) int {
	left, right := 0, n

	for left <= right {
		mid := left + (right-left)/2
		coinsUsed := mid * (mid + 1) / 2

		if coinsUsed == n {
			return mid
		} else if coinsUsed < n {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return right
}

func main() {
	n := 8
	output := arrangeCoins(n)
	fmt.Println("For n=8 the complete rows formed is", output)
}

// Output
// For n=8 the complete rows formed is 3