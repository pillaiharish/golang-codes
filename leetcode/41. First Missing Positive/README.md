# 41 — First Missing Positive

[← Index](../README.md) · [Building block: index placement](../0-problem-solving-building-blocks/README.md#8-cyclic-placement-the-index-is-the-answers-address) · [Go solution](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/first-missing-positive/)

## Recognition

**Clue:** "smallest missing positive integer", expected **O(n) time** and **O(1) extra space**. Sorting violates time; using a set violates auxiliary-space requirements.

**Mental trigger:** **A positive value x belongs at array index x−1.** Use the input array itself as a placement table.

## Plain-language intuition

For n items, the answer must be in `[1,n+1]`. Values <= 0 or > n cannot fill any of the first n expected positions, so ignore them. Repeatedly swap each eligible value into its expected slot, then scan for the first incorrect slot.

The method **mutates the input**. That is essential to achieving O(1) extra space.

## Loop invariant / progress argument

Every successful swap places **at least one in-range value into its correct index**. Once a position contains the correct value, it need not move away for a different value: the check `nums[nums[i]-1] != nums[i]` prevents swapping a duplicate onto itself. Thus the total number of successful swaps is O(n).

## Reconstruct the recipe

1. Let `n := len(nums)`.
2. For each position `i`, while current value x is in `1..n` **and** `nums[x-1] != x`, swap `nums[i]` with `nums[x-1]`.
3. Check the same `i` again after each swap, since a different number has arrived.
4. After placement, scan `i=0..n-1`; first `nums[i] != i+1` → return `i+1`.
5. If all match, return `n+1`.

### Trace: `[3,4,-1,1]`

| Step | Array | Why |
|---|---|---|
| Start | [3,4,-1,1] | n=4 |
| i=0; 3 belongs at index 2 | [-1,4,3,1] | place 3 |
| i=1; 4 belongs at index 3 | [-1,1,3,4] | place 4 |
| i=1; 1 belongs at index 0 | [1,-1,3,4] | place 1 |
| scan index 1 | [1,-1,3,4] | expected 2, found -1 → answer 2 |

## Compare your source solution with this version

Your supplied solution used `nums[i] < n` as the upper check. This new implementation uses the conventional **`nums[i] <= n`**. The previous `< n` variation can still find the first missing positive: placing values 1 through n−1 is enough to determine the answer, and the remaining slot gives n by elimination. The `<= n` version makes the *explicit position x−1 for every x in 1..n* invariant easier to recite.

Both require the **target-slot inequality check** to avoid looping forever on duplicate values (e.g. `[1,1]`).

## Complexity and limitations

- **O(n)** time: not O(n²) despite nested loops, because successful swaps make irreversible placement progress.
- **O(1)** auxiliary space, excluding input and a handful of local variables.
- Input is rearranged; copy it first if mutation is prohibited, which costs O(n) extra memory.

## Alternatives

Sort: O(n log n). Build set of positives and scan from 1: expected O(n) time but O(n) space. Index placement satisfies both strict constraints.

## 30-second spoken answer

> For n entries, the missing answer lies between 1 and n+1. I use the input array as a map of value to index: x belongs at x−1. I swap in-range values to their proper places, skipping duplicates already sitting at their destinations. Every swap fixes at least one position so total time is O(n). Finally I return the first index where nums[i] isn't i+1, otherwise n+1. Extra space is constant.

## Cold recall

1. Why can values <= 0 and > n be ignored?
2. Why is the inner loop a **while**, not an if?
3. Why does `nums[target] != nums[i]` stop infinite swaps?
4. Why is the overall loop O(n)?
5. Explain what changes if preserving input order is required.

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—`.
