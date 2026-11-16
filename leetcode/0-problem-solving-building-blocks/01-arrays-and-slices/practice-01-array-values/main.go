package main

import "fmt"

// ChangeCopy receives and returns an entire array value. The caller is unchanged.
func ChangeCopy(original [3]int) [3]int {
	original[0] = 99
	return original
}

// ChangeThroughPointer modifies the caller's array instead.
func ChangeThroughPointer(original *[3]int) {
	original[0] = 99
}

// ArrayKey demonstrates that an array is a map key if all its elements are comparable.
func ArrayKey(values [3]int) int {
	counts := map[[3]int]int{values: 1}
	counts[values]++
	return counts[values]
}

func main() {
	original := [3]int{1, 2, 3}
	changed := ChangeCopy(original)
	fmt.Println(original, changed, ArrayKey(original)) // [1 2 3] [99 2 3] 2
}
