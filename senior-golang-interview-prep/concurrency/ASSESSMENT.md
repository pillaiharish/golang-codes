# Senior concurrency assessment

[Concurrency index](README.md) · [Tasks](tasks/README.md) · [Evidence notes](NOTES.md)

## What the senior bar means

A senior answer states the requirements and invariants, picks ownership and synchronization, anticipates failure, measures the result, and explains alternatives. The levels below are practice labels rather than employer-specific scoring:

- **Medium:** explain the mechanism accurately and illustrate it.
- **Hard:** reason through cancellation, failure, ordering and resource bounds.
- **Expert:** defend distributed effects, recovery, fairness, observability and operational trade-offs.

## Review rubric

Score each dimension 0 (missing), 1 (happy path), 2 (tested failure behavior), or 3 (evidence plus defended trade-offs).

| Dimension | Evidence expected |
|---|---|
| Correctness | Explicit invariants, ownership, synchronization and deterministic adversarial tests |
| Lifecycle | Admission boundary, cancellation, joins and bounded shutdown |
| Resource control | Bounds across records, bytes, active work, retries, waiters and client buffering |
| Delivery and ordering | Defined acknowledgement, replay, ordering and unknown-outcome semantics |
| Observability | Offered/admitted/completed work, failure reasons, queue age and useful diagnostics |
| Engineering judgment | Alternatives, scope, measured limits, operational rollout and rollback |

**Mandatory gates:** an unresolved data race, skipped unfinished consumer offset, send/close panic or unaccounted accepted work is a correctness gap regardless of total score. A timeout test cannot justify saying an uncooperative goroutine was terminated. Tests must exercise the stated contract; passing a race run alone is insufficient.

## Mock interview formats

These are suggested practice formats, not claims about a company's interview process.

- **45-minute coding:** 5 minutes clarify a task contract; 25 implement a core behavior; 10 test failure cases; 5 defend limitations.
- **60-minute design:** 10 requirements/capacity; 20 architecture and ownership; 15 failure/recovery; 10 observability/rollout; 5 alternatives.
- **30-minute incident review:** state impact, mitigation, evidence, root cause, follow-up and the teammate ownership needed to prevent recurrence.

## Design-review prompts

1. One synchronous Kafka sender falls behind. What evidence separates application serialization from broker saturation?
2. A proposal launches a goroutine for every record but adds a semaphore inside it. What remains unbounded?
3. A consumer commits the largest completed offset. Construct a counterexample with out-of-order completion.
4. A producer returns success before the delivery callback. What is the caller actually promised?
5. The service closes the queue while HTTP handlers still publish. Where is the admission barrier?
6. Every tenant ID is a metric label and has a permanent limiter entry. How do memory and cardinality scale?
7. A retry uses a fresh 30-second timeout at each layer. What happens to the caller's 200-ms budget?
8. A database update and Kafka publish happen separately. Explain the crash windows and compare outbox versus transaction boundaries.

For each review, describe severity, a concrete failure schedule, the smallest suitable change, a test and a rollout plan. Separate correctness objections from preferences. Explain what evidence would change your recommendation.
