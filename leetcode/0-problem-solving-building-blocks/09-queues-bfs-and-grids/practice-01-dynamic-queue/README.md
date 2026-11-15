# Practice 01 — Dynamic Queue, using your Go struct

[← Lesson 09](../README.md) · [Your complete Go code](main.go) · [Tests](main_test.go) · [Next: BFS](../practice-02-graph-bfs/README.md)

## Recognition and intuition

A **queue** is FIFO (first in, first out). The first item added is the first removed. This is the mechanism that lets BFS explore nodes in increasing distance order.

This `dynamicQueue` code is preserved from your exercise, including your method names, printed demo, and `err` variable. Only HTML-escaped brackets and pointer symbols were decoded, and whitespace normalized for valid Go. Unlike class-based OOP, Go uses a struct with **methods and pointer receivers** to express an object-like abstraction.

## Operations and invariant

- `items []int` is a dynamically growing slice of integers. Its zero value is usable as an empty queue.
- `Enqueue` uses `append` to add to the **back**.
- `Dequeue` checks `len(q.items)==0` before accessing `q.items[0]`, reads the **front** element, then drops it using `q.items = q.items[1:]`.
- The receiver `*dynamicQueue` modifies the original struct rather than a copy.
- `Dequeue() (int, bool)` returns **(removedValue, success)**. The `bool` is **not a Go error**. The variable `err` in your original demo would be more accurately named `ok`, but is intentionally unchanged.

**Invariant:** `items[0]`, when present, is always the oldest unremoved item.

## Trace

| Action | Queue slice | Return |
|---|---|---|
| `dynamicQueue{}` | `[]` | — |
| `Enqueue(5)` | `[5]` | — |
| `Enqueue(3)` | `[5 3]` | — |
| `Enqueue(11)` | `[5 3 11]` | — |
| `Dequeue()` | `[3 11]` | `(5, true)` |

Your demo prints:

```text
Out [5 3 11]
First Dequeue 5 [3 11]
```

**Why the boolean?** Empty `Dequeue()` returns `(0,false)`, but enqueueing an actual `0` and removing it returns `(0,true)`.

## Complexity and Go-specific caveats

- `Enqueue`: amortized O(1); `Dequeue`: O(1) slice reslicing (does **not** shift all items).
- The slice may **retain the old backing array** after many dequeues. For a long-lived service, a circular buffer or head-index queue may better control memory.
- This queue is **not thread-safe**. Ordinary single-goroutine BFS doesn't need synchronization.

## Try it yourself

1. Add the numbers `5,3,11` and predict which exits first.
2. Dequeue until empty: explain why the return value is `0,false`.
3. Enqueue `0`: how do you distinguish it from an empty queue?
4. Compare a `*dynamicQueue` receiver with a value receiver.
5. Use a `[]int` FIFO directly in [your BFS implementation](../practice-02-graph-bfs/main.go).

```sh
cd leetcode/0-problem-solving-building-blocks/09-queues-bfs-and-grids/practice-01-dynamic-queue
GO111MODULE=off go run main.go
GO111MODULE=off go test -race -count=1 -v .
```

**Connection to [#1293](../../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md):** Your queue uses `int` as its element type. A grid shortest-path queue instead stores a **state struct** containing the row, column, and remaining obstacle eliminations.
