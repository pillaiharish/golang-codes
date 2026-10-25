# 128 — Longest Consecutive Sequence (Go)

[← LeetCode index](../README.md) · [Go solution with numbered comments](main.go) · [Table-driven tests](main_test.go) · [Quick revision card](REVISION-CARD.md) · **[My failed attempts + learned concepts](LEARNING-NOTES.md)** · [Official LeetCode problem](https://leetcode.com/problems/longest-consecutive-sequence/) · [Unsolved NeetCode practice starter](../../neetcode/128.%20Longest%20Consecutive%20Sequence/README.md) · [Revision log](../REVISION-LOG.md)

**Pattern:** Hash set + identify sequence starts · **Difficulty:** Medium · **Expected:** O(n) average time and O(n) space.

## Problem statement (paraphrased from LeetCode)

Given an unsorted slice of integers `nums`, return the **length** of the longest sequence of consecutive integer *values*. The values need not be adjacent in the input. Repeated values do **not** extend a sequence. The problem asks for an O(n)-time algorithm.

```go
func longestConsecutive(nums []int) int
```

**Constraints:** `0 <= len(nums) <= 100000`, and `-1000000000 <= nums[i] <= 1000000000`. The input need not be sorted; the function should not change it. Consult [LeetCode](https://leetcode.com/problems/longest-consecutive-sequence/) for the authoritative wording.

### Official examples and expected outputs

| Input (`nums`) | Output | Longest run |
|---|---:|---|
| `[100,4,200,1,3,2]` | `4` | `1,2,3,4` |
| `[0,3,7,2,5,8,4,6,0,1]` | `9` | `0,1,2,3,4,5,6,7,8` |
| `[1,0,1,2]` | `3` | `0,1,2` (duplicate `1` counts once) |

Additional tests cover empty/nil input, one value, only duplicates, reversed runs, negative-to-positive runs, disjoint runs, gaps, allowed extreme values, and a 100,000-value consecutive run.

## Before you solve: prerequisite basics

Work through these in order, explaining each Go operation out loud:

1. [01 — Arrays and slices](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md): `[]int`, `range` values, duplicates.
2. [02 — Loops and bounds](../0-problem-solving-building-blocks/02-loops-and-bounds/README.md): `for`, `continue`, loop invariants.
3. [03 — Maps, sets, frequency and grouping](../0-problem-solving-building-blocks/03-maps-sets-frequency-and-grouping/README.md): `make(map[int]bool)`, membership, zero values versus key existence.
4. [17 — Complexity and input contracts](../0-problem-solving-building-blocks/17-complexity-amortization-and-input-contracts/README.md): hash-map expectations and aggregate loop costs.

## 1. Recognition trigger

**Clue:** Unsorted numbers, *consecutive values* rather than consecutive positions, and an O(n) requirement. Sorting would make neighbors easy to compare, but costs O(n log n). Use a hash set to answer "does `x` exist?" in expected O(1).

**Mental shortcut:** **Only count a run at its beginning.** A number starts a run precisely when `num-1` is absent.

## 2. Plain-language intuition: how the approach evolved

**First idea (valid but slower):** Sort, ignore repeated numbers, count each consecutive run, and keep the largest count. It works but is O(n log n). You also do not need an array of counts; update one `maxLength` variable.

**Second idea (the observed failing experiment):** Put numbers into `map[int]bool`, walk backward from *every* element, and set visited entries to `false`. This loses connections: processing a number can mark part of a run visited before later starts discover the remainder. Removing the `false` assignments avoids that bug but repeatedly rescans prefixes and can take O(n²) time.

**Preferred idea:** Keep the map as an immutable membership set. Skip all numbers that have a predecessor. Only for numbers without predecessors, walk forward counting `num+1`, `num+2`, ... until a gap. Every value gets counted in at most one run.

### My mistakes and Go concepts to revisit

The full [learning notes](LEARNING-NOTES.md) preserve the original unsuccessful Go function **verbatim** and explain each bug with traces, not just the final fix:

- **Incorrect visited state:** `map1[k] = false` keeps the key in the map; `_, ok := map1[k-1]` still finds that key even though `map1[k-1]` may now be false. The original version returns **2** for `[1,2,3,4]` instead of 4.
- **One-direction traversal:** starting from a middle value and marking it visited while walking only backward can leave its forward neighbors disconnected.
- **Maximum update on just one exit path:** a `for map1[k]` loop may terminate without reaching the `largest` update. Moving the update outside the loop is needed but not sufficient without correcting the traversal invariant.
- **TLE after removing marking:** scanning a consecutive prefix again from each element creates `1+2+...+n` checks → O(n²), regardless of average O(1) map lookup.
- **Syntax and reasoning fundamentals:** `seen[k]` vs `_, ok := seen[k]`; `delete` vs setting false; map keys vs slice elements; map iteration order; zero values; `continue`; nested loops with amortized total O(n); keeping just a running maximum.

The notes also explain why a **two-direction visited** variant is valid, while the preferred immutable-set version is simpler to reason about.

## 3. Invariant / why it is correct

Every maximal consecutive run has **exactly one smallest element** `s` for which `s-1` is absent. The outer loop will eventually see `s`, regardless of Go's random map iteration order. The inner loop then counts every element of that run until the first missing successor. Any other member has a predecessor and is skipped. Therefore each run is counted exactly once, and the maximum count is the answer.

Duplicates become one map key and cannot inflate the length. No visited markers or input mutations are needed.

## 4. Five-step reconstruction recipe

1. Build `seen := make(map[int]bool, len(nums))`; set `seen[num] = true` for each input.
2. Initialize `maxLength := 0` (correct for empty input).
3. Visit **unique map keys** with `for num := range seen`.
4. If `seen[num-1]` is true, `continue`. Otherwise start at `current := num`, `length := 1` and increment both while `seen[current+1]`.
5. Update `maxLength` after the walk. Return it.

**Why these steps matter:** the set removes duplicates; starting the answer at zero handles empty input; iterating map keys avoids repeated duplicates; testing `num-1` prevents repeated expansion from middle values; updating the maximum *after* the complete run prevents the early-exit bug from the [first failed attempt](LEARNING-NOTES.md).

## 5. Hand trace

Input: `[100,4,200,1,3,2]`; the membership set is `{1,2,3,4,100,200}`.

| Candidate key | Is predecessor present? | Work | Longest seen |
|---|---|---|---:|
| `1` | `0` absent | Count `1 -> 2 -> 3 -> 4`, stop at `5` | `4` |
| `2`, `3`, `4` | Yes | Skip: these are inside the same run | `4` |
| `100` | `99` absent | Count only `100` | `4` |
| `200` | `199` absent | Count only `200` | `4` |

This table groups keys for explanation; **the actual map iteration order can differ** and the answer remains `4`.

## 6. Complexity and mistakes to avoid

- **Expected O(n) time:** O(n) to insert keys; O(u) predecessor checks for `u <= n` unique keys; O(u) inner-loop increments **combined** because the maximal runs are disjoint. Assumes average O(1) hash-map operations. This is not a deterministic worst-case bound for Go maps.
- **O(n) extra space:** map storing at most `n` distinct integers; no sorting or auxiliary list of run lengths.
- **Key existence versus stored value:** If `seen[2] = false`, then `_, ok := seen[2]` still gives `ok == true` but `seen[2]` is `false`. This solution never writes `false`, so direct `seen[x]` is a valid membership check.
- **Do not expand from every element:** For `[1,2,3,4]`, repeated forward/backward scans cost `4+3+2+1`, leading to quadratic growth.
- **Do not rely on map iteration order:** Go ranges over map keys in unspecified order. Sequence-start detection is independent of that order.
- **Do not count duplicates twice:** Looping over map keys and using a set automatically removes duplicates.
- **Update the maximum after finishing a run:** Don't update only inside a special `break` branch; that can miss a completed count.
- **Integer bounds:** The official limits of +/-1e9 make `num-1` and `current+1` safe for Go `int`. For unrestricted machine-sized extremes, arithmetic-overflow guards would be needed.

### Alternative techniques

| Idea | Time | Extra space | Trade-off |
|---|---|---|---|
| Sort + skip duplicates + count | O(n log n) | Depends on sort/copy choice | Easy first implementation; misses required O(n) |
| Hash set + scan from every number | O(n²) worst case | O(n) | Repeated work despite O(1) lookup |
| Hash set + only start without predecessor (**this solution**) | Expected O(n) | O(n) | Simple, non-mutating, proof via unique sequence starts |
| Mark visited + expand both directions | Expected O(n) | O(n) | Also correct, but requires careful state mutation |

## 7. 30-second explanation and blank-screen recall

> I put all distinct numbers into a hash set. For each number, if its predecessor exists, it cannot be the start of a run, so I skip it. Otherwise I walk through consecutive successors, count the run and update the maximum. Each unique value belongs to only one counted run, so total average time is O(n), with O(n) space.

**Before opening the Go implementation, answer these from memory:**

1. Why can `3` be skipped if `2` exists, even if `1` has not been processed yet?
2. Why is the *sum* of all `for seen[current+1]` iterations O(n), not O(n²)?
3. What is the difference between `if seen[k]` and `if _, ok := seen[k]; ok` when some values are false?
4. What causes TLE if you scan backwards from every input value but remove the visited marker?
5. How do `[1,0,1,2]`, `[]`, and `[-2,-1,0,1]` behave?

**Personal review:** Confidence `— / 5` · Last cold recall `—` · Next recall `—`. Record genuine attempts in the [revision log](../REVISION-LOG.md).

## Run locally and submit

From the repository root:

```sh
cd 'leetcode/128. Longest Consecutive Sequence'
GO111MODULE=off go test -count=1 -race -v
GO111MODULE=off go run main.go
```

The standalone demo prints `4`. The test file runs all three official examples plus boundary/regression tests. On LeetCode submit only `longestConsecutive` from [main.go](main.go), without the demo `main()` or `fmt` import. The [separate NeetCode starter](../../neetcode/128.%20Longest%20Consecutive%20Sequence/README.md) stays intentionally unsolved for future cold practice.
