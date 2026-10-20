# Task 13: Race, deadlock, and leak diagnosis lab

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Construct isolated test fixtures for five symptoms: shared map mutation, opposing lock order, producer blocked after early consumer exit, retained slice backing arrays, and a retry loop that ignores cancellation. Keep deliberately broken fixtures opt-in and separate from normal passing tests. Diagnose each using the appropriate race report, stack dump, profile or trace, then repair it and add a regression test.

## Acceptance tests and failure injection

1. Each failure is reproduced through barriers or explicit triggers, not a lucky sleep.
2. A repair removes the demonstrated failure and preserves functional behavior under cancellation and shutdown.
3. Regression tests assert invariants or owned-worker completion, not an exact process-wide goroutine count.
4. Capture one example where the race detector is clean but a logical race or deadlock remains.

## Evidence to collect

For each incident, save hypothesis, evidence, root cause, minimal repair and remaining uncertainty. Bound failing runs so a deliberate deadlock does not hang the suite.

## Senior follow-up

How would you mitigate the issue in production before a permanent fix is deployed?

## Prerequisites

- [Goroutines, scheduling, and lifetime](../../topics/01-goroutines-scheduler-and-lifetime.md)
- [Testing, profiling, and observability](../../topics/11-testing-profiling-and-observability.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
