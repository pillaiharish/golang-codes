# Q15. Use `default` in `select` for nonblocking operations

[Concurrency index](../README.md) · [Q6: buffering](../6-buffered-channel-capacity/README.md) · [Revision notes](../NOTES.md)

**Level 3 — select, timeouts, and cancellation.**

## Task

Attempt a send only if the channel is ready to accept it; otherwise skip the value. Then implement a nonblocking receive.

## Reference answer

```go
select {
case ch <- value:
    fmt.Println("accepted")
default:
    fmt.Println("dropped: not ready")
}

select {
case v, ok := <-ch:
    if ok { fmt.Println("received", v) }
default:
    fmt.Println("nothing available")
}
```

`default` runs if no communication case can proceed immediately. For an unbuffered channel, a send needs a ready receiver; for a buffered channel, it needs space. A nonblocking, deliberately lossy queue should measure drops and define whether dropping is acceptable.

## Scenario Q&A

| Question | Answer |
|---|---|
| Does default guarantee no blocking anywhere in the branch? | No. It only makes the `select` communication decision nonblocking; other code can block. |
| What if you spin in a tight loop around this `select`? | High CPU usage and potential starvation; use blocking waits, timers, or bounded pacing. |
| What if the channel is closed during a send? | Sending to a closed channel panics; a `default` branch does not make send/close races safe. |
| When is dropping acceptable? | For explicitly lossy telemetry or hints, with counters and documented data-loss semantics; not automatically for business-critical writes. |

**Recall:** Explain how nonblocking enqueue differs from delivery acknowledgment and how you would diagnose queue-full drops in a Kafka telemetry pipeline.

Reference: [Go specification — select](https://go.dev/ref/spec#Select_statements).
