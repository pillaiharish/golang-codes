# Task 07: Batched asynchronous publisher

[Task index](../README.md) · [Concurrency index](../../README.md) · [Assessment](../../ASSESSMENT.md)

**Status: not implemented.** This is a task specification; write the implementation and tests yourself.

## Contract to implement

Build a publisher around a fake asynchronous sink before integrating Kafka. Publish returns an admission result, while a receipt or callback reports delivery. Bound queued bytes, records and outstanding sends. Flush on maximum batch records, bytes or age. Define payload ownership and callbacks exactly once per accepted record at the local API boundary. A timed-out remote operation may have unknown outcome; do not label it definitely undelivered. Pin a Go Kafka client only for the integration phase.

## Acceptance tests and failure injection

1. A low-rate batch flushes by age; a high-rate batch flushes by record or byte threshold; oversize records follow explicit policy.
2. Withhold sink acknowledgements and prove total buffered/outstanding work stays bounded.
3. Deliver callbacks out of order, return transient errors and simulate an acknowledgement lost after remote acceptance.
4. Close rejects new work, flushes within its budget and reports unresolved records without inventing successful delivery.

## Evidence to collect

Compare synchronous and asynchronous modes at identical offered load and fake latency. For Kafka, record broker/client versions and exact ack, retry, buffering and idempotence settings.

## Senior follow-up

Which ordering guarantee survives concurrent Publish calls for the same key, and how will you test it?

## Prerequisites

- [Memory ordering and data ownership](../../topics/02-memory-model-and-ownership.md)
- [Backpressure, rate limits, and capacity](../../topics/07-backpressure-rate-limits-and-capacity.md)
- [Kafka producer concurrency and batching](../../topics/08-kafka-producer-concurrency.md)
- [Testing, profiling, and observability](../../topics/11-testing-profiling-and-observability.md)

## Deliverables

- Implementation, contract-focused tests, and a small local demo where useful.
- An ownership/lifecycle explanation and explicit limits on records, bytes, active work and waiting callers where applicable.
- Actual commands/results, unresolved cases and trade-offs recorded using the [evidence template](../../NOTES.md).

Use the [local workflow](../README.md#local-workflow) when adding code. No pass result is implied by this task brief.
