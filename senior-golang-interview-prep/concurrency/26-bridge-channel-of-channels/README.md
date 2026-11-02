# Q26. Implement a `bridge` for a channel of channels

[Concurrency index](../README.md) · [Q22: fan-in/fan-out](../22-fan-out-fan-in/README.md) · [Revision notes](../NOTES.md)

**Level 5 — core concurrency patterns.**

## Task

Flatten `streams <-chan <-chan int` into one `<-chan int`. If the outer channel yields streams `[1,2,3]`, `[4,5]`, `[6]`, output should be `1,2,3,4,5,6` for a **sequential bridge**.

## Reference approach

```go
func bridge(ctx context.Context, streams <-chan <-chan int) <-chan int
```

Implement one forwarding goroutine:

1. Receive the next inner stream from `streams` with a cancellation-aware `select`.
2. Drain that **one** inner stream in order. Every receive and send must also select on `ctx.Done()`.
3. Once the inner stream closes, move to the next inner stream.
4. Close the output when the outer stream closes and all received inner streams have been drained, or on cancellation.

The sequential bridge preserves inner-stream order and outer-stream sequence. A concurrently merged bridge is a **different contract**: it can interleave values and needs a WaitGroup to close its output safely.

## Scenario Q&A

| Question | Answer |
|---|---|
| What if an inner channel never closes? | Sequential forwarding stalls on that stream; cancellation must provide an escape. |
| What if the outer channel produces a nil inner channel? | A receive from nil blocks forever unless you skip it or select on cancellation. |
| What if the consumer stops reading output? | The forwarding send can block; use a cancellation-aware send. |
| Who closes the flattened output? | The bridge's one owner, after its forwarding goroutine exits. |

**Recall:** State whether your bridge guarantees global order, and compare its throughput and head-of-line blocking to fan-in.

Reference: [Go blog — pipelines](https://go.dev/blog/pipelines).
