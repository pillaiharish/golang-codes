# Task 02: Bounded producer-consumer queue

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Design Put(ctx, item), TryPut(item), Get(ctx), and Close(). Specify errors for full, closed, canceled and oversize. Bound queued item count and retained payload bytes; state separately how caller-owned and active payloads are accounted. Copy payloads on acceptance or document exclusive ownership. Close rejects subsequent admissions and lets accepted queued items drain; blocked waiters must wake. Define the linearization point for admission versus Close. A canceled operation that races with successful admission may succeed if your documented contract allows it.

## Acceptance tests and failure injection

1. Fill exactly the count limit, then prove TryPut rejects without waiting for a consumer; separately hit the byte limit with fewer records.
2. Block Put and Get, then cancel or close; every waiter returns with the documented outcome.
3. Race many producers with Close: no panic, no double acceptance, no lost accepted item during a successful drain.
4. Reject an individually oversized record immediately instead of waiting forever.

## Evidence to collect

Expose queued records/bytes and rejection reasons. Explain why queue bounds do not bound unlimited blocked callers.

## Senior follow-up

Add a deadline-based blocking API while preserving the try API. Would a drop-oldest policy change ordering guarantees?

## Prerequisites

- [Memory ordering and data ownership](../../topics/02-memory-model-and-ownership.md)
- [Channels, select, and queue contracts](../../topics/04-channels-select-and-queue-contracts.md)
- [Backpressure, rate limits, and capacity](../../topics/07-backpressure-rate-limits-and-capacity.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
