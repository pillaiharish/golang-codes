# 11. Testing, profiling, and observability

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

Test the protocol, not the scheduler's luck. Use barriers, controlled fake dependencies and injectable clocks to establish precise states: queue full, worker started, effect applied, callback withheld. A timeout bounds a failed test; sleeping for a guessed duration does not prove a state transition happened.

Race detection finds data races in exercised executions; it does not prove absence of all races, logical races, deadlocks or leaks. Assert resource bounds and joins directly. Run adversarial sequences involving cancellation, retry, close and completion. A fake dependency permits deterministic failure injection; it does not replace real Kafka tests for protocol behavior.

CPU profiles explain CPU execution, not every cause of wall-clock latency. Heap profiles show allocation/retention evidence; goroutine stacks reveal blocking locations; block/mutex profiles and execution traces help explain waiting and scheduling. Enable relevant profiles deliberately and understand overhead. Compare evidence before and after one targeted change.

Use low-cardinality metrics: queue records/bytes, active work, oldest age, enqueue latency, delivery latency, rejection reason, retry count and unresolved outcomes. Put individual event IDs in logs/traces, not metric labels. Report offered and successful throughput alongside latency and loss; avoid benchmarks that hide overload by pacing only on completions.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. What does a clean race run prove?

**Medium. Strong answer:** No race was reported in that execution; it is evidence, not a universal correctness proof.

### Q2. How do you test a full queue without sleeps?

**Hard. Strong answer:** Block the fake sink at a barrier, establish active work, fill known capacity, then exercise the next admission.

### Q3. Why is CPU low while p99 is high?

**Hard. Strong answer:** Investigate queueing, I/O and lock waits, throttling and runnable delay with appropriate profiles/traces.

### Q4. How do you test shutdown leaks reliably?

**Hard. Strong answer:** Assert owned workers exit via join signals after controlled release/cancellation; use stack evidence for failures rather than exact global goroutine counts.

### Q5. How can a benchmark claim 100k/s while delivery is poor?

**Expert. Strong answer:** It may measure attempts or immediate rejections. Report completion, failure, queue age and end-to-end latency under open-loop offered load.

## Implement and demonstrate

- [Task 01: Concurrent state and immutable snapshots](../tasks/01-shared-state-and-snapshots/README.md)
- [Task 07: Batched asynchronous publisher](../tasks/07-batched-async-publisher/README.md)
- [Task 13: Race, deadlock, and leak diagnosis lab](../tasks/13-debugging-lab/README.md)
- [Task 14: Capstone: telemetry ingestion to Kafka](../tasks/14-telemetry-kafka-capstone/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://go.dev/doc/articles/race_detector](https://go.dev/doc/articles/race_detector)
- [https://go.dev/doc/diagnostics](https://go.dev/doc/diagnostics)
- [https://prometheus.io/docs/practices/naming/](https://prometheus.io/docs/practices/naming/)
