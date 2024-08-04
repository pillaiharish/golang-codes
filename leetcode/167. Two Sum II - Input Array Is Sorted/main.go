package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(numbers, target)
	fmt.Println(result) // Output: [1, 2]
}
