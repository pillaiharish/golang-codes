# 05. Context, deadlines, and errors

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

Context propagates cancellation and deadlines; it does not forcibly stop arbitrary code. Every blocking dependency must cooperate or have a separate timeout/isolation strategy. Cancellation signals intent, whereas joining proves workers have exited. Always release resources associated with a derived context using its cancel function.

A request context normally ends with the request. Accepted background work needs an explicitly owned service lifetime plus its own bounded execution budget. Replacing every context with Background hides cancellation; retaining request contexts indefinitely can retain request-scoped data. Pass contexts as operation arguments and avoid using values as optional configuration.

Define total time as admission wait + processing + retries + backoff + cleanup. Each retry consumes the same remaining budget rather than starting a fresh end-to-end timeout. With errgroup.WithContext, the derived context is canceled on the first non-nil worker error or when Wait returns. Do not reuse it for post-Wait work. SetLimit limits active group functions, but Go may block on admission; it does not independently bound every upstream producer.

Fail-fast, collect-all and first-success are different error policies. The first error is not necessarily the root cause; preserve useful error context without masking cancellation or leaking sensitive payloads.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. Does cancel wait for completion?

**Medium. Strong answer:** No. Join separately and define what happens if dependencies ignore cancellation.

### Q2. How do you bound a fan-out request to 200 ms?

**Hard. Strong answer:** Allocate one overall deadline including admission and retries; propagate it and collect all owned workers.

### Q3. Can first-success use a first-error policy unchanged?

**Hard. Strong answer:** No. One failure must not cancel potential winners; cancel only after success or exhaustion according to the contract.

### Q4. Why did work after errgroup.Wait fail immediately?

**Hard. Strong answer:** The group-derived context is already canceled; use the appropriate parent or separately scoped context.

### Q5. What does a timeout mean for a remote write?

**Expert. Strong answer:** Its outcome may be unknown. Cancellation does not undo a committed remote effect; retry with an idempotency strategy.

## Implement and demonstrate

- [Task 03: Bounded worker pool](../tasks/03-bounded-worker-pool/README.md)
- [Task 04: Cancellable fan-out and first-success](../tasks/04-cancellable-fanout/README.md)
- [Task 10: Rebalance-safe consumer coordinator](../tasks/10-rebalance-coordinator/README.md)
- [Task 12: Graceful service shutdown](../tasks/12-graceful-shutdown/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://pkg.go.dev/context](https://pkg.go.dev/context)
- [https://pkg.go.dev/golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
