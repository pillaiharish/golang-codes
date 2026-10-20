# Task 10: Rebalance-safe consumer coordinator

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Extend the progress tracker with assignment epochs. Assign starts a partition generation; revoke stops new admissions and drains or cancels within a bounded policy. Fence stale work and commit callbacks by partition plus generation. Do not assume cancellation reverses an external effect. Use a fake group coordinator first, then exercise actual rebalance behavior with a pinned Go client and classic consumer group.

## Acceptance tests and failure injection

1. Revoke while a handler is blocked, reassign the partition, then release the old handler; its completion cannot advance the new generation.
2. A delayed commit response from the old generation cannot overwrite current state.
3. Repeated assign/revoke cycles release resources and preserve bounds.
4. With a real broker, restart one consumer during processing and reconcile replayed IDs and effects; record configuration and observed behavior.

## Evidence to collect

Document poll-loop, callback and drain ownership. Record how the selected client permits safe commits during revocation; do not invent an API-independent sequence.

## Senior follow-up

How would external fencing or transactional deduplication prevent duplicate effects that local epoch checks cannot stop?

## Prerequisites

- [Context, deadlines, and errors](../../topics/05-context-deadlines-and-errors.md)
- [Kafka consumers, offsets, and rebalances](../../topics/09-kafka-consumer-offsets-and-rebalances.md)
- [Delivery semantics, idempotency, and retries](../../topics/10-delivery-idempotency-and-retry.md)
- [Graceful shutdown and senior design review](../../topics/12-shutdown-and-design-review.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
