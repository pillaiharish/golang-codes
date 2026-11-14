# #525 Contiguous Array — 90-second cold-recall card

[Full explanation](README.md) · [Go solution](main.go) · [Tests](main_test.go) · [LeetCode index](../README.md)

**Trigger:** Longest *contiguous* span with equal counts of `0` and `1`.

**Transform:** Treat `0` as `-1`, `1` as `+1`. Equal prefix balances enclose a zero-sum span.

**Invariant:** `firstIndex[balance]` always holds the earliest index at which that balance appeared; `0: -1` represents the empty prefix.

| Step | Reconstruct | Why |
|---|---|---|
| 1 | `firstIndex := map[int]int{0: -1}` | Include balanced prefixes starting at index 0 |
| 2 | `balance, best := 0, 0` | Track prefix sum and longest interval |
| 3 | Loop indices: `1 → +1`, `0 → -1` | Equal counts become zero sum |
| 4 | If balance exists, maximize `i - firstIndex[balance]` | Same prefix balance means balanced middle |
| 5 | Else store the current index once | Earliest index makes the distance longest |

**Trace:** `[0,1,0,1]` has balances `-1,0,-1,0`; initial balance `0` at index `-1` gives length `3-(-1)=4`.

**Traps to avoid:** Missing the `0: -1` seed; overwriting the first index; reading absent map keys as index `0`; confusing a contiguous subarray with a subsequence.

**Complexity:** Expected O(n) time, O(n) auxiliary space.

**Cold drill:** Explain what `firstIndex` stores without naming it. Code the function from scratch. Check `[0,1]`, `[0,1,0]`, `[1,1,1,0,0]`, and all zeros.

Confidence: `—` · Last cold recall: `—` · Next recall: `—`
