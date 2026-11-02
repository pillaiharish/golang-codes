# Q30. Implement graceful shutdown

[Concurrency index](../README.md) · [Q29: leaks](../29-detect-goroutine-leak/README.md) · [Advanced shutdown task](../tasks/12-graceful-shutdown/README.md)

**Level 6 — production correctness.**

## Task

Run a service with workers, listen for SIGINT, stop accepting new work, let admitted jobs finish within a bounded grace period, and then release resources. Use `signal.NotifyContext`.

## Reference approach

```go
signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()
// Start workers; admission and work processing have explicit owners.
<-signalCtx.Done()
// 1. Stop accepting new jobs atomically with Submit.
// 2. Signal the producer that no more jobs can be enqueued.
// 3. Drain/join accepted work using a SEPARATE shutdown deadline.
// 4. Close sinks and report work that did not finish.
```

Use a **separate drain context** if jobs must finish after the signal: the signal context is already canceled. A deadline cannot forcibly kill an uncooperative worker; define what the service reports and what the process does when the budget expires.

## Scenario Q&A

| Question | Answer |
|---|---|
| Why stop admission before closing a jobs channel? | Otherwise a concurrent submitter may send to a closed channel and panic or lose accepted work. |
| Is canceling workers immediately always graceful? | No. It may abandon already accepted work; choose drain versus abort contract. |
| What if Kafka has buffered records? | Flush within the deadline, distinguish accepted/enqueued from broker-acknowledged, and report unresolved deliveries. |
| What if shutdown is called twice? | Make state transitions idempotent; only the owner closes each channel/resource once. |
| Does the deadline guarantee the process has no goroutines? | No, if uncooperative work ignores shutdown; the process-level policy must be explicit. |

**Recall:** Draw admission-open -> stopping -> draining -> terminated states. Specify invariants and failure accounting at every transition.

References: [`signal.NotifyContext`](https://pkg.go.dev/os/signal#NotifyContext), [advanced task 12](../tasks/12-graceful-shutdown/README.md).
