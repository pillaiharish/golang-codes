package main

import "fmt"

// maxSlidingWindow uses a monotonic deque of INDICES, never raw values.
// The deque is maintained in decreasing order by nums[index].
// Time: amortized O(n). The simple Go slice implementation below may retain
// extra backing-array capacity due to repeated front reslices; see README.
func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    if k <= 0 || k > n {
        return []int{} // outside standard problem constraints
    }
    result := make([]int, 0, n-k+1)
    deque := make([]int, 0, k)
    for i := range nums {
        // Evict the index that no longer belongs to the window.
        if len(deque) > 0 && deque[0] <= i-k {
            deque = deque[1:]
        }
        // The newer value dominates smaller-or-equal older candidates.
        for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        if i >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }
    return result
}

func main() {
    fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) // [3 3 5 5 6 7]
}
