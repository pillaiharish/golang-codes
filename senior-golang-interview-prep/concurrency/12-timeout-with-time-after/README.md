# Q12. Implement a timeout with `time.After`

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../11-select-multiple-channels/README.md) · [Next question](../13-context-cancellation/README.md)

**Level 3 — select, timeouts, and cancellation.**

## Your task

Start a simulated slow operation. Print `operation timed out` if no result arrives within 500 ms.

## Scenario questions — answer before opening the solution

1. What if the result and timer are both ready?
2. Does `time.After` kill the worker?
3. Why use `context.WithTimeout` in a server?
4. What if the worker sends on an unbuffered result channel after the timeout?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q12 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
