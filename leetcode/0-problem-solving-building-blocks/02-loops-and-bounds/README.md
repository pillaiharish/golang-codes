# 02 — Loops, range, indices, and bounds

[← All foundations](../README.md) · [← LeetCode index](../../README.md)

## Why this matters
Traversal is the prerequisite for brute force, sliding windows, two pointers, BFS grids and binary search. Learn exact loop stopping conditions before optimizations.

## Go operations: learn and explain
- `for i:=0;i<len(a);i++` provides index control; `for i,v := range a` gives index and value.
- Use `for _,v := range a` when index is unnecessary. Slice range values are copies.
- `for r := range grid { for c := range grid[r] {...} }` traverses a 2D slice.
- Guard `0 <= index && index < len(a)` before reading; slicing `[lo:hi]` uses exclusive hi.
- Nested loops over all subarray starts and ends are O(n²), but nested loops that move pointers forward can be amortized linear.

## Runnable example
```go
package main
import "fmt"
func main() {
    a := []int{2, 5, 8}
    for i, v := range a { if v % 2 == 0 { a[i] = v * 2 } }
    for start := 0; start < len(a); start++ {
        sum := 0
        for end := start; end < len(a); end++ {
            sum += a[end]
            fmt.Println(start, end, sum)
        }
    }
}
```

## Dry run
First loop transforms `[2,5,8]` → `[4,5,16]`. With start=0, running subarray sums are 4, 9, 25.

## Tiny exercise (do before opening a hard problem)
Write both an index-based loop and `range` loop to count values >60 in `[18,92,61]`. Why must you avoid `i<=len(a)`?

## Typical mistakes
- Writing `v=10` inside `range` does *not* change `a[i]`.
- Go map iteration has no stable order; do not derive an ordered result by ranging a map.
- When modifying a grid, distinguish `(r,c)` from `(x,y)` coordinate conventions.

## Apply to the existing repository
- [1. Two Sum](../../1.%20Two%20Sum/main.go)
- [102. Binary Tree Level Order Traversal](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go)
- [200. Number of Islands](../../200.%20Number%20of%20Islands/main.go)
- [2678. Number of Senior Citizens](../../2678.%20Number%20of%20Senior%20Citizens/main.go)
- [1508. Range Sum of Sorted Subarray Sums](../../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go)

## 30-second closed-book recall
Say what state this building block stores, how to traverse/update it, and which boundary case could invalidate your solution. Then rewrite the snippet without copying.
