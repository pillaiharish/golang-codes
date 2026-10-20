# 08. Kafka producer concurrency and batching

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

Separate four boundaries: the application accepted a record; the client buffered it; the broker acknowledged it; a consumer applied its effect. Returning success at the first boundary does not prove the later ones. An asynchronous producer needs delivery-result handling, bounded outstanding work and a shutdown flush policy.

Batch by byte count, record count and maximum age. Larger batches can improve efficiency but increase waiting time and memory. Partition keys influence ordering and load distribution. Concurrent application calls and retries require careful ordering semantics: the same key alone does not define which concurrently submitted record was first.

Producer idempotence suppresses supported protocol-level retry duplicates; it does not deduplicate arbitrary application resubmissions or external database effects. `acks=all` relates to in-sync replicas, not every configured replica regardless of health; review replication, minimum ISR and failure assumptions together. Timeouts can leave outcomes uncertain.

Kafka broker documentation uses configuration names common to the Java client. Map behavior to the selected Go client's documented options and defaults; do not copy names blindly. Pin the client and broker versions for experiments. Check thread safety, callback execution, delivery guarantees, buffering limits, cancellation behavior and flush/close semantics in that exact client. More application workers can simply move backlog into the client.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. Does Produce returning nil prove delivery?

**Medium. Strong answer:** Only according to that API contract; distinguish enqueue return from delivery callback or acknowledged result.

### Q2. Why might one synchronous sender lose telemetry upstream?

**Hard. Strong answer:** Its service rate limits drain speed. Measure latency; evaluate bounded asynchronous production or batching against downstream capacity.

### Q3. Can retries reorder records?

**Hard. Strong answer:** Potentially, depending on producer configuration and client guarantees. Test per-partition sequence under failure, not only normal operation.

### Q4. What must Close guarantee?

**Hard. Strong answer:** Reject new admission, finish or classify accepted records, flush within a budget, and report unresolved outcomes honestly.

### Q5. Is enabling idempotence enough for exactly-once business effects?

**Expert. Strong answer:** No. Define the transaction boundary, application resubmissions, consumer isolation and external side effects.

## Implement and demonstrate

- [Task 07: Batched asynchronous publisher](../tasks/07-batched-async-publisher/README.md)
- [Task 08: Per-key ordered dispatcher](../tasks/08-per-key-dispatcher/README.md)
- [Task 12: Graceful service shutdown](../tasks/12-graceful-shutdown/README.md)
- [Task 14: Capstone: telemetry ingestion to Kafka](../tasks/14-telemetry-kafka-capstone/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://kafka.apache.org/41/configuration/producer-configs/](https://kafka.apache.org/41/configuration/producer-configs/)
- [https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo](https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo)
