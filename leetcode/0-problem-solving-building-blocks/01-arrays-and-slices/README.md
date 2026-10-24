# 01 — Arrays and slices

[← All foundations](../README.md) · [← LeetCode index](../../README.md)

## Why this matters
Before any optimized solution, know indexing, length, capacity, traversal, swapping, and the difference between fixed arrays and dynamically sized slices.

## Go operations: learn and explain
- `var counts [26]int` creates a comparable, zero-initialized fixed array.
- `nums := []int{3,1,2}` creates a slice. `nums[i]` indexes, `nums[i:j]` reslices, `len(nums)` measures elements.
- `make([]int,0,n)` reserves capacity but leaves length zero; append grows length.
- `append` requires a **slice**: `out := []string{}` followed by `out = append(out, "hi")`. A `string` is not a slice; concatenate strings with `+`/`+=`, or use `strings.Builder` for repeated writes.
- `var out []string` creates a nil slice, while `out := []string{}` creates a non-nil empty slice; both can be appended to. `var out string` creates one empty string, not a result collection.
- `for i,v := range nums` traverses indices and *copies* of values. Write to `nums[i]` to modify.
- `nums[i], nums[j] = nums[j], nums[i]` swaps values, used for cyclic placement.

## Runnable example
```go
package main
import "fmt"
func main() {
    nums := []int{3, 1, 2}
    for i, v := range nums { fmt.Println(i, v) }
    nums[0], nums[1] = nums[1], nums[0]
    output := make([]int, 0, len(nums))
    output = append(output, nums[0])
    var counts [3]int
    counts[0]++
    fmt.Println(nums, output, counts) // [1 3 2] [1] [1 0 0]
}
```

## Dry run
Swap `[3,1,2]` at indices 0,1 → `[1,3,2]`. `output` length changes from 0 to 1; counts starts zero.

## Tiny exercise (do before opening a hard problem)
Swap first and last of `[9,2,5]` → `[5,2,9]`; explain the length/capacity of `make([]int,0,5)` after 2 appends.

## Typical mistakes
- `append("hello", "world")` does not compile: its first argument must be a slice, not a string.
- `var out string` cannot be returned from a function declared to return `[]string`; read the function's return type before choosing a variable type.
- Index `i == len(nums)` is out of bounds.
- Reslicing does not copy; backing-array memory can remain retained.
- An array such as `[26]int` may be used as a map key; `[]int` cannot.

## Apply to the existing repository
- [1. Two Sum](../../1.%20Two%20Sum/main.go)
- [41. First Missing Positive](../../41.%20First%20Missing%20Positive/main.go)
- [94. Binary Tree Inorder Traversal](../../94.%20Binary%20Tree%20Inorder%20Traversal/main.go)
- [1508. Range Sum of Sorted Subarray Sums](../../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go)
- [239. Sliding Window Maximum](../../239.%20Sliding%20Window%20Maximum/main.go)
- [641. Design Circular Deque](../../641.%20Design%20Circular%20Deque/main.go)

## 30-second closed-book recall
Say what state this building block stores, how to traverse/update it, and which boundary case could invalidate your solution. Then rewrite the snippet without copying.
