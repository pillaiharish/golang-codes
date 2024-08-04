package main

import "fmt"

func canBeEqual(target []int, arr []int) bool {
	mapNumbers := make(map[int]int)
	for _, i := range target {
		mapNumbers[i] += 1
	}
	for _, i := range arr {
		// fmt.Println("Here ", i)
		mapNumbers[i] -= 1
		if mapNumbers[i] < 0 {
			// fmt.Println("Here")
			return false
		}
	}
	return true
}

func main() {

	target1, arr1 := []int{1, 2, 3, 4}, []int{2, 4, 1, 3}
	fmt.Println(canBeEqual(target1, arr1)) // Output: true

	target2, arr2 := []int{7}, []int{7}
	fmt.Println(canBeEqual(target2, arr2)) // Output: true

	target3, arr3 := []int{3, 7, 9}, []int{3, 7, 11}
	fmt.Println(canBeEqual(target3, arr3)) // Output: false
}

// Output:
// true
// true
// false
