# Task 01: Concurrent state and immutable snapshots

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Build a CounterStore with Add(key, delta), Get(key), and Snapshot(). Start from Module 8's instance-state model. All operations must be safe concurrently. Define whether a snapshot represents one consistent point in time; require that guarantee for this task. Returned snapshots must not expose mutable internal storage. Reject or explicitly handle integer overflow in your chosen domain. Do not copy a live store containing synchronization primitives.

## Acceptance tests and failure injection

1. Run concurrent updates to one key and many keys; final totals equal the known sum after joining all workers.
2. Mutating a returned snapshot cannot change store state; a snapshot cannot combine fields from different logical transactions if your API updates multiple fields.
3. Force reads during writes with barriers and run the race detector; test independent store instances and zero-value behavior.

## Evidence to collect

Compare a simple mutex design with one alternative under read-heavy and write-heavy loads. Record allocations, throughput and contention without assuming RWMutex wins.

## Senior follow-up

How would you publish a large read-mostly configuration snapshot without copying it on every read?

## Prerequisites

- [Memory ordering and data ownership](../../topics/02-memory-model-and-ownership.md)
- [Mutexes, atomics, and coordination](../../topics/03-mutexes-atomics-and-coordination.md)
- [Testing, profiling, and observability](../../topics/11-testing-profiling-and-observability.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
