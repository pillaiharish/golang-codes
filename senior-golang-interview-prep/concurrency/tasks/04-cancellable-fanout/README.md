# Task 04: Cancellable fan-out and first-success

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Build two separately specified operations over context-aware backends: FetchAll collects results with a maximum parallelism; FirstSuccess returns a successful result, canceling losing work. In FirstSuccess, one backend failure must not cancel other possible winners. Define empty input, all-failed behavior and partial results. No worker may remain blocked sending a result after its caller returns; choose a bounded result channel or explicit joining protocol.

## Acceptance tests and failure injection

1. Control backend completion order; a fast error followed by a slower success returns success.
2. All failures return useful aggregated information; cancellation terminates all cooperative backends.
3. Race two successes and parent cancellation; assert only outcomes permitted by the documented contract.
4. Verify the concurrency cap and absence of blocked result senders after early return.

## Evidence to collect

Draw the lifetime of request context, backend contexts and join ownership. Use deterministic backend barriers.

## Senior follow-up

When does hedged execution improve tail latency, and how would you limit its extra downstream load?

## Prerequisites

- [Goroutines, scheduling, and lifetime](../../topics/01-goroutines-scheduler-and-lifetime.md)
- [Context, deadlines, and errors](../../topics/05-context-deadlines-and-errors.md)
- [Worker pools, fan-out, and ordering](../../topics/06-worker-pools-and-patterns.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
