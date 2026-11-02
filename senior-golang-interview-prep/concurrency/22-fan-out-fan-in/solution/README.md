# Q22. Implement fan-out and fan-in — worked solution

[Back to Q22 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q21: worker pool](../../21-worker-pool/README.md) · [Revision notes](../../NOTES.md)

**Level 5 — core concurrency patterns.**

## Task

Fan-out: multiple workers read the **same** input channel. Fan-in: merge values from several input channels into one output channel. Write `merge(channels ...<-chan int) <-chan int`.

## Reference answer

```go
func merge(inputs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    wg.Add(len(inputs))
    for _, input := range inputs {
        go func(ch <-chan int) {
            defer wg.Done()
            for v := range ch { out <- v }
        }(input)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}
```

The fan-in owner closes `out` **after** every forwarder finishes; no worker closes a shared output channel. The function preserves per-input ordering as received, but not a global order between the inputs.

## Scenario Q&A

| Question | Answer |
|---|---|
| What if no inputs are given? | The WaitGroup count is zero; the coordinator closes `out`. |
| What if one input never closes? | Its forwarder may never finish, so `out` cannot close. |
| What if the downstream consumer stops early? | Forwarders can block on `out <- v`. Add cancellation-aware sends/receives. |
| Does fan-out make each input item reach every worker? | No. Workers competing on one channel distribute items; broadcast requires a different design. |

**Recall:** Add `context.Context` to `merge`, and explain how to stop all forwarders without waiting for stalled input sources.

Reference: [Go blog — pipelines and cancellation](https://go.dev/blog/pipelines).
