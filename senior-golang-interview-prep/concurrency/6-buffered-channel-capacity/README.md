# Q6. Understand buffered channels

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../5-unbuffered-channel-send-receive/README.md) · [Next question](../7-producer-consumer/README.md)

**Level 2 — channel fundamentals.**

## Your task

Create a capacity-two channel. Send two integers without starting a receiver. Predict what happens on a third send, then add a receiver to unblock it.

## Scenario questions — answer before opening the solution

1. What if capacity is zero?
2. What if the buffer is empty and no sender exists?
3. Is capacity two a promise that the application never blocks?
4. Does a buffered channel cap the number of spawned goroutines?
5. What if the buffer has values when closed?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q6 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
