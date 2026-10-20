# 06. Worker pools, fan-out, and ordering

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

A worker pool bounds active work only if admission and all additional buffers are also controlled. Launching a goroutine per request that waits on a semaphore bounds execution but can leave unlimited waiting goroutines. Acquire admission before spawning or use a fixed set of workers plus an explicit queue/rejection policy.

Fan-out distributes work; fan-in combines results. Close a shared result stream only when every sender has exited. An early-return consumer must cancel producers or drain their output according to the protocol. Preserve result ordering only if required: completion order often differs from input order. A reorder buffer needs a bound when an early item stalls.

Per-key ordering can be implemented through serial execution lanes, but a hot key can dominate its lane and collide with unrelated keys. Define idle-key cleanup and limits on active keys. Global ordering often sacrifices parallelism unnecessarily. A singleflight group coalesces concurrent same-key calls; it is neither a persistent cache nor durable effect deduplication. Shared work cancellation must be separated from one waiter's cancellation.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. What exactly does a pool of eight workers bound?

**Medium. Strong answer:** At most eight active worker operations, not payload bytes, upstream waiters or downstream client buffers.

### Q2. How can an early fan-in return leak goroutines?

**Hard. Strong answer:** Remaining senders wait forever unless cancellation or draining is included in the protocol.

### Q3. How do you preserve per-device order with parallel processing?

**Hard. Strong answer:** Choose key assignment and serial per-key execution; discuss hot keys, retries and bounded pending state.

### Q4. Should one canceled cache caller abort all shared loads?

**Hard. Strong answer:** Specify ownership of shared work and independent waiter cancellation; do not accidentally couple unrelated callers.

### Q5. How do you bound ordered output with one stalled item?

**Expert. Strong answer:** Bound the reorder window and apply admission backpressure or a documented failure policy.

## Implement and demonstrate

- [Task 03: Bounded worker pool](../tasks/03-bounded-worker-pool/README.md)
- [Task 04: Cancellable fan-out and first-success](../tasks/04-cancellable-fanout/README.md)
- [Task 05: Cache loading with shared work](../tasks/05-cache-and-singleflight/README.md)
- [Task 08: Per-key ordered dispatcher](../tasks/08-per-key-dispatcher/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://go.dev/blog/pipelines](https://go.dev/blog/pipelines)
- [https://pkg.go.dev/golang.org/x/sync/singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight)
- [https://pkg.go.dev/golang.org/x/sync/semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
