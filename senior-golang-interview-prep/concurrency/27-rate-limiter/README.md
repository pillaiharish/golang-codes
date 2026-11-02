# Q27. Implement a rate limiter

[Concurrency index](../README.md) · [Q20: semaphore](../20-semaphore-limit-concurrency/README.md) · [Revision notes](../NOTES.md)

**Level 5 — core concurrency patterns.**

## Task

Process 20 tasks at a pace of **at most five starts per second**. Implement a simple ticker-based pacer and compare it with `golang.org/x/time/rate`. Explain why a concurrency bound may also be necessary.

## Reference approach

```go
ticker := time.NewTicker(200 * time.Millisecond)
defer ticker.Stop()
sem := make(chan struct{}, 3)
var wg sync.WaitGroup
for _, task := range tasks {
    <-ticker.C         // pace admissions
    sem <- struct{}{}  // cap active work before go
    wg.Add(1)
    go func(t Task) {
        defer wg.Done()
        defer func() { <-sem }()
        process(t)
    }(task)
}
wg.Wait()
```

This shows a roughly 5-starts/second pacer plus an active-work limit of three. For strict request-rate semantics including bursts and per-tenant limits, choose and document a token-bucket policy (`x/time/rate`) instead of assuming ticker timing is exact.

## Scenario Q&A

| Question | Answer |
|---|---|
| Does a ticker alone cap active goroutines? | No. Work lasting longer than the interval can accumulate; add a semaphore/worker pool. |
| Does a semaphore guarantee five requests per second? | No. It limits concurrency, not start rate. |
| What happens when a downstream service slows? | The semaphore fills and admission blocks; define timeout, buffering and rejected-work behavior. |
| What does a token bucket add? | Configurable average rate and burst capacity. |
| Can you safely use this for per-tenant fairness? | Not without per-tenant accounting and a fairness policy. |

**Recall:** Compare rate, concurrency, queue capacity and drop counters using your asynchronous telemetry/Kafka experience.

References: [`time.Ticker`](https://pkg.go.dev/time#Ticker), [`x/time/rate`](https://pkg.go.dev/golang.org/x/time/rate).
