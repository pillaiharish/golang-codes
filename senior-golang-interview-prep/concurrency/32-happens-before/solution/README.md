# Q32. Explain happens-before in your code — worked solution

[Back to Q32 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q16: data race](../../16-data-race-counter/README.md) · [Revision notes](../../NOTES.md)

**Level 6 — correctness and the Go memory model.**

## Task

For earlier exercises, explain how one goroutine knows another goroutine's writes are visible. Identify the *precise synchronization operation* creating each happens-before edge. Find an unsynchronized example and fix it.

## Reference answer

| Mechanism | Useful guarantee |
|---|---|
| Channel send/receive | A send synchronizes before completion of its matching receive; additional rules apply to unbuffered channels. |
| Channel close/receive | Closing synchronizes before a receive that returns because it observes the closed state. |
| `Mutex.Unlock` -> later `Mutex.Lock` | Establishes the lock-protected ordering edge. |
| `WaitGroup.Done` -> `Wait` it unblocks | Done synchronizes before the return from that Wait; it is **not** a worker-to-worker mutual exclusion mechanism. |
| Atomic operations | Have specified synchronization/ordering semantics; they do not automatically make multi-field invariants atomic. |
| Context cancellation | `Done()` channel closure provides a channel-close signal; merely passing a context does not synchronize unrelated shared variables. |

Example:

```go
var result int
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); result = 42 }()
wg.Wait()
fmt.Println(result) // write precedes observation after join
```

If `fmt.Println(result)` runs **before** `wg.Wait()`, that unsynchronized read may race with the write.

## Scenario Q&A

| Question | Answer |
|---|---|
| Is a scheduling delay a happens-before edge? | No. `Sleep` does not create a memory-ordering relationship. |
| Does `WaitGroup` make multiple concurrent writes safe? | No; workers still need exclusive ownership, atomics or locking for conflicting accesses. |
| Does an atomic counter protect a related map? | No. A separate invariant needs its own synchronization protocol. |
| Why do we care beyond the race detector? | Memory visibility and protocol ordering are correctness properties for **all** executions, not just observed runs. |

**Recall:** For Q1, Q3, Q7 and Q21, draw arrows indicating exactly which operations synchronize. Do not substitute wall-clock ordering for memory-model guarantees.

References: [The Go memory model](https://go.dev/ref/mem), [`sync.WaitGroup`](https://pkg.go.dev/sync#WaitGroup).
