# Q8. What happens when a channel is closed?

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../7-producer-consumer/README.md) · [Next question](../9-directional-channels/README.md)

**Level 2 — channel fundamentals.**

## Your task

Write independent tiny programs to demonstrate: receiving from a closed channel, sending to a closed channel, closing twice, and closing a nil channel. Predict whether each returns, blocks, or panics.

## Scenario questions — answer before opening the solution

1. Why can a receiver safely range over a closed channel?
2. Should the receiver close the channel?
3. How do many senders close safely?
4. Can `sync.Once` make repeated `close` harmless?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q8 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
