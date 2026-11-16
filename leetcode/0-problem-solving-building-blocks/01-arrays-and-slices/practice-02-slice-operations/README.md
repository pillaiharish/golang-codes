# Practice 02 — slice operations from fundamentals to reusable blocks

[Slice mastery](../SLICE-MASTERY.md) · [Traps](../TRAPS-AND-MEMORY.md) · [Recall](../RECALL-DRILLS.md)

## Task

Implement `InsertAt`, `DeleteAt`, `PopLast`, `FilterInPlace`, `DeduplicateStable`, `SortedCopy`, `BinarySearch`, `ReverseInPlace` and `DeepCloneMatrix`.

| Operation | What happens | Edge to remember |
|---|---|---|
| Insert | append space, shift suffix right using overlapping `copy` | index in [0,n] |
| Delete | shift left, clear obsolete tail, shorten | invalid/empty and pointer retention |
| Pop | read last, clear, shorten | empty slice |
| Filter | separate read/write cursors; clear remainder | may modify aliases |
| Dedupe | set of comparable values, first-seen order | hash-key eligibility |
| Sort | clone before sorting when original must remain unchanged | sort mutates |
| Binary search | lookup in ascending sorted slice | unsorted is unsupported |
| Reverse | swap in place | modifies shared backing |
| Matrix clone | clone outer slice and each row | nested sharing |

### Hand trace

```text
original  [3 1 3 2]
insert 9 at index 2 -> [3 1 9 3 2]
delete index 2      -> [3 1 3 2]
SortedCopy          -> [1 2 3 3] (original unchanged)
```

**API design for this learning exercise:** invalid indices return `(unchangedSlice, false)`. `InsertAt`, `DeleteAt` and `FilterInPlace` can reuse and mutate shared backing storage; clone before calling if isolation is needed. `SortedCopy` and `DeduplicateStable` provide a new top-level slice. `DeepCloneMatrix` clones both levels for `[][]int`.

Time: insert/delete O(n), pop O(1), filter O(n), average dedupe O(n) time/O(n) extra, sort O(n log n), binary search O(log n), matrix clone O(total elements).

## Run

```sh
GO111MODULE=off go run .
GO111MODULE=off go test -race -count=1 -v .
GO111MODULE=off go vet .
```

[Complete source](main.go) · [Table-driven tests](main_test.go), including nil, invalid indices, pointer tail cleanup, nested clone and sorted search.

## Closed-book challenge

Implement insert/delete without `slices.Insert` / `slices.Delete`; explain overlapping `copy` and why the vacated tail is cleared. Implement an in-place stable filter with a write cursor. Transfer to existing [#41](../../../41.%20First%20Missing%20Positive/README.md) or [#238](../../../238.%20Product%20of%20Array%20Except%20Self/README.md).
