# Q4. What happens if a goroutine panics?

[Concurrency index](../README.md) · [Unrecovered panic](panic.go) · [Recovered panic](recover.go) · [Revision notes](../NOTES.md)

**Level:** 1 — goroutine basics | **Focus:** panic scope, deferred functions, recovery boundary, worker failures.

## Interview task

1. Start a goroutine that calls `panic("something went wrong")`; observe process exit and stack trace.
2. In a **separate run**, install a deferred `recover` inside that goroutine.
3. Make `main` wait for the worker, and explain why recovering in `main` cannot catch the worker's panic.

## Reference answers

Run the intentionally failing program and expect a **nonzero exit status**:

```sh
go run panic.go
```

It crashes the **process**, not just that worker. Any deferred cleanup in the panicking goroutine runs while its stack unwinds, but without recovery the panic terminates the program.

Run the safe comparison separately:

```sh
go run recover.go
go run -race recover.go
```

Expected recovered version:

```text
recovered: something went wrong
main completed
```

The essential structure is:

```go
go func() {
    defer wg.Done()
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    panic("something went wrong")
}()
wg.Wait()
```

`recover` must be called **directly by a deferred function executing on the panicking goroutine**. A defer in `main` runs on a different goroutine and cannot recover the worker's panic. After successful recovery, that panicking function does not resume where it panicked; it returns after its deferred functions finish.

## Scenario questions and answers

| Follow-up | Answer |
|---|---|
| Put `defer recover` in `main`; will it catch the worker? | No. Recovery is goroutine-local. |
| Call `recover()` in ordinary code (not a defer); does it stop a panic? | No; outside the required deferred context, it has no recovery effect. |
| Does `defer wg.Done()` prevent a crash? | No. It accounts for task completion during unwinding, but does not recover. |
| Can a worker pool wrap each job in one recovery boundary? | Yes, if it converts the panic to a reported failed job and isolates the recovery to the job. Do not silently drop a job. |
| Should every panic be ignored to keep the service alive? | No. Recovery should log/report with context and stack evidence, maintain state invariants, and apply an explicit fail/restart policy. |
| Will `recover` catch a fatal runtime error or external process kill? | No. It is not a universal process crash guard. |

## Cold recall

- Why must Q4 use two separate programs or runs?
- What happens to subsequent statements inside a function after a panic is recovered?
- Where would you put a recovery boundary in a long-running Kafka worker, and how would you account for the failed record?
- How is a recoverable panic different from an ordinary returned `error`?

**Evidence:** Record the failing exit status, recovered output, and how you would surface the worker failure in [NOTES.md](../NOTES.md). Do not run the deliberately failing example inside a shared long-running service.

References: [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover), [Go specification — handling panics](https://go.dev/ref/spec#Handling_panics).
