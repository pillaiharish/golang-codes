# Q20. Limit concurrency using a semaphore — worked solution

[Back to Q20 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q27: rate limiting](../../27-rate-limiter/README.md) · [Revision notes](../../NOTES.md)

**Level 4 — shared state and synchronization.**

## Task

Process many independent tasks while allowing at most **three simultaneously active tasks**. Use a buffered channel as a counting semaphore.

## Reference answer

```go
sem := make(chan struct{}, 3)
var wg sync.WaitGroup
for _, task := range tasks {
    sem <- struct{}{} // acquire BEFORE starting the goroutine
    wg.Add(1)
    go func(t Task) {
        defer wg.Done()
        defer func() { <-sem }() // release even on early return
        process(t)
    }(task)
}
wg.Wait()
```

Acquiring before `go` limits **active plus launched work** rather than creating an unbounded number of goroutines blocked inside semaphore acquisition. In an error/cancellation-aware version, acquire with a `select` that also watches `ctx.Done()`; ensure rejected tasks are not added to the WaitGroup.

## Scenario Q&A

| Question | Answer |
|---|---|
| Is a semaphore a rate limiter? | No. It bounds simultaneous work, not starts per second. |
| What if release is forgotten? | The producer can block permanently once permits are exhausted. |
| Is acquiring inside each goroutine equivalent? | It can bound `process`, but still creates an unbounded backlog of waiting goroutines. |
| What if cancellation occurs while waiting for a permit? | Use a cancelable admission `select`; document what happens to rejected work. |
| Could a fixed worker pool be simpler? | Yes, especially for long-lived jobs and bounded queues. |

**Recall:** Given 5 requests/second but 30-second jobs, explain why limiting rate alone still permits concurrency to grow.

References: [Go blog — pipelines](https://go.dev/blog/pipelines), [`x/sync/semaphore`](https://pkg.go.dev/golang.org/x/sync/semaphore).
