# Q13. Implement cancellation using `context.Context` — worked solution

[Back to Q13 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q14: done channel](../../14-done-channel-cancellation/README.md) · [Revision notes](../../NOTES.md)

**Level 3 — select, timeouts, and cancellation.**

## Task

Start a long-running worker, cancel it from `main` after about one second, and wait until it exits. Use `context.WithCancel` or `context.WithTimeout`.

## Reference answer

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            return
        case <-time.After(100 * time.Millisecond):
            fmt.Println("working")
        }
    }
}()
wg.Wait()
fmt.Println("worker stopped:", ctx.Err())
```

The worker checks `ctx.Done()` and cooperatively returns. The deadline makes cancellation available; `WaitGroup` separately proves that the worker actually exited. Cancellation is a request, not a forced goroutine kill.

## Scenario Q&A

| Question | Answer |
|---|---|
| If the worker ignores `ctx.Done()`? | `Wait` may block forever; context cannot forcibly terminate it. |
| What if a worker is blocked sending to a channel? | Put the send and `<-ctx.Done()` into a `select` so it can exit. |
| Why call `cancel()` even with a deadline? | Releases associated timer/parent resources promptly. |
| Can two child contexts share one parent deadline? | Yes. Parent cancellation propagates to descendants. |
| How should a library treat context? | Pass it as the first parameter to request-scoped operations, propagate it, and do not store it as arbitrary long-lived object state. |

**Recall:** Explain how cancellation, joining, and deadlines differ. Replace the timer-driven worker with a receiver on an input channel that honors cancellation.

Reference: [`context` package](https://pkg.go.dev/context).
