# Task 05: Cache loading with shared work

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Build Get(ctx, key) around a bounded cache with TTL and a loader. Coalesce simultaneous loads for the same key; different keys may load concurrently under a global cap. Canceling one waiter must not cancel a load still needed by other waiters. Give shared work a service-owned timeout and specify the last-waiter policy. Cache successful values only for the base task; document negative caching separately. Inject a clock for expiration tests.

## Acceptance tests and failure injection

1. Many callers for one missing key invoke the loader once while it is in flight; unrelated keys are not globally serialized.
2. Cancel one waiter and let another receive success; all waiter registrations and permits are released.
3. Expire an entry using the fake clock; stale completion cannot overwrite a newer generation after invalidation.
4. Bound entry count and load concurrency; test errors, eviction and service shutdown.

## Evidence to collect

Report cache hits, coalesced waiters and loader calls separately. Document value ownership to prevent caller mutation.

## Senior follow-up

What changes for stale-while-revalidate, and why is this not durable event deduplication?

## Prerequisites

- [Mutexes, atomics, and coordination](../../topics/03-mutexes-atomics-and-coordination.md)
- [Worker pools, fan-out, and ordering](../../topics/06-worker-pools-and-patterns.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
