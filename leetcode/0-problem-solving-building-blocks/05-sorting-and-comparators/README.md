# 05 — Sorting and custom comparators

[← All foundations](../README.md) · [← LeetCode index](../../README.md)

## Why this matters
Sorting can canonicalize anagrams, rank votes, sort interval endpoints, or order subarray sums. Learn sort APIs and why tie-breakers matter.

## Go operations: learn and explain
- `sort.Ints(a)` sorts ints ascending in place; `sort.Strings(a)` sorts strings.
- `sort.Slice(items, func(i,j int) bool {...})` orders by a custom criterion.
- Comparator must implement a consistent strict less-than relation; avoid `<=`.
- For multiple criteria, compare primary first, then secondary on ties.
- Sorting mutates the original slice, so copy if preserving input order matters.

## Runnable example
```go
package main
import ("fmt";"sort")
func main() {
    intervals := [][]int{{1,4},{2,5},{3,5},{1,3}}
    sort.Slice(intervals, func(i,j int) bool {
        if intervals[i][1] != intervals[j][1] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] > intervals[j][0]
    })
    fmt.Println(intervals) // [[1 3] [1 4] [3 5] [2 5]]
}
```

## Dry run
End ascending puts `[1,3]` first. When ends tie at 5, start descending places `[3,5]` before `[2,5]`.

## Tiny exercise (do before opening a hard problem)
Why does interval #757 require descending starts on ties? Show how a two-field comparator handles rank counts then alphabetical team names.

## Typical mistakes
- A comparator that is not transitive may produce unexpected results.
- `sort.Slice` changes its input; copy first if needed.
- Hash grouping and sorting are alternatives, not interchangeable for every problem.

## Apply to the existing repository
- [49. Group Anagrams](../../49.%20Group%20Anagrams/main.go)
- [2273. Find Resultant Array After Removing Anagrams](../../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go)
- [757. Set Intersection Size At Least Two](../../757.%20Set%20Intersection%20Size%20At%20Least%20Two/main.go)
- [1366. Rank Teams by Votes](../../1366.%20Rank%20Teams%20by%20Votes/main.go)
- [1508. Range Sum of Sorted Subarray Sums](../../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go)

## 30-second closed-book recall
Say what state this building block stores, how to traverse/update it, and which boundary case could invalidate your solution. Then rewrite the snippet without copying.
