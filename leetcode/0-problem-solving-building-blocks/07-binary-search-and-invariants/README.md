# 07 — Binary search and monotonic predicates

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognition and intuition
Binary search requires an ordered or monotonic decision, not necessarily a simple sorted list. Find a boundary by proving which half can be discarded.

## Operations to know first
- Set `left,right` bounds; midpoint `left+(right-left)/2` avoids unnecessary addition overflow.
- For rotated sorted arrays without duplicates, compare `nums[mid]` with `nums[right]` to find the minimum.
- For a monotonic feasibility predicate `possible(k)`, find the last valid k or first invalid k.
- A `left<right` boundary search and `left<=right` exact-match search need different update rules.

## Runnable Go micro-example
```go
package main
import "fmt"
func minRotated(a []int) int {
    left,right:=0,len(a)-1
    for left<right {
        mid:=left+(right-left)/2
        if a[mid]>a[right] {left=mid+1} else {right=mid}
    }
    return a[left] // precondition: non-empty, distinct values
}
func main(){fmt.Println(minRotated([]int{4,5,6,7,0,1,2}))}
```

## Hand trace
`[4,5,6,7,0,1,2]`: mid=3 (7)>2 → left=4; mid=5 (1)<=2 → right=5; mid=4 (0)<=1 → right=4; answer 0.

## Recall exercise (no notes)
For arranging 8 coins, check triangular counts k(k+1)/2: 3 rows use 6, 4 rows need 10 → return 3.

## Common mistakes
- #153's implementation assumes nonempty and no duplicates (as per its problem contract).
- `mid*(mid+1)` can overflow an `int`; use a safe comparison/64-bit or division when generalizing.
- Binary search does not work unless you can justify a monotone side to discard.

## Stronger problems using this block
- [#153 Find Minimum in Rotated Sorted Array](../../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/main.go)
- [#441 Arranging Coins](../../441.%20Arranging%20Coins/main.go)

## Review in 60 seconds
Name the invariant, rewrite the code without looking, then explain one mistake the tests should catch.
