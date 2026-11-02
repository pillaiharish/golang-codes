# Q25. Implement `tee` with two outputs

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../24-or-done/README.md) · [Next question](../26-bridge-channel-of-channels/README.md)

**Level 5 — core concurrency patterns.**

## Your task

Write a `tee` function that accepts one input channel and returns two channels. **Every input value must reach both outputs**, then both outputs must close when the source ends.

## Scenario questions — answer before opening the solution

1. Is tee equivalent to two goroutines reading the same input?
2. What if one output has no receiver?
3. Are the two outputs guaranteed to receive a value simultaneously?
4. Can a tee silently drop one branch?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q25 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
