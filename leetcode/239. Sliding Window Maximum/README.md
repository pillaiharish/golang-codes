# 239 — Sliding Window Maximum

## Before you solve: prerequisite basics

Study these in sequence and type the tiny Go example in each before reopening this solution:

1. [01 arrays and slices](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md)
2. [08 sliding window and monotonic deque](../0-problem-solving-building-blocks/08-sliding-window-and-monotonic-deque/README.md)
3. [12 circular buffers and deques](../0-problem-solving-building-blocks/12-circular-buffers-and-deques/README.md)
4. [17 complexity amortization and input contracts](../0-problem-solving-building-blocks/17-complexity-amortization-and-input-contracts/README.md)

[Full mapping for all 23 LeetCode problems](../0-problem-solving-building-blocks/README.md#b-problem--exact-prerequisites-all-23-problem-folders) · [Practice log](../REVISION-LOG.md)

[← Index](../README.md) · [Building block: monotonic deque](../0-problem-solving-building-blocks/README.md#4-slices-indices-and-queue-behavior) · [Go solution](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/sliding-window-maximum/)

## Recognition

**Clue:** “maximum of every contiguous window of size k.” Recalculating each max takes O(nk). We must reuse knowledge from the previous window as its leftmost item exits and one new item enters.

**Mental trigger:** **Monotonic deque of indices: front is best; back loses to newer, larger values.**

## Intuition in a sentence

Keep only the indices that could still become maximum. If the newly arrived value is greater than or equal to the last candidate's value, the older candidate is never useful again: it is both smaller/equal **and** expires sooner.

## The two invariants

1. Deque indices are in increasing order and belong to the current window.
2. Their corresponding **values decrease from front to back**. The front is always the current maximum.

The deque holds **indices**, not values, because a value alone cannot tell whether it has expired.

## Recipe (same order every time)

1. Prepare results and an empty deque of indices.
2. For incoming index `i`, evict expired front index if `index < i-k+1` (equivalently `index <= i-k`).
3. While back value is `<= nums[i]`, pop it: new element dominates it.
4. Push `i` at the back.
5. Once `i>=k-1`, append `nums[deque[0]]` to results.

### Trace: `nums=[1,3,-1,-3,5]`, `k=3`

| i | incoming | Expire / dominate | Deque indices → values | Output |
|---|---:|---|---|---|
| 0 | 1 | nothing | [0] → [1] | — |
| 1 | 3 | pop 0 (1 <= 3) | [1] → [3] | — |
| 2 | -1 | nothing | [1,2] → [3,-1] | 3 |
| 3 | -3 | nothing | [1,2,3] → [3,-1,-3] | 3 |
| 4 | 5 | expire 1; pop 3,2 from back | [4] → [5] | 5 |

Continue the full input `[1,3,-1,-3,5,3,6,7]` → `[3,3,5,5,6,7]`.

## The proof you should remember

Each index enters the deque once and can be removed at most once. Therefore the total number of deque insertions/removals is O(n), despite the nested `for`. Every discarded back candidate is dominated by a newer candidate that is at least as large and survives at least as long.

## Complexity and Go-specific trap

- **Amortized O(n) time**.
- **O(k) live deque indices** as an abstract data structure, plus O(n-k+1) result.
- **Go implementation nuance:** `deque=deque[1:]` is O(1) for that operation but retains its backing array; repeatedly reslicing the front and appending can grow allocations/capacity beyond k over a long stream. Thus the exact simple slice implementation is **not guaranteed O(k) allocated auxiliary memory**. For hard memory bounds, use a circular-buffer deque or controlled compaction.
- Standard LeetCode requires `1<=k<=n`; this Go version defensively returns an empty result for invalid k.

## Alternatives

Naïve scan per window: O(nk). Heap with lazy stale removal: O(n log n) but more state. Monotonic deque is the standard O(n) solution.

## 30-second spoken answer

> I keep a deque of indices whose values are decreasing. For each incoming number, I remove expired indices from the front and remove smaller-or-equal indices from the back because the new value dominates them. The front is always the max once a full window exists. Every index is pushed and popped at most once, so time is amortized linear.

## Cold recall questions

1. Why indices rather than values?
2. Why pop smaller/equal entries from the **back**, not front?
3. What exact inequality makes an index expired?
4. Why is a nested while loop still O(n)?
5. Why might Go's naïve slice-based deque retain more memory than O(k)?

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—`.
