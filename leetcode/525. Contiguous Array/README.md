# 525 — Contiguous Array (Go)

[← LeetCode index](../README.md) · [Go solution](main.go) · [Table-driven tests](main_test.go) · [Cold-recall card](REVISION-CARD.md) · [Official problem](https://leetcode.com/problems/contiguous-array/) · [Revision log](../REVISION-LOG.md)

**Pattern:** Prefix sum / running balance + earliest-index hash map · **Difficulty:** Medium · **Target:** expected O(n) time, O(n) extra space.

## Problem statement (paraphrased)

Given a binary array `nums`, return the length of the **longest contiguous** subarray containing the same number of zeros and ones. A subarray uses adjacent positions, not arbitrary elements.

```go
func findMaxLength(nums []int) int
```

**Constraints:** `1 <= len(nums) <= 100000` and every `nums[i]` is either `0` or `1`. The implementation also returns `0` for an empty or nil slice, though those inputs are outside the stated constraints. It does not modify `nums`.

| Input | Output | Explanation |
|---|---:|---|
| `[0,1]` | `2` | Both entries form a balanced subarray |
| `[0,1,0]` | `2` | Either adjacent pair contains one zero and one one |
| `[0,1,0,1]` | `4` | The whole array is balanced |
| `[1,1,1,0,0]` | `4` | The longest balanced subarray is `[1,1,0,0]` |

## Prerequisites to review

1. [Arrays and slices](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md): positions and contiguous ranges.
2. [Loops and bounds](../0-problem-solving-building-blocks/02-loops-and-bounds/README.md): `for i, value := range nums` and index differences.
3. [Maps and key existence](../0-problem-solving-building-blocks/03-maps-sets-frequency-and-grouping/README.md): `map[int]int`, comma-ok lookup, preserving a first-seen value.
4. [Subarrays and prefix sums](../0-problem-solving-building-blocks/15-subarrays-running-sums-and-modular-arithmetic/README.md): equal prefix sums cancel over the interval between them.
5. [Complexity and constraints](../0-problem-solving-building-blocks/17-complexity-amortization-and-input-contracts/README.md): average-case map lookup cost and linear scans.

## 1. Recognition trigger

**Clues:** The question asks for the *longest contiguous* interval satisfying an equality between counts. Directly checking all subarrays works, but would take O(n²) time even when counts are maintained incrementally.

Replace each `0` with `-1` and each `1` with `+1`. Now a balanced subarray is precisely a **zero-sum subarray**. Equal prefix balances mark its boundaries.

**Memory hook:** Same balance twice → zero net change between them → equal zeros and ones.

## 2. Intuition: balance plus the first time it appeared

Walk left to right, keeping `balance = onesSeen - zerosSeen`:

- Reading `1` adds `1`; reading `0` subtracts `1`.
- If balance is `b` at an earlier index `j` and again at current index `i`, the segment `nums[j+1 : i+1]` sums to zero.
- Its length is `i - j`, so save **the earliest index** for each balance to maximize the length.
- Balance `0` exists *before* the array, at index `-1`. Seed `firstIndex[0] = -1` so a valid prefix from index `0` is measured correctly.

The map records `balance -> earliest index`, **not** a count and not the most recent index. Using a single balance variable is enough; separate zero/one counters would carry redundant state.

## 3. Invariant and correctness argument

After processing index `i`, `balance` equals the sum of the transformed values in `nums[0:i+1]`. For every balance seen so far, `firstIndex` stores its **first** occurrence, including the virtual `(balance=0, index=-1)` before the input.

If the current balance was seen at `j`, then the transformed sum from `j+1` through `i` is `balance(i) - balance(j) = 0`: that subarray has equal zeros and ones. Choosing the earliest such `j` yields the longest valid subarray ending at `i`. If the current balance is new, saving `i` keeps the invariant true. Every valid subarray must have equal balances at its endpoints, so checking each endpoint and taking the maximum finds the global optimum.

## 4. Five-step reconstruction recipe

1. Make `firstIndex := map[int]int{0: -1}` to represent the balance before the array.
2. Start `balance := 0` and `best := 0`.
3. Scan `nums` by index; add one for `1` and subtract one for `0`.
4. If `firstIndex[balance]` exists, update `best` with `i - firstIndex[balance]`.
5. Otherwise, store `firstIndex[balance] = i` **once**. Return `best` after the scan.

The `else` is essential: a later occurrence must never overwrite the earliest index.

## 5. Hand trace

Input: `[0,1,0,1]`. Initially `firstIndex = {0: -1}`, `balance = 0`, `best = 0`.

| Index | Value | New balance | Earliest index of balance | Candidate length | Best |
|---:|---:|---:|---|---:|---:|
| 0 | 0 | -1 | New: store `-1 → 0` | — | 0 |
| 1 | 1 | 0 | `-1` (seeded) | `1 - (-1) = 2` | 2 |
| 2 | 0 | -1 | `0` (do not overwrite) | `2 - 0 = 2` | 2 |
| 3 | 1 | 0 | `-1` (still earliest) | `3 - (-1) = 4` | 4 |

Answer: `4`. The equal balance at indices `-1` and `3` proves the entire array is balanced.

## 6. Complexity, boundaries, and traps

- **Expected time O(n):** one loop of `n` values, with average O(1) hash-map lookup/insertion. Map operations do not promise deterministic worst-case O(1).
- **Extra space O(n):** at most `n + 1` distinct prefix balances in `firstIndex`.
- **Seed `0: -1`:** without it, a balanced subarray starting at index `0` may be missed.
- **Do not overwrite:** keeping the latest index shortens later intervals; keep the earliest.
- **Check map presence:** `index, ok := firstIndex[balance]` distinguishes a missing key from a stored index `0`. Directly reading a missing `map[int]int` key returns `0`.
- **Maintain a running maximum:** the final balance tells only whether the whole array is balanced, not the longest balanced *portion*.
- **Binary input required:** treating every non-`1` value as `-1` is valid only under the stated `0`/`1` constraint.
- **Even length:** any balanced nonempty subarray must have even length. When no such interval exists, answer `0`.

**Alternative for learning:** Enumerate all starts and extend each end while tracking the balance; this is straightforward O(n²) time with O(1) extra space. The prefix-balance map avoids revisiting ranges and is the preferred solution here.

## 7. Interview explanation and blank-screen recall

> I treat zero as minus one and one as plus one. The problem becomes finding the longest zero-sum subarray. I walk once with a prefix balance and map each balance to its earliest index, starting with balance zero at index minus one. Seeing a balance again means the values in between sum to zero, so I update the maximum distance. I never overwrite the earliest index. This takes expected O(n) time and O(n) space.

Before looking at [main.go](main.go), answer:

1. Why does equal prefix balance imply equal zeros and ones in the middle?
2. Why is `firstIndex[0] = -1`, not `0`?
3. Why does retaining the first occurrence give the longest candidate?
4. What is wrong with assuming a missing `map[int]int` key means index `0` was stored?
5. For `[1,1,1,0,0]`, why is the answer `4` even though the final balance is not zero?

**Personal review:** Confidence `— / 5` · Last cold recall `—` · Next recall `—`. Record only genuine recall attempts in the [revision log](../REVISION-LOG.md).

## Run locally and submit

From the repository root:

```sh
cd 'leetcode/525. Contiguous Array'
GO111MODULE=off go test -count=1 -race -v
GO111MODULE=off go run main.go
```

The demo prints `4`. Submit only `findMaxLength` from [main.go](main.go) on LeetCode; omit the demo `main` and `fmt` import.
