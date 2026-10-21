# 03 — Maps, sets, frequencies and grouping

[← All foundations](../README.md) · [← LeetCode index](../../README.md)

## Why this matters
The same Go map implements four different patterns: membership, value→index, frequency counts, and canonical-key grouping. Know which information a problem needs.

## Go operations: learn and explain
- `m := make(map[int]int)` creates a writable map. A nil map can be read but cannot be assigned to.
- `value, ok := m[key]` distinguishes missing from stored zero; use `delete(m,key)` to remove.
- `map[int]struct{}` is a membership set; `struct{}{}` carries zero-sized value payload but map metadata still occupies memory.
- `count[x]++` relies on zero-value reads for missing keys.
- `groups[key] = append(groups[key], item)` builds map-of-slices.
- `make(map[int]int,n)` is a capacity *hint*, not fixed capacity or proof of no further allocations.

## Runnable example
```go
package main
import "fmt"
func main() {
    previous := map[int]int{7: 0}
    pos, ok := previous[7] // pos=0 but ok=true
    seen := map[int]struct{}{7: {}}
    _, repeated := seen[7]
    count := map[string]int{}
    count["go"]++
    count["go"]++
    groups := map[string][]string{}
    groups["aet"] = append(groups["aet"], "eat", "tea")
    fmt.Println(pos, ok, repeated, count["go"], groups["aet"])
}
```

## Dry run
Map lookup 7 returns `(0,true)`, so testing only `value!=0` would be wrong. Counter `go` becomes 2. Bucket `aet` contains two original strings.

## Tiny exercise (do before opening a hard problem)
Choose one map type for Two Sum, Contains Duplicate, Valid Anagram, Group Anagrams and Logger Rate Limiter, and explain why each value type differs.

## Typical mistakes
- Maps are not safe for unsynchronized concurrent writes.
- A set using `struct{}` does not make a whole map entry zero bytes.
- Never rely on iteration order of `map` when tests expect grouping or sorting.

## Apply to the existing repository
- [1. Two Sum](../../1.%20Two%20Sum/main.go)
- [49. Group Anagrams](../../49.%20Group%20Anagrams/main.go)
- [128. Longest Consecutive Sequence](../../128.%20Longest%20Consecutive%20Sequence/README.md) — use a membership set; only start walks where the predecessor is missing.
- [146. LRU Cache](../../146.%20LRU%20Cache/main.go)
- [1460. Make Two Arrays Equal by Reversing Subarrays](../../1460.%20Make%20Two%20Arrays%20Equal%20by%20Reversing%20Subarrays/main.go)
- [217. Contains Duplicate](../../217.%20Contains%20Duplicate/main.go)
- [242. Valid Anagram](../../242.%20Valid%20Anagram/main.go)
- [353. Design Snake Game](../../353%20-%20Design%20Snake%20Game/main.go)
- [359. Logger Rate Limiter](../../359.%20Logger%20Rate%20Limiter/main.go)

## 30-second closed-book recall
Say what state this building block stores, how to traverse/update it, and which boundary case could invalidate your solution. Then rewrite the snippet without copying.
