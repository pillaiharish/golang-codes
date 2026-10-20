# 04. Channels, select, and queue contracts

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

Bounded describes storage capacity; blocking describes what an operation does when it cannot proceed. They are independent. Buffered channels have fixed slot capacity; an unbuffered channel is a rendezvous, not an unbounded queue. Go channels never grow automatically like slices.

| Queue contract | When full or no receiver | Consequence |
|---|---|---|
| Bounded + blocking | Wait for space, cancellation or deadline | Backpressure; bound waiting callers too |
| Bounded + try/nonblocking | Return rejection immediately | Caller retries, drops or routes elsewhere explicitly |
| Unbounded + enqueue | Grow storage while resources allow | Overload becomes memory growth; not infinite or necessarily lock-free |
| Unbuffered + blocking | Wait for a receiver | Direct handoff, no queued slots |
| Unbuffered + try | Succeed only if handoff is ready | Rejection can occur despite an otherwise healthy consumer |

`select` with `default` avoids waiting for channel readiness; it does not promise a lock-free algorithm or bounded total method latency. If cancellation and send are both ready, selection does not prioritize cancellation. Define whether concurrent cancellation may still admit work. `len(ch)` is an observation, not a reservation.

Nil channel operations block; a nil case is disabled in select. Sending to a closed channel panics; a closed channel drains buffered values, then receives return the zero value with ok=false. A permanently ready closed receive can spin unless the loop exits or disables that case. The closer must prove no sends remain; multiple producers need coordinated ownership.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. Why can a buffered send block?

**Medium. Strong answer:** All slots are occupied; buffering postpones backpressure rather than eliminating it.

### Q2. Can a bounded queue still consume unbounded memory?

**Hard. Strong answer:** Yes, through blocked callers, retained payloads, retries or downstream buffers outside the queue.

### Q3. Does select prioritize ctx.Done()?

**Hard. Strong answer:** No. Specify and test allowed outcomes for simultaneous readiness.

### Q4. Who closes a multi-producer channel?

**Hard. Strong answer:** A coordinator after producers stop, or an owner with a synchronized admission/closure protocol.

### Q5. Does nonblocking mean lossless?

**Expert. Strong answer:** No. Rejection requires an explicit caller policy; hidden retries may recreate unbounded buffering.

## Implement and demonstrate

- [Task 02: Bounded producer-consumer queue](../tasks/02-bounded-queue/README.md)
- [Task 03: Bounded worker pool](../tasks/03-bounded-worker-pool/README.md)
- [Task 12: Graceful service shutdown](../tasks/12-graceful-shutdown/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://go.dev/ref/spec#Channel_types](https://go.dev/ref/spec#Channel_types)
- [https://go.dev/ref/spec#Select_statements](https://go.dev/ref/spec#Select_statements)
- [https://go.dev/blog/pipelines](https://go.dev/blog/pipelines)
