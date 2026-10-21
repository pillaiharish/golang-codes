# Task 11: Idempotent effects and transactional outbox

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

In phase A, build a transactional fake store supporting effect plus dedup marker in one atomic commit. In phase B, write domain changes and an outbox row in the same transaction, then relay to a fake broker. In phase C, integrate a real transactional database and Kafka if available. Use stable event IDs, define conflicting payload reuse and retention. Mark relay progress only according to acknowledgement policy; accept that replay can occur.

## Acceptance tests and failure injection

1. Crash before and after each transaction boundary; retry and verify one intended effect within the dedup horizon.
2. Simulate remote acceptance followed by lost acknowledgement; replay produces no duplicate protected effect.
3. Two consumers race the same ID; transaction/uniqueness semantics allow one effect.
4. An unavailable quarantine sink cannot justify committing past a poison event under a durable-quarantine policy.

## Evidence to collect

Create a crash-window table: durable state before crash, restart action and expected effect count. Clearly separate mock evidence from integration evidence.

## Senior follow-up

Can the same guarantee cover an external HTTP service without a compatible idempotency or transaction contract?

## Prerequisites

- [Kafka consumers, offsets, and rebalances](../../topics/09-kafka-consumer-offsets-and-rebalances.md)
- [Delivery semantics, idempotency, and retries](../../topics/10-delivery-idempotency-and-retry.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
