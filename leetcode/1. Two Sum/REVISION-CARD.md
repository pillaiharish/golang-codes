# #1 Two Sum — 90-second cold recall card

[Full explanation](README.md) · [Commented code](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md)

**Trigger:** Two different original indices in an unsorted slice add to a target.

**Invariant:** Before visiting index `i`, `seen[value]` stores only indices **before** `i`.

| Step | Reconstruct from blank file | Why |
|---|---|---|
| 1 | `seen := make(map[int]int)` | Need original indices, not just existence |
| 2 | `for i, v := range nums` | Scan each current candidate once |
| 3 | `need := target - v` | Convert pair search into one membership lookup |
| 4 | Check `j, ok := seen[need]`; return `[j,i]` | Prior index is distinct from current index |
| 5 | Otherwise `seen[v] = i` | Make current value available to later candidates |

**Two tiny traces:** `[2,7,11,15]`, target 9 → `[0,1]`; `[3,3]`, target 6 → `[0,1]` **only if lookup comes before insertion**.

**Traps to explain:** `map[int]bool` cannot return an index; inserting before checking may reuse the same position; two pointers on sorted values do not preserve original indices without extra bookkeeping; #167 uses 1-based indices.

**Complexity:** expected O(n) time, O(n) extra space.

**Cold drill:** Implement in <10 minutes, then answer why `[3,3]` works. Run `GO111MODULE=off go test -race` from this folder. Record *actual* confidence/date in the [log](../REVISION-LOG.md), not here in advance.

Confidence: `—` · Last cold recall: `—` · Next recall: `—`
