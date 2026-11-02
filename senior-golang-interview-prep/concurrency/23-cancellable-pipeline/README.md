# Q23. Build a cancellable pipeline

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../22-fan-out-fan-in/README.md) · [Next question](../24-or-done/README.md)

**Level 5 — core concurrency patterns.**

## Your task

Build `generator -> square -> printer` with independent goroutines and channels. Then add cancellation so stopping the final consumer does not leak upstream producers.

## Scenario questions — answer before opening the solution

1. Why are both receive and send cancelable?
2. Does a closed output tell upstream to stop?
3. What happens when the printer consumes only one value?
4. How does a slow printer affect generator throughput?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q23 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
