# Q24. Implement `orDone` — worked solution

[Back to Q24 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q23: pipeline](../../23-cancellable-pipeline/README.md) · [Revision notes](../../NOTES.md)

**Level 5 — core concurrency patterns.**

## Task

Implement `func orDone(ctx context.Context, ch <-chan int) <-chan int`: relay values, exit if the source closes or context cancels, and close the returned channel exactly once.

## Reference answer

```go
func orDone(ctx context.Context, ch <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for {
            select {
            case <-ctx.Done(): return
            case v, ok := <-ch:
                if !ok { return }
                select {
                case out <- v:
                case <-ctx.Done(): return
                }
            }
        }
    }()
    return out
}
```

Both input receive and output send need cancellation paths; otherwise a downstream consumer that stops early can strand this helper. The returned channel's owner is the relay goroutine.

## Scenario Q&A

| Question | Answer |
|---|---|
| What if `ctx` is canceled before the first value? | Output closes without any values. |
| What if `ch` closes normally? | Relay exits and closes its output after forwarding received values. |
| What if both a value and cancellation are ready? | `select` may pick either; cancellation is not an exact cutoff unless a stricter protocol is added. |
| Does this function shut down the original source goroutine? | Not by itself. The source must also observe cancellation. |

**Recall:** Explain what ownership the relay has, what it cannot promise about its upstream producer, and how you would test cancellation while blocked on a send.

Reference: [Go blog — pipelines](https://go.dev/blog/pipelines).
