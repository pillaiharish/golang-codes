# 09. Kafka consumers, offsets, and rebalances

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

For a classic consumer group, a partition is assigned to one group member at a time, but application code may still process fetched records concurrently. These notes do not describe Kafka share groups. Partition log order is not automatically effect-completion order.

A committed offset identifies the next position to resume from. Suppose delivered offsets 40, 41 and 42 run concurrently: 42 finishing first cannot justify committing 43 while 40 is unfinished. Track the earliest unfinished delivered record and advance only the safe completed frontier. Kafka offsets can have gaps; do not wait for every numerical offset to exist.

Commit before an effect risks loss after a crash; commit after an effect risks replay after a crash. Idempotency or a suitable transaction boundary handles replay. A successful fetch is not a completed business operation. Limit fetched-but-unfinished records, per-partition queues and completion tracking.

On revocation, stop admitting work for revoked partitions, cancel or drain according to a deadline, and fence stale completions with an assignment epoch. Cancellation cannot undo an external effect already committed. Commit only valid safe progress while respecting the client's rebalance protocol. Slow processing, blocked polling and rebalance callbacks have client-specific implications; test them against the pinned client.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. What does committing offset 43 mean?

**Medium. Strong answer:** Resume from 43; it does not mean only record 43 completed.

### Q2. 42 completes before 40: what can be committed?

**Hard. Strong answer:** Only the next position justified by the completed frontier; never skip an earlier unfinished delivered record.

### Q3. How do you preserve per-key effects within a partition?

**Hard. Strong answer:** Serial per-key execution or ordered application of results; partition order alone does not order parallel effects.

### Q4. Why fence callbacks after revocation?

**Hard. Strong answer:** A delayed completion from an old assignment must not mutate current progress or submit a stale commit.

### Q5. Can cancellation guarantee no duplicate external effect during rebalance?

**Expert. Strong answer:** No. Use an external idempotency/transaction mechanism or a supported fencing protocol as well as local lifecycle control.

## Implement and demonstrate

- [Task 08: Per-key ordered dispatcher](../tasks/08-per-key-dispatcher/README.md)
- [Task 09: Parallel consumer completion frontier](../tasks/09-consumer-offset-frontier/README.md)
- [Task 10: Rebalance-safe consumer coordinator](../tasks/10-rebalance-coordinator/README.md)
- [Task 11: Idempotent effects and transactional outbox](../tasks/11-idempotency-and-outbox/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://kafka.apache.org/41/design/design/](https://kafka.apache.org/41/design/design/)
- [https://kafka.apache.org/41/configuration/consumer-configs/](https://kafka.apache.org/41/configuration/consumer-configs/)
- [https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo](https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo)
