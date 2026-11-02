# Q3. Prove that goroutine execution order is not guaranteed

[Concurrency index](../README.md) · [Unordered demo](unordered.go) · [Ordered solution](ordered.go) · [Revision notes](../NOTES.md)

**Level:** 1 — goroutine basics | **Focus:** scheduling versus synchronization.

## Interview task

Launch three goroutines: A prints `A`, B prints `B`, and C prints `C`. Wait for all of them. Predict the output, run it repeatedly, and explain why launching A before B before C does **not** guarantee that print order. Then modify the program to guarantee `A -> B -> C`.

## Reference answer

- [unordered.go](unordered.go) launches the three goroutines in A/B/C order and joins them with `sync.WaitGroup`. Their **completion and print order are unspecified**.
- [ordered.go](ordered.go) passes completion signals through channels: B waits for `aDone`; C waits for `bDone`. These signals establish the necessary order.

Run each program **as an explicit file**, because both files are separate `package main` examples:

```sh
go run unordered.go
go run -race unordered.go
go run ordered.go
go run -race ordered.go
```

Repeat the first command several times. You **might** observe `C B A`, `B A C`, or `A B C`. Seeing the same order repeatedly does not establish a scheduling guarantee. A finite set of runs illustrates possible schedules; the lack of ordering follows from the absence of synchronization between the three print operations.

### Why the channel solution works

```go
go func() { defer wg.Done(); fmt.Println("A"); close(aDone) }()
go func() { defer wg.Done(); <-aDone; fmt.Println("B"); close(bDone) }()
go func() { defer wg.Done(); <-bDone; fmt.Println("C") }()
wg.Wait()
```

The receive from a closed channel proceeds only after the corresponding `close`. The output order becomes A then B then C. **A plain mutex only makes each critical section exclusive; it does not determine which goroutine gets the lock first.** To guarantee order with a mutex, you would also need a condition/turn protocol.

## Scenario questions and answers

| Follow-up | Answer |
|---|---|
| Does `WaitGroup` print goroutines in launch order? | No. It guarantees that registered work has signaled completion before `Wait` returns. |
| Does adding `time.Sleep` in B enforce order? | No. Delays are not a synchronization contract and can be flaky. |
| What if one of the ordered goroutines forgets to close its signal? | The successor can block indefinitely, potentially deadlocking the program. |
| Can results be computed concurrently but printed in order? | Yes. Store each result at a unique index, wait for all workers, then print sequentially; ensure no conflicting writes. |
| Does `GOMAXPROCS(1)` force A/B/C? | No. Limiting execution parallelism does not make scheduling order an API guarantee. |

## Cold recall

1. Distinguish *concurrency*, *parallelism*, *launch order*, *execution order*, and *completion order*.
2. Name the synchronization edge that guarantees B prints after A.
3. Rewrite the ordered solution using a single coordinator rather than chaining workers.

**Evidence:** Record observed orders separately from the formal ordering guarantee in [NOTES.md](../NOTES.md). The sample implementation is not proof you can reproduce it unaided.

References: [Go specification — go statements](https://go.dev/ref/spec#Go_statements), [Go memory model](https://go.dev/ref/mem).
