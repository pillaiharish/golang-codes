# Q29. Detect and fix a goroutine leak — worked solution

[Back to Q29 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q23: pipeline](../../23-cancellable-pipeline/README.md) · [Revision notes](../../NOTES.md)

**Level 6 — production correctness.**

## Task

Deliberately start a goroutine blocked on an unreceived channel operation. Observe goroutine counts and stack dumps; then add cancellation and prove the worker exits.

## Reference approach

```go
ch := make(chan int)
go func() { ch <- 1 }() // blocks until somebody receives
fmt.Println("goroutines:", runtime.NumGoroutine())
```

This is a **leak candidate** if the intended receiver never arrives. A single `NumGoroutine` observation is not proof: counts change naturally. Reproduce under repeated operations, compare a baseline, inspect goroutine stacks/profiles and verify eventual return to a stable count. Fix the blocked send with cancellation-aware `select`, or guarantee a receiver/drain path.

## Scenario Q&A

| Question | Answer |
|---|---|
| Does `WaitGroup` prevent leaks? | No. It can wait forever if workers never reach `Done`. |
| Does closing a context force a blocked syscall to finish? | Not generically. The operation must support cancellation or have a deadline/unblock mechanism. |
| Can a leaked goroutine keep data in memory? | Yes. Its stack and referenced heap objects remain reachable. |
| How do you test leak freedom? | Deterministically stop work, join workers, and observe profiles/counts over repeated runs; tools such as `goleak` are useful additional checks. |

**Recall:** Map a Kafka telemetry producer stuck on a full queue: identify ownership, escape path, drop policy and the metric that reveals the leak.

References: [Go blog — pipelines](https://go.dev/blog/pipelines), [goroutine profiles](https://pkg.go.dev/runtime/pprof).
