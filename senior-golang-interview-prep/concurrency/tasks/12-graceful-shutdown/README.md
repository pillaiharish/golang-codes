# Task 12: Graceful service shutdown

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Build a service lifecycle coordinator around admission, workers and a publisher. Expose a Shutdown(ctx) operation and wire a local demo to SIGTERM/SIGINT when you implement it. Transition accepting → draining → stopped with synchronized admission. Use a shared shutdown budget and reserve time for flush and joins. Specify repeated shutdown calls and callers with different deadlines. Separate cooperative completion from unresolved work.

## Acceptance tests and failure injection

1. Race new submissions with shutdown; each is either rejected or included in accepted-work accounting.
2. Stop with queued work, active work and pending acknowledgements; verify dependency teardown order.
3. Expire the drain budget and report unresolved work; never claim an uncooperative operation has stopped.
4. Repeated shutdown and cancellation do not double-close channels, lose permits or panic.

## Evidence to collect

Use deterministic tests for lifecycle behavior, then a local signal test. Document readiness propagation and Kubernetes grace-budget assumptions without adding deployment manifests.

## Senior follow-up

How do long-lived HTTP connections and background workers affect the shutdown proof?

## Prerequisites

- [Channels, select, and queue contracts](../../topics/04-channels-select-and-queue-contracts.md)
- [Context, deadlines, and errors](../../topics/05-context-deadlines-and-errors.md)
- [Kafka producer concurrency and batching](../../topics/08-kafka-producer-concurrency.md)
- [Graceful shutdown and senior design review](../../topics/12-shutdown-and-design-review.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
