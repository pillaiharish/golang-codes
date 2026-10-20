# Task 03: Bounded worker pool

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Implement a pool with W workers, Q queued jobs, Submit(ctx, job), and Shutdown(ctx). Bound admission before spawning additional work. Choose fail-fast or continue-on-job-error and document it. Every accepted job gets exactly one terminal local result: completed, failed, canceled-before-start, or unresolved at forced shutdown. The handler accepts context; a separate adversarial test may use a handler that ignores it to demonstrate the limit of cancellation.

## Acceptance tests and failure injection

1. A barrier-controlled handler records a maximum of W active calls; Q additional queued jobs fit and the next follows admission policy.
2. Cancel before admission, while queued and during processing; preserve the chosen accounting and release permits on errors.
3. Shutdown during active submissions is race-free and joins cooperative handlers; repeated Shutdown does not panic.
4. An uncooperative handler cannot be falsely reported as stopped just because Shutdown timed out.

## Evidence to collect

Provide a state diagram and a table of accepted versus terminal outcomes after quiescence. Benchmark completed jobs, not just submissions.

## Senior follow-up

How would priorities or tenant fairness change queue design and starvation risks?

## Prerequisites

- [Goroutines, scheduling, and lifetime](../../topics/01-goroutines-scheduler-and-lifetime.md)
- [Mutexes, atomics, and coordination](../../topics/03-mutexes-atomics-and-coordination.md)
- [Channels, select, and queue contracts](../../topics/04-channels-select-and-queue-contracts.md)
- [Context, deadlines, and errors](../../topics/05-context-deadlines-and-errors.md)
- [Worker pools, fan-out, and ordering](../../topics/06-worker-pools-and-patterns.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
