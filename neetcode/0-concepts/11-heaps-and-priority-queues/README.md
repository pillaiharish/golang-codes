# Heaps and the Go priority-queue interface

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A heap exposes one extreme value without keeping the whole array sorted. In a min-heap the parent is no larger than its children; this invariant does not order siblings or arbitrary positions. Typical operations are peek O(1), insertion/removal O(log n), and heap construction O(n).

Go container/heap uses a type implementing Len, Less, Swap, Push(any), and Pop() any. Your Push/Pop methods change the storage endpoints; heap.Push and heap.Pop perform the heap adjustment. Calling your endpoint method directly does not substitute for the package operation. A max-heap reverses Less. Tie-breaking fields make priority deterministic where required.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
h := &IntHeap{9, 2, 6}
heap.Init(h)
heap.Push(h, 4)
fmt.Println(heap.Pop(h).(int), (*h)[0])
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **2 4**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Change the heap comparator to a max-heap and predict the next two removals. Add a tie-break field to a record type.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Calling storage Pop directly; assuming the full array is sorted; retaining removed pointer entries; forgetting pointer receivers for resizing.

## Questions that use this building block

- [#23 Merge k Sorted Lists](../../23.%20Merge%20k%20Sorted%20Lists/README.md)
- [#703 Kth Largest Element in a Stream](../../703.%20Kth%20Largest%20Element%20in%20a%20Stream/README.md)
- [#1046 Last Stone Weight](../../1046.%20Last%20Stone%20Weight/README.md)
- [#973 K Closest Points to Origin](../../973.%20K%20Closest%20Points%20to%20Origin/README.md)
- [#215 Kth Largest Element in an Array](../../215.%20Kth%20Largest%20Element%20in%20an%20Array/README.md)
- [#621 Task Scheduler](../../621.%20Task%20Scheduler/README.md)
- [#355 Design Twitter](../../355.%20Design%20Twitter/README.md)
- [#295 Find Median from Data Stream](../../295.%20Find%20Median%20from%20Data%20Stream/README.md)
- [#332 Reconstruct Itinerary](../../332.%20Reconstruct%20Itinerary/README.md)
- [#1584 Min Cost to Connect All Points](../../1584.%20Min%20Cost%20to%20Connect%20All%20Points/README.md)
- [#743 Network Delay Time](../../743.%20Network%20Delay%20Time/README.md)
- [#778 Swim in Rising Water](../../778.%20Swim%20in%20Rising%20Water/README.md)
- [#269 Alien Dictionary](../../269.%20Alien%20Dictionary/README.md)
- [#787 Cheapest Flights Within K Stops](../../787.%20Cheapest%20Flights%20Within%20K%20Stops/README.md)
- [#1851 Minimum Interval to Include Each Query](../../1851.%20Minimum%20Interval%20to%20Include%20Each%20Query/README.md)

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
