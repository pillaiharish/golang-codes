# 217 — Contains Duplicate

[← Index](../README.md) · [Building block: sets](../0-problem-solving-building-blocks/README.md#2-a-set-stores-existence-not-a-count) · [Go solution](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/contains-duplicate/)

## Recognition

**Clue:** “does any element appear at least twice?” We do not need *where*, *how many*, or *which* repeated value. We only need to know whether a value has already appeared.

**Mental trigger:** **Have I seen it before? → set.** In Go, use `map[int]struct{}` as a set.

## Intuition

Scan left to right. If the next number is in `seen`, we found a duplicate. If it is new, add it. Because each element is processed once, we can stop immediately when a duplicate appears.

## Invariant

**Before index i**, the set contains every distinct value from `nums[:i]`. Hence membership of `nums[i]` means it occurred earlier and the answer is true.

## Rebuild from scratch

1. Create `seen := make(map[int]struct{}, len(nums))`.
2. For each number, check `_, ok := seen[number]`.
3. If present, return true.
4. Otherwise assign `seen[number] = struct{}{}`.
5. End of scan → false.

### Trace: `[1,2,3,1]`

| Next number | Seen before | Resulting action |
|---|---|---|
| 1 | {} | insert 1 |
| 2 | {1} | insert 2 |
| 3 | {1,2} | insert 3 |
| 1 | {1,2,3} | already present; return true |

## Why `map[int]struct{}` rather than `map[int]int`?

`map[int]int` can solve the problem but suggests that you are counting occurrences or storing another piece of information. **An empty struct value has zero-sized payload**, so a map set expresses the intent and can use less payload memory than an integer map.

**Important correction to an oversimplified explanation:** The whole map entry is **not zero bytes**. Keys, metadata, buckets/groups and allocations still occupy memory. Also, `make(map[int]struct{}, len(nums))` gives a **hint**, not “exactly enough allocated buckets” and not a formal performance guarantee. It can reduce resizing when many distinct values are inserted but may over-allocate if a duplicate appears early.

`map[int]bool` is also a legitimate readable set when you need true/false flags. Use `struct{}` for existence-only semantics.

## Complexity

Expected O(n) time, O(n) additional space in the all-unique case. Sorting first gives O(n log n) time and may mutate the input; no sorting is needed for the set method.

## 30-second spoken answer

> I scan once with a hash set. Before inserting a number, I check whether it already exists; if yes, I return true immediately. If the whole input is unique, I return false. I use map[int]struct{} because I only need membership, and the length hint is an optional allocation optimization.

## Recall without notes

1. What information differs between a **set**, **counter**, and **value→index map**?
2. Is `struct{}` equivalent to a zero-memory map entry? Why not?
3. Why might `make(..., len(nums))` waste memory on some inputs?
4. Recode it, including empty and negative-number tests.

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—`.
