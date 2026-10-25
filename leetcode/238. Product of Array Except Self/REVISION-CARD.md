# #238 Product of Array Except Self — 90-second cold recall card

[Full reasoning](README.md) · [Commented code](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md)

**Trigger:** Every position needs an aggregate of everything **except itself**, without division.

**Invariant:** On the forward pass, `prefix` contains strictly earlier values. On the backward pass, `suffix` contains strictly later values.

| Step | Reconstruct | Why |
|---|---|---|
| 1 | `ans := make([]int,len(nums))` | Reuse required output for prefix products |
| 2 | `prefix := 1` | Empty product is 1 |
| 3 | Forwards: `ans[i]=prefix`, then `prefix*=nums[i]` | Excludes current element on the left |
| 4 | `suffix := 1`; loop backwards | Accumulate all elements strictly to the right |
| 5 | `ans[i]*=suffix`, then `suffix*=nums[i]` | Excludes current element on the right |

**Trace:** `[1,2,3,4]` → prefix answer `[1,1,2,6]` → final `[24,12,8,6]`.

**Your earlier approach:** Repeated adjacent products/pass-by-pass multiplication can revisit too many elements. Two running passes keep it O(n) while still using prior work.

**Traps:** `prefix`/`suffix` must start at 1; both accumulator updates occur **after** using them at index `i`; division fails at zeros and is forbidden; no extra prefix/suffix arrays are required.

**Cold drill:** Reconstruct both loops and trace `[1,0,3,4]`; explain why output counts as O(n) but auxiliary space is O(1).

Confidence: `—` · Last cold recall: `—` · Next recall: `—`
