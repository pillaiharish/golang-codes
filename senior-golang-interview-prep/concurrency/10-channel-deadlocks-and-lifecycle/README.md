# Q10. Create and fix deadlocks

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../9-directional-channels/README.md) · [Next question](../11-select-multiple-channels/README.md)

**Level 2 — channel fundamentals.**

## Your task

Create four small cases, predict the outcome, and fix them:

1. Send on an unbuffered channel without a receiver.
2. Receive from an open empty channel without a sender.
3. Range over a channel when the producer never closes it.
4. Start a goroutine but return from `main` before joining it.

## Scenario questions — answer before opening the solution

1. A buffered channel of size 1:
2. A never-closing channel with cancellation:
3. Two mutexes locked in opposite orders:
4. A test that hangs:

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q10 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
