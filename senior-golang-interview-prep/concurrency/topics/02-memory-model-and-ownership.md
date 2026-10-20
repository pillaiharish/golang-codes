# 02. Memory ordering and data ownership

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

A data race is unsynchronized conflicting access to the same memory with at least one write. A logical race can also exist in race-free code: separately locking a capacity check and insertion still permits over-admission. State the invariant and protect the whole transition.

A channel send orders prior writes before the corresponding receive completes. Closing a channel orders prior writes before a receive that observes closure; merely reading an earlier buffered item is not that closure observation. Goroutine exit alone does not publish results. Use documented synchronization rather than assumptions about timing.

Sending a slice copies its descriptor, not its backing array. Sending a pointer does not transfer ownership automatically. Agree on immutability, copying, or exclusive ownership until an acknowledgement returns. A shallow copy of a struct containing a map still aliases its entries. Publishing an immutable snapshot through an atomic pointer only works if the entire reachable snapshot remains immutable.

Example to reason about: a producer reuses `buf` immediately after sending `buf[:n]`; a slower consumer may read the next message's bytes. A channel buffer can make the timing failure more likely without being the underlying cause.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. Why is count++ unsafe across goroutines?

**Medium. Strong answer:** It is a read-modify-write operation; protect the shared counter with a mutex or appropriate atomic operation.

### Q2. Can race-free code oversell a capacity limit?

**Hard. Strong answer:** Yes. A check followed by a separately synchronized update can violate the combined invariant.

### Q3. Is sending []byte safe ownership transfer?

**Hard. Strong answer:** Only by contract. Copy bytes or stop accessing the backing array until ownership returns.

### Q4. Can two atomics publish a consistent pair of fields?

**Hard. Strong answer:** Individual atomicity does not make a multi-field transaction. Use a lock or atomically publish an immutable aggregate.

### Q5. How would you review a snapshot API?

**Expert. Strong answer:** Trace all aliases, nested maps and slices, writers and publication. Readers must not mutate shared snapshots.

## Implement and demonstrate

- [Task 01: Concurrent state and immutable snapshots](../tasks/01-shared-state-and-snapshots/README.md)
- [Task 02: Bounded producer-consumer queue](../tasks/02-bounded-queue/README.md)
- [Task 07: Batched asynchronous publisher](../tasks/07-batched-async-publisher/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://go.dev/ref/mem](https://go.dev/ref/mem)
- [https://pkg.go.dev/sync/atomic](https://pkg.go.dev/sync/atomic)
