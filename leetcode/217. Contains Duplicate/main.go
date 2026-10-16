package main

import "fmt"

// containsDuplicate uses a set: we only need membership, not counts.
// Expected O(n) time and O(n) auxiliary space.
func containsDuplicate(nums []int) bool {
    seen := make(map[int]struct{}, len(nums))
    for _, number := range nums {
        if _, exists := seen[number]; exists {
            return true
        }
        seen[number] = struct{}{}
    }
    return false
}

func main() {
    fmt.Println(containsDuplicate([]int{1, 2, 3, 1})) // true
}
