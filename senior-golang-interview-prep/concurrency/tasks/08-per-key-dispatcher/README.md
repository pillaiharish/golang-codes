# Task 08: Per-key ordered dispatcher

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Schedule records across bounded workers while preserving serial handler execution for each key. Define submission order for concurrent calls, active-key limits, idle-key eviction and retry ordering. Independent keys should progress concurrently. Bound per-key and global backlog. If retrying a failed record, do not let later same-key effects overtake it unless the contract explicitly permits that behavior.

## Acceptance tests and failure injection

1. Hold key A at a barrier; key B progresses while a second A operation does not execute concurrently.
2. A hot key reaches its bound without exhausting unlimited memory or violating the global limit.
3. Retry, cancellation and key eviction preserve sequencing and do not allow two owners for one key.
4. Quiescence releases per-key bookkeeping; shutdown joins all lanes.

## Evidence to collect

Report skewed-key latency, fairness and memory. Compare hash-sharded lanes with explicit per-key scheduling and explain collision trade-offs.

## Senior follow-up

How would you preserve ordering when keys move between processes or Kafka partitions?

## Prerequisites

- [Worker pools, fan-out, and ordering](../../topics/06-worker-pools-and-patterns.md)
- [Kafka producer concurrency and batching](../../topics/08-kafka-producer-concurrency.md)
- [Kafka consumers, offsets, and rebalances](../../topics/09-kafka-consumer-offsets-and-rebalances.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
