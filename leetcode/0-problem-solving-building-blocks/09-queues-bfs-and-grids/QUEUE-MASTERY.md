# Queue mastery — FIFO, from zero values to BFS

[Mastery system](../MASTERY-SYSTEM.md) · [Combined lesson](README.md) · [Grid](GRID-MASTERY.md) · [BFS](BFS-MASTERY.md) · [Original #1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md)

## Quick recall — 60 seconds

**Queue = FIFO**: enqueue at the back, dequeue at the front. Unlike a stack, the earliest inserted live element leaves first.

```go
queue := []int{10}
head := 0
for head < len(queue) {
    current := queue[head]
    head++
    _ = current
    // queue = append(queue, discoveredNeighbor)
}
```

**Invariant:** live elements are `queue[head:]`. The next dequeue is `queue[head]`. **Trap:** advancing `head` avoids shifting but does **not** automatically reclaim consumed capacity.

## A–B. Identity, representation, zero values, creation

A queue is an **abstract data structure**. Go does not have a built-in queue type; implement FIFO with a slice, head index, or ring buffer. An unbounded slice-backed queue is appropriate for beginner BFS.

| Representation | Zero value | Constructor? | When to choose |
|---|---|---|---|
| `[]T` and reslice | nil slice, usable | No | Small one-off BFS, learn FIFO |
| `struct{items []T; head int}` | zero, usable | No | Simple head-index queue; amortized appends |
| Fixed ring `[]T,head,size` | capacity 0; cannot enqueue | Usually yes | Bounded stream/queue; see [lesson 12](../12-circular-buffers-and-deques/README.md) |
| `container/list` | usable zero value | No | Stable list-element references; generally more allocation |

`var q Queue[int]` is valid for the head-index [generic version](practice-03-head-index-queue/main.go). In the learner's [original `dynamicQueue`](practice-01-dynamic-queue/main.go), `items` starts nil; `append` works. Generic `[T any]` means the type can store int, pointers, structs, etc.

## C. Full operations lifecycle

| Operation | Slice/head implementation | Complexity |
|---|---|---|
| Initialize | `var q Queue[int]`, or `queue := []int{}` | O(1) |
| Enqueue | `items = append(items, x)` | Amortized O(1) |
| Dequeue | read `items[head]`; clear; `head++` | O(1) |
| Peek | check nonempty, read `items[head]` | O(1) |
| Length / empty | `len(items)-head`, compare with zero | O(1) |
| Traverse FIFO | `for i:=head; i<len(items); i++` | O(n) |
| Search | linear scan over live elements | O(n) |
| Arbitrary insert/delete/sort | Not a normal queue API; use a slice if needed | N/A |

A queue is **ordered by insertion**, not numerically sorted. `len(items)` in a head-index queue includes already consumed slots; size is `len(items)-head`.

**Trace:** enqueue 5,3,11 → dequeue returns 5 → remaining 3,11 → dequeue returns 3 → empty-dequeue yields zero and false. `(T,bool)` differentiates empty from a legitimate zero element.

## D. Ownership, memory and cleanup

- A slice value copies its **header**; copies or subslices can share a backing array. Passing a slice to a function does not deep-copy elements.
- `items = items[1:]` is O(1), not O(n); however it may retain the original backing array.
- In V2, `head++` is O(1), but consumed slots still occupy capacity; **clear pointer-bearing slots** so unreachable objects are not retained. After full drain, reset the slice and head. A long-lived queue with interleaved enqueue/dequeue needs compaction or a ring (not provided by simple V2).
- The queue is **not safe for concurrent mutation**. Wrap a queue in synchronization if producers/consumers are concurrent.

## E. Invariants and Layer 3 traps

1. `0 <= head <= len(items)`.
2. `Len() == len(items)-head`; only `items[head:]` is live.
3. Empty dequeue must not panic; do **not** treat a returned 0 as proof of empty.
4. Snapshot BFS level size before enqueueing next-level work; otherwise level boundaries move.
5. Pointer/struct values and slice headers have different copy/ownership semantics.
6. Never assume a slice-based queue automatically reclaims storage as it drains.

## F. Progressive implementations (preserve original)

1. **V1 — learner-written** [dynamic queue](practice-01-dynamic-queue/README.md): a struct, pointer receiver, `append`, front reslicing. [Tests](practice-01-dynamic-queue/main_test.go).
2. **V2 — head-index** [practice and trace](practice-03-head-index-queue/README.md): `head`, `Len`, clearing pointer slots, full-drain reset. [Tests](practice-03-head-index-queue/main_test.go).
3. **V3 — circular queue/deque** [existing lesson](../12-circular-buffers-and-deques/README.md): bounded ring, modulo arithmetic, full/empty distinction. Study after V2.

No change to the learner's V1 source was necessary.

## G. LeetCode applications

**Clues:** shortest path with unit-cost edges → BFS queue; tree level order → BFS queue; multi-source spreading → initial queue has several seeds. Queue alone does **not** choose what "visited" means.

- [#102 level-order traversal](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/README.md)
- [#200 number of islands](../../200.%20Number%20of%20Islands/README.md)
- [#1293 obstacle-elimination grid](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md) (uses a three-part state)

## H–I. Design and optimization (later)

Future design: choose `(T,bool)` for empty, hide internal storage, document ownership, compare bounded ring vs growable slice. Future profiling: allocation/retention, backing-array growth, compaction and benchmark against ring. **Do not optimize before reproducing correct FIFO behavior.**

## J. Recall and tests

1. From a blank editor, write enqueue/dequeue with head index and zero-value support.
2. Explain why `queue=queue[1:]` is O(1) but can retain memory.
3. Demonstrate why snapshotting `levelEnd := len(queue)` is necessary.
4. Write tests for empty, one element, FIFO, drain/reuse and pointer clearance.
5. Transfer: apply the same six-line head-index pattern to [plain-grid BFS](practice-05-grid-bfs/README.md), then solve #1293 independently.

```sh
cd leetcode/0-problem-solving-building-blocks/09-queues-bfs-and-grids/practice-03-head-index-queue
GO111MODULE=off go test -race -count=1 .
GO111MODULE=off go vet .
```

Confidence: — · Last independent recall: — · Next review: —
