# Q21. Implement a worker pool

[Concurrency index](../README.md) · [Q22: fan-out/fan-in](../22-fan-out-fan-in/README.md) · [Advanced worker-pool task](../tasks/03-bounded-worker-pool/README.md)

**Level 5 — core concurrency patterns.**

## Task

Create **3 workers**, **10 jobs**, and a results channel. Every worker consumes jobs and returns `job * 2`; collect all ten results without leaking goroutines or closing the results channel early.

## Reference answer

```go
jobs := make(chan int)
results := make(chan int)
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for j := range jobs { results <- j * 2 }
    }()
}
go func() {
    for j := 1; j <= 10; j++ { jobs <- j }
    close(jobs) // producer owns jobs
}()
go func() {
    wg.Wait()
    close(results) // coordinator owns results
}()
for result := range results { fmt.Println(result) }
```

Do not wait for workers **before** receiving from an unbuffered results channel: the workers could block on result sends and never reach `Done`. `results` closes only after every worker exits.

## Scenario Q&A

| Question | Answer |
|---|---|
| Is result order guaranteed? | No. Worker completion order varies; attach job IDs to reorder if needed. |
| Who closes `results`? | One coordinator after all workers finish, not individual workers. |
| What if the consumer stops reading early? | Workers can block on sends; add cancellation or a drain policy. |
| How do you bound memory? | Bound worker count **and** admission queue, plus waiting submitters and result buffering. |
| What about job failures? | Define per-job outcomes, error channels/structs, and whether failure cancels the pool. |

**Recall:** Defend an admission policy, a shutdown protocol and one deterministic test showing that max active workers never exceeds 3.

References: [Go blog — pipelines](https://go.dev/blog/pipelines), [advanced task 03](../tasks/03-bounded-worker-pool/README.md).
