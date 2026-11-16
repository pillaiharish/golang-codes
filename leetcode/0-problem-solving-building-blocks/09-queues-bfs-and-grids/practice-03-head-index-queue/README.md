# Practice 03 — head-index queue (Go)

[← Mastery card](../QUEUE-MASTERY.md) · [Original learner-written V1](../practice-01-dynamic-queue/README.md) · [Next: plain-grid BFS](../practice-05-grid-bfs/README.md)

## Problem

Your original queue uses `q.items=q.items[1:]` (correct O(1) reslicing). Write a second version that leaves the slice start unchanged, advances `head`, and uses `len(items)-head` for live length.

## Solution and why

[main.go](main.go) defines `Queue[T]` with `Enqueue`, `Dequeue`, and `Len`. The zero value is usable. The `(T,bool)` result distinguishes an empty queue from a valid zero value. Clearing a consumed slot prevents pointer retention; a fully drained queue resets its backing reference and head.

**Invariant:** `0 <= head <= len(items)` and live items are `items[head:]`.

### Hand trace

| Operation | Stored slice | head | live length |
|---|---|---:|---:|
| Enqueue 5,3,11 | [5 3 11] | 0 | 3 |
| Dequeue → 5 | [0 3 11] | 1 | 2 |
| Dequeue → 3 | [0 0 11] | 2 | 1 |
| Dequeue → 11 | nil (reset) | 0 | 0 |

Time: enqueue amortized O(1); dequeue O(1); length O(1). **Tradeoff:** when long-lived and not fully drained, consumed storage remains allocated. A compaction policy or circular queue solves that later; this version teaches the head-index mental model.

## Tests and run

```sh
GO111MODULE=off go run .
GO111MODULE=off go test -race -count=1 -v .
GO111MODULE=off go vet .
```

Tests cover empty, FIFO, full drain, reuse and clearing pointer-bearing slots. [source](main.go) · [tests](main_test.go).

## Cold recall

Hide main.go. Implement `Queue[int]` from scratch, then change it to store `*int`. Explain why `head++` does not mean the backing array shrinks. Do **not** replace the original V1; keep both variants for comparison.
