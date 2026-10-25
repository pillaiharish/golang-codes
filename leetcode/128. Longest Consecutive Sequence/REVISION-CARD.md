# #128 Longest Consecutive Sequence — 90-second cold recall card

[Full explanation](README.md) · [Original mistakes](LEARNING-NOTES.md) · [Commented code](main.go) · [Tests](main_test.go)

**Trigger:** Unsorted integers, longest run of consecutive **values**, expected O(n) requirement.

**Invariant:** Expand only from a value without a predecessor; every maximal run has exactly one such start.

| Step | Reconstruct | Why |
|---|---|---|
| 1 | Fill `map[int]bool` with all numbers | Deduplicate and answer membership quickly |
| 2 | Initialize `largest := 0` | Empty input must return 0 |
| 3 | Iterate **map keys** | Each distinct value visited once, arbitrary order |
| 4 | `if seen[n-1] { continue }` | A predecessor means this is not a run start |
| 5 | From `n`, count while `seen[current+1]` | Scan that run once; update max *after* inner loop |

**Trace:** `[100,4,200,1,3,2]` → start 1 counts 1,2,3,4 → answer 4.

**Your earlier traps:** `seen[k]=false` does not remove map key; `_,ok := seen[k]` remains true; marking a backward-only partial run loses connections; removing marking but scanning from every value does O(n²) total work; inner loop does **not** automatically mean quadratic time.

**Complexity:** expected O(n) total time (each unique value expanded once), O(n) map space.

**Cold drill:** Trace the failed approach on `[1,2,3,4]` (it returned 2), reconstruct the corrected algorithm, then test duplicates and negatives. [Full failure analysis](LEARNING-NOTES.md).

Confidence: `—` · Last cold recall: `—` · Next recall: `—`
