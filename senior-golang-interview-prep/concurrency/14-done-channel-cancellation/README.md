# Q14. Implement a `done` channel pattern

[Concurrency index](../README.md) · [Q13: context](../13-context-cancellation/README.md) · [Revision notes](../NOTES.md)

**Level 3 — select, timeouts, and cancellation.**

## Task

Use `done := make(chan struct{})` to broadcast a stop signal to several workers without `context.Context`.

## Reference answer

```go
done := make(chan struct{})
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        ticker := time.NewTicker(100 * time.Millisecond)
        defer ticker.Stop()
        for {
            select {
            case <-done: return
            case <-ticker.C: /* one short unit of work */
            }
        }
    }()
}
close(done) // owner broadcasts to all receivers
wg.Wait()   // then joins all workers
```

Closing `done` makes every current and future receive from it ready. A `chan struct{}` carries only a signal. The ticker here prevents a busy loop; if the worker instead blocks on I/O or a channel operation, use `select { case <-done: ...; case ...: ... }` around that blocking operation.

## Scenario Q&A

| Question | Answer |
|---|---|
| Why close rather than send one value? | One close broadcasts; one sent value is consumed by one receiver. |
| What if two goroutines call `close(done)`? | A second close panics. Assign a single owner or synchronize closure. |
| Does close mean all workers finished? | No. Use `WaitGroup` or explicit acknowledgements to join. |
| Does a `done` channel carry errors or deadlines? | Not by itself; `context` adds standard propagation and deadline APIs. |

**Recall:** Design cancellation for a worker waiting on input and a worker blocked sending output; ensure both can exit after `done` closes.

Reference: [Go blog — pipelines and cancellation](https://go.dev/blog/pipelines).
