# Q23. Build a cancellable pipeline

[Concurrency index](../README.md) · [Q24: or-done](../24-or-done/README.md) · [Revision notes](../NOTES.md)

**Level 5 — core concurrency patterns.**

## Task

Build `generator -> square -> printer` with independent goroutines and channels. Then add cancellation so stopping the final consumer does not leak upstream producers.

## Reference answer

An example cancellation-aware `square` stage:

```go
func square(ctx context.Context, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for {
            select {
            case <-ctx.Done(): return
            case n, ok := <-in:
                if !ok { return }
                select {
                case out <- n*n:
                case <-ctx.Done(): return
                }
            }
        }
    }()
    return out
}
```

The generator should similarly `select` on cancellation when sending, and close only **its own** output channel. The printer should cancel the shared context if it exits early. Each stage owns closing its output; no receiver should race to close another stage's channel.

## Scenario Q&A

| Question | Answer |
|---|---|
| Why are both receive and send cancelable? | Either may block indefinitely if an adjacent stage stops. |
| Does a closed output tell upstream to stop? | Not safely. Sending to a closed output can panic; cancellation is a separate signal. |
| What happens when the printer consumes only one value? | Without cancellation, upstream sends may block, leaking goroutines. |
| How does a slow printer affect generator throughput? | Backpressure propagates through channel sends; buffers change burst capacity, not underlying throughput. |

**Recall:** Specify an ownership/close policy and prove every launched stage has an exit path on cancellation.

Reference: [Go blog — pipelines](https://go.dev/blog/pipelines).
