package main

import "fmt"

// twoSum finds indices of two distinct positions that sum to target.
// Pattern: hash map of values seen BEFORE the current position.
// Expected time O(n), extra space O(n).
func twoSum(nums []int, target int) []int {
    indexByValue := make(map[int]int, len(nums))
    for i, value := range nums {
        complement := target - value
        if previous, ok := indexByValue[complement]; ok {
            return []int{previous, i}
        }
        indexByValue[value] = i
    }
    return nil // LeetCode 1 promises one solution; nil makes the fallback explicit.
}

func main() {
    fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
}
