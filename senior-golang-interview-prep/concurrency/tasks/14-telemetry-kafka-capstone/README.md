# Task 14: Capstone: telemetry ingestion to Kafka

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Start with the stated scenario: many producers, a 1,024-slot channel and one reader synchronously sending to Kafka, with roughly 9% reported drops. Treat 9% as the scenario input, not a measurement from this lab. Reproduce controlled overload and redesign the pipeline using measured evidence. Compare the baseline with at least two alternatives, such as bounded batching, bounded asynchronous production, multiple ordered lanes, sampling or durable spooling. Choose based on the data's business value and latency/loss budget.

## Acceptance tests and failure injection

1. Sweep offered rate, payload size, sink latency, error rate and key skew; include sustained overload and a temporary outage.
2. Bound all retained work, including retries and client buffers; reconcile admitted work with delivered, failed, unresolved and shutdown outcomes.
3. Demonstrate stated ordering and duplicate behavior under retry/rebalance where applicable.
4. Stop under load; test recovery from sink outage without a retry storm or unbounded memory growth.
5. Separate load-generator limitations from service limitations; report actual measurements with hardware and versions.

## Evidence to collect

Write a short RFC with requirements, capacity estimate, design alternatives, tests, metrics, rollout and rollback. Report offered/admitted/delivered rates, queue age, p50/p95/p99 latency, loss by reason, memory, CPU and shutdown time.

## Senior follow-up

How would the design change for billing events, security audit logs or tenant-isolated inference requests?

## Prerequisites

- [Backpressure, rate limits, and capacity](../../topics/07-backpressure-rate-limits-and-capacity.md)
- [Kafka producer concurrency and batching](../../topics/08-kafka-producer-concurrency.md)
- [Delivery semantics, idempotency, and retries](../../topics/10-delivery-idempotency-and-retry.md)
- [Testing, profiling, and observability](../../topics/11-testing-profiling-and-observability.md)
- [Graceful shutdown and senior design review](../../topics/12-shutdown-and-design-review.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
