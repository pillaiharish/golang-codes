# 1 — Two Sum

[← Index](../README.md) · [Building block: previously seen map](../0-problem-solving-building-blocks/README.md#1-a-hash-map-remembers-information-you-already-saw) · [Go solution](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/two-sum/)

## Recognition: how do I know it's this pattern?

**Clue:** “find two distinct indices whose values add to a target” in an *unsorted* array. A brute-force pair check takes O(n²). If I already knew the complementary number's index, I could solve this in one scan.

**Mental trigger:** **Target minus current = value I hope I have seen.** Use a map: number → earliest/previous index.

## Plain-language intuition

Given `nums[i] = v`, the pair must satisfy `v + other = target`; therefore `other = target - v`. Instead of searching the full array, maintain a dictionary of numbers observed **before** `i`.

Do *not* put the current value in the map before checking: `[3,3], target=6` needs two positions, not one position matched to itself.

## Invariant and why the algorithm works

**Before handling index i**, every map entry describes an index **strictly earlier than i**. So a successful lookup proves both indices are distinct. If a pair exists, when we reach the later index, its earlier partner is already stored.

## Memorize this five-step recipe

1. Make `map[int]int`: value → index.
2. Visit each pair `i, value`.
3. Compute `need := target - value`.
4. If `need` exists in the map, return `[storedIndex, i]`.
5. Otherwise remember `value → i`; return nil if none found (outside the standard guarantee).

### Trace: `[2,7,11,15]`, target=9

| i | value | need | Map *before* lookup | Action |
|---|---:|---:|---|---|
| 0 | 2 | 7 | `{}` | save 2→0 |
| 1 | 7 | 2 | `{2:0}` | find 2 at index 0; return [0,1] |

### Trick case: `[3,3]`, target=6

At i=0, `need=3` is not yet stored; save 3→0. At i=1, `need=3` exists with index 0. Answer [0,1].

## Complexity

- Average/expected **O(n)** time: n map lookups and inserts, typically O(1) each; no guaranteed worst-case O(1) hash lookup.
- **O(n)** extra space in the worst case: at most n seen values.

## Alternative and boundary

- Brute force checks each `i<j`: O(n²) time, O(1) extra space. Useful as a correctness oracle.
- Sorting + two pointers changes original order/indices unless we keep pairs with original indices (O(n log n)). For an **already sorted** input, see [Two Sum II #167](../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go), which uses O(1) extra space and returns **1-based** indices. Do not confuse #1 and #167.
- Input constraints promise a solution; `nil` is a documented fallback rather than a special LeetCode answer.

## 30-second spoken answer

> I scan once, keeping a map from each value seen so far to its index. For the current value I compute the complement, target minus value. If it was already seen, I return the stored index and my current index. I check before inserting so I cannot reuse the same position. Expected O(n) time and O(n) space.

## Recall quiz — close the solution first

1. Why is the map value an **index**, not just `bool`?
2. Why lookup **before** insert? Show what fails with `[3,3]`.
3. What are the differences between #1 and #167?
4. Reconstruct the solution with no notes, then write two tests.

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—` (record real attempts in the [log](../REVISION-LOG.md)).
