# Task 06: Rate, concurrency, and tenant admission

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Implement admission with a global active-work cap and per-tenant rate/burst limits. Choose rejection or bounded waiting; never silently allocate an unbounded waiter queue. Return a release handle that cannot leak capacity on handler failure, and define duplicate release behavior. Bound tenant metadata and define eviction without giving active tenants an easy quota reset. Use an injectable clock where feasible.

## Acceptance tests and failure injection

1. One tenant exhausts its burst without blocking other tenants that remain eligible under the global budget.
2. Slow work saturates concurrency even when rate is low; admission and rejection metrics distinguish the causes.
3. Cancellation while waiting does not consume a permit forever; release after errors restores capacity exactly once.
4. Adversarial tenant IDs do not create unlimited metadata or unbounded metric cardinality.

## Evidence to collect

Plot or tabulate offered, admitted and completed work under burst and sustained overload. Explain rate, concurrency and queue limits separately.

## Senior follow-up

What additional guarantees are needed when limits must hold across several service replicas?

## Prerequisites

- [Backpressure, rate limits, and capacity](../../topics/07-backpressure-rate-limits-and-capacity.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
