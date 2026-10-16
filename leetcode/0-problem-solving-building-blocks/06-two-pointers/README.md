# 06 — Two pointers: converging indices

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognition and intuition
When a slice is sorted, you can use information from both ends to discard impossible pairs without a hash map.

## Operations to know first
- Initialize `left,right:=0,len(a)-1`; continue while `left<right`.
- If sum is too small, move left rightward; if too large, move right leftward.
- Invariant: a valid solution (if any) remains inside the current `[left,right]` search region.
- The sorted input of #167 makes these pointer moves valid; this is not the same algorithm as unsorted #1.

## Runnable Go micro-example
```go
package main
import "fmt"
func pairSorted(a []int,target int) []int {
    left,right:=0,len(a)-1
    for left<right {
        sum:=a[left]+a[right]
        if sum==target {return []int{left+1,right+1}}
        if sum<target {left++} else {right--}
    }
    return nil
}
func main(){fmt.Println(pairSorted([]int{2,7,11,15},9))}
```

## Hand trace
`[2,7,11,15]` target 9: `2+15=17` >9 → right--, `2+11=13` >9 → right--, `2+7=9` → return [1,2].

## Recall exercise (no notes)
For sorted `[-2,1,4,7]` and target 5, predict 1-based indices [2,3]. Why is O(1) extra memory enough?

## Common mistakes
- An unsorted array does not justify moving pointers by comparisons.
- LeetCode #167 returns 1-based indices, unlike #1.
- Two pointer progress is O(n) because left only increases and right only decreases.

## Stronger problems using this block
- [#167 Two Sum II - Input Array Is Sorted](../../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go)
- [#1 Two Sum](../../1.%20Two%20Sum/main.go)

## Review in 60 seconds
Name the invariant, rewrite the code without looking, then explain one mistake the tests should catch.
