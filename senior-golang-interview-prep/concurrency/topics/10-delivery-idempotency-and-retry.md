# 10. Delivery semantics, idempotency, and retries

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

At-most-once allows loss to avoid replay; at-least-once allows replay to avoid silent skipping within stated durability assumptions. Exactly-once must name its boundary. A Kafka transaction can coordinate Kafka writes and consumer offsets in a supported workflow, but does not atomically commit an unrelated database or HTTP payment.

For a database effect, store the deduplication marker and the effect in the same transaction when possible. A marker written before a failed effect can suppress necessary work; a marker written after an effect can permit duplication. Define key scope, retention, payload conflicts and replay horizon. An in-memory set loses its guarantees on restart.

The transactional outbox commits domain changes and publication intent together. Its relay can publish duplicates after an acknowledgement/crash window; consumers still need suitable duplicate handling. A retry is appropriate only when the operation is safe and the failure is plausibly transient. Use bounded attempts, jitter and the original deadline; retry storms across layers multiply load.

A poison message requires a documented decision: stop the partition, quarantine durably, or intentionally discard with accountability. If quarantining permits offset advancement, ensure durable quarantine acceptance first. Unknown outcomes deserve their own classification rather than being mislabeled definitely failed.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. Why can at-least-once processing repeat an effect?

**Medium. Strong answer:** The effect can complete before the offset commit is durable, and the record reappears after restart.

### Q2. Why is a dedup marker outside the effect transaction unsafe?

**Hard. Strong answer:** Either ordering leaves a crash window that loses or duplicates the intended effect.

### Q3. Does an outbox eliminate all duplicates?

**Hard. Strong answer:** No. It closes the database/intent dual-write gap; relay acknowledgement uncertainty still needs duplicate handling.

### Q4. When may a poison record be skipped?

**Hard. Strong answer:** Only under the agreed policy, with durable quarantine or explicit authorized loss and auditable accounting.

### Q5. How do you choose dedup retention?

**Expert. Strong answer:** Match replay and recovery horizons, key scope and storage budget; explain what happens after expiration.

## Implement and demonstrate

- [Task 09: Parallel consumer completion frontier](../tasks/09-consumer-offset-frontier/README.md)
- [Task 10: Rebalance-safe consumer coordinator](../tasks/10-rebalance-coordinator/README.md)
- [Task 11: Idempotent effects and transactional outbox](../tasks/11-idempotency-and-outbox/README.md)
- [Task 14: Capstone: telemetry ingestion to Kafka](../tasks/14-telemetry-kafka-capstone/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://docs.aws.amazon.com/prescriptive-guidance/latest/cloud-design-patterns/transactional-outbox.html](https://docs.aws.amazon.com/prescriptive-guidance/latest/cloud-design-patterns/transactional-outbox.html)
- [https://kafka.apache.org/41/design/design/](https://kafka.apache.org/41/design/design/)
