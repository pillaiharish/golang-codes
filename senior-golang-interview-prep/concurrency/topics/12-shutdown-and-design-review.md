# 12. Graceful shutdown and senior design review

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

Model shutdown as states: accepting, draining, stopped. Specify the admission boundary, who owns each worker, how queued and active work are treated, and what callers observe during concurrent Close. Repeated Close should not panic. A bounded shutdown deadline does not magically stop a dependency that ignores cancellation.

In a service, withdraw readiness, stop admission, drain or cancel accepted work, flush downstream clients, join owned workers and release resources in a deliberate order. Readiness changes are not instantaneous traffic barriers. Kubernetes preStop work consumes the same termination grace budget; reserve time for cleanup. HTTP Server.Shutdown does not automatically join arbitrary background workers or handle every long-lived connection.

Review a design by asking for invariants, measured demand, failure modes and business consequences before choosing primitives. Kafka delivery, MQTT QoS, a Redis write and ClickHouse ingestion each have their own acknowledgement and durability boundary; none should be assumed to prove an end-to-end business effect merely from a successful local call.

A senior explanation names alternatives, explains why they were rejected, proposes evidence and a rollback, and delegates clear ownership. For telemetry, choose an explicit loss/latency budget; for billing or security audit data, revisit durability rather than importing a telemetry drop policy.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. What order should shutdown follow?

**Medium. Strong answer:** Stop admission before dismantling consumers; classify accepted work, honor a shared budget and join owners.

### Q2. Will readiness=false prevent all further requests immediately?

**Hard. Strong answer:** No. Account for propagation, existing connections and concurrent requests with an application admission barrier.

### Q3. Does HTTP Shutdown stop the Kafka publisher?

**Hard. Strong answer:** Only if the service explicitly coordinates that publisher and its lifetime.

### Q4. How do you review a proposed tenfold buffer increase?

**Hard. Strong answer:** Ask for evidence of burst duration versus sustained overload, memory cost, queue delay and failure behavior.

### Q5. Defend a lossy telemetry pipeline to another senior engineer.

**Expert. Strong answer:** State the data classification, measured loss/latency goals, sampling bias, alternatives, rollout and rollback; do not generalize it to durable business events.

## Implement and demonstrate

- [Task 10: Rebalance-safe consumer coordinator](../tasks/10-rebalance-coordinator/README.md)
- [Task 12: Graceful service shutdown](../tasks/12-graceful-shutdown/README.md)
- [Task 14: Capstone: telemetry ingestion to Kafka](../tasks/14-telemetry-kafka-capstone/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://pkg.go.dev/net/http#Server.Shutdown](https://pkg.go.dev/net/http#Server.Shutdown)
- [https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)
