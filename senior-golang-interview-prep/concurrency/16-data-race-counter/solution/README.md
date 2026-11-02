# Q16. Create and fix a data race — worked solution

[Back to Q16 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q17: shared maps](../../17-protect-shared-map/README.md) · [Revision notes](../../NOTES.md)

**Level 4 — shared state and synchronization.**

## Task

Start 1,000 goroutines that each increment one shared counter. Observe the incorrect implementation using `go run -race main.go` in your own experiment. Fix it first with a mutex and then using atomics.

## Reference answers

**Broken:** `counter++` expands into a read, update and write; concurrent updates race and can be lost.

```go
var mu sync.Mutex
counter := 0
// Inside each goroutine:
mu.Lock()
counter++
mu.Unlock()
// After wg.Wait(): fmt.Println(counter) // 1000
```

**Atomic alternative:**

```go
var counter atomic.Int64
// Inside each goroutine: counter.Add(1)
// After wg.Wait(): fmt.Println(counter.Load()) // 1000
```

Create the `WaitGroup` and join all workers in both versions. The mutex protects the entire read-modify-write operation; typed atomic methods perform an atomic update. The race detector finds observed data races in an executed schedule, but a clean run is not a proof that every possible schedule is safe.

## Scenario Q&A

| Question | Answer |
|---|---|
| Why doesn't `wg.Wait()` alone protect `counter++`? | It joins completion, but does not serialize simultaneous worker writes. |
| When use atomics versus mutex? | Atomics suit simple independent values and carefully designed lock-free protocols; mutexes suit multi-field invariants, maps and compound operations. |
| Can you read `counter` while workers update? | Only with corresponding synchronization (lock or atomic load), not a plain read. |
| Does `counter++` always produce a wrong number without a mutex? | No. A race may coincidentally produce the expected value; correctness is still not guaranteed. |

**Recall:** Contrast data race with higher-level race condition; explain happens-before for mutex unlock/lock and `WaitGroup.Done`/`Wait`.

References: [Go race detector](https://go.dev/doc/articles/race_detector), [`sync/atomic`](https://pkg.go.dev/sync/atomic).
