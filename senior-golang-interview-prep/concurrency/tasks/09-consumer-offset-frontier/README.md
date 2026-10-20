# Task 09: Parallel consumer completion frontier

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Implement a pure per-partition progress tracker before attaching it to Kafka. Feed ordered delivered offsets and completion events; return the next safe commit position. Offsets may have gaps. Track only delivered work, bound outstanding entries and reject duplicate/stale completion changes. A completion means the effect reached your stated durable boundary. Keep each partition independent.

## Acceptance tests and failure injection

1. Deliver 40, 41, 42; complete 42 first and prove the tracker cannot skip 40. Complete 40 then 41 and verify safe advancement.
2. Deliver 40 then 43; finishing both must not wait for nonexistent offsets 41 and 42.
3. A failed earliest record blocks advancement unless the configured durable quarantine policy completes.
4. Interleave two partitions, duplicate completions and restart from a committed position; never regress or cross unfinished work.

## Evidence to collect

Write a table-driven state-transition suite and a model/property test that generated completion orders never commit past unfinished delivered work.

## Senior follow-up

What changes when commits themselves fail or complete out of order?

## Prerequisites

- [Kafka consumers, offsets, and rebalances](../../topics/09-kafka-consumer-offsets-and-rebalances.md)
- [Delivery semantics, idempotency, and retries](../../topics/10-delivery-idempotency-and-retry.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
