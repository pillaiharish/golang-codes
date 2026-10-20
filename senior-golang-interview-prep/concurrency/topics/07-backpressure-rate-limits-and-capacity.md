# 07. Backpressure, rate limits, and capacity

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

Rate limits bound work admitted over time; concurrency limits bound work active simultaneously; queues absorb temporary bursts. None substitutes for the others. Budget bytes as well as record count because payload sizes vary. Include client buffers, retry queues, blocked publishers and batches in the memory budget.

If arrivals are 10,000 records/s and sustained completion is 8,000 records/s, a previously empty 1,024-slot queue fills in approximately 1,024 / 2,000 = 0.512 seconds under those simplified constant-rate assumptions. Increasing capacity buys time, not sustainable throughput. Little's law relates average in-system work, effective throughput and average time in a stable system; it is not a p99 latency predictor.

Policies include blocking with deadlines, reject-new, drop-oldest, sampling, priority shedding and durable spooling. Choose according to the value and durability requirements of the data. Blocking telemetry on a business request may damage the service it observes. A disk spool also requires bounds, disk-full handling, recovery and replay deduplication.

Measure offered, admitted and completed load separately. A fast benchmark that rejects most work proves little about delivery throughput. Fairness needs per-tenant or per-key budgets plus a global budget; per-tenant limiters alone allow aggregate overload and unbounded tenant-state growth.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. Rate limiting versus concurrency limiting?

**Medium. Strong answer:** Requests per interval versus active requests; a slow dependency can exhaust concurrency even at a modest rate.

### Q2. How would you size a buffer?

**Hard. Strong answer:** Use measured burst excess, payload sizes and latency budgets; include all retained state and validate under realistic load.

### Q3. Why did a larger queue worsen the outage?

**Hard. Strong answer:** It retained stale work, increased memory and queue delay while service capacity remained inadequate.

### Q4. Which metrics distinguish queue loss from broker failure?

**Hard. Strong answer:** Separate admission rejection, accepted work, terminal success/failure, unknown outcome and shutdown abandonment.

### Q5. How do you stop a hot tenant monopolizing capacity?

**Expert. Strong answer:** Combine global and per-tenant admission with fair scheduling, bounded tenant metadata and explicit cleanup.

## Implement and demonstrate

- [Task 02: Bounded producer-consumer queue](../tasks/02-bounded-queue/README.md)
- [Task 06: Rate, concurrency, and tenant admission](../tasks/06-rate-and-concurrency-limits/README.md)
- [Task 07: Batched asynchronous publisher](../tasks/07-batched-async-publisher/README.md)
- [Task 14: Capstone: telemetry ingestion to Kafka](../tasks/14-telemetry-kafka-capstone/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://pkg.go.dev/golang.org/x/time/rate](https://pkg.go.dev/golang.org/x/time/rate)
- [https://pkg.go.dev/golang.org/x/sync/semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
- [https://prometheus.io/docs/practices/naming/](https://prometheus.io/docs/practices/naming/)
