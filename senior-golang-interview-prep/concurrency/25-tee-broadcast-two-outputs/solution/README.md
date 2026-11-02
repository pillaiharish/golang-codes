# Q25. Implement `tee` with two outputs — worked solution

[Back to Q25 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q24: or-done](../../24-or-done/README.md) · [Revision notes](../../NOTES.md)

**Level 5 — core concurrency patterns.**

## Task

Write a `tee` function that accepts one input channel and returns two channels. **Every input value must reach both outputs**, then both outputs must close when the source ends.

## Reference approach

1. Start one distributor goroutine that exclusively receives input and owns closing both outputs.
2. For each received value, send it once to A and once to B. Track whether each send succeeded; do not send the same value twice to an already-served output.
3. Use `select` with the shared context during both sends so an abandoned receiver cannot strand the distributor.
4. When input closes or context cancels, close both outputs once. The upstream producer separately needs a cancellation path.

Conceptual core (inside the distributor after receiving `v`):

```go
a, b := outA, outB
for i := 0; i < 2; i++ {
    select {
    case a <- v: a = nil
    case b <- v: b = nil
    case <-ctx.Done(): return
    }
}
```

A nil channel disables its case after successful delivery. The outer input receive must also select on `ctx.Done()`; on return, defer closing both owned outputs.

## Scenario Q&A

| Question | Answer |
|---|---|
| Is tee equivalent to two goroutines reading the same input? | No. Competing consumers divide messages; they do not each receive every message. |
| What if one output has no receiver? | Without cancellation or a drop policy, the other output may be held up by the slow branch. |
| Are the two outputs guaranteed to receive a value simultaneously? | No; each send is a separate operation. |
| Can a tee silently drop one branch? | Only if the contract explicitly permits lossy delivery; then count/report drops. |

**Recall:** Define what happens on slow subscribers, cancellation and partial delivery of a value to A but not B.

Reference: [Go blog — pipelines](https://go.dev/blog/pipelines).
