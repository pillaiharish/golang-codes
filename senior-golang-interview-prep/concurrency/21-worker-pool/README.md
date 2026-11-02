# Q21. Implement a worker pool

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../20-semaphore-limit-concurrency/README.md) · [Next question](../22-fan-out-fan-in/README.md)

**Level 5 — core concurrency patterns.**

## Your task

Create **3 workers**, **10 jobs**, and a results channel. Every worker consumes jobs and returns `job * 2`; collect all ten results without leaking goroutines or closing the results channel early.

## Scenario questions — answer before opening the solution

1. Is result order guaranteed?
2. Who closes `results`?
3. What if the consumer stops reading early?
4. How do you bound memory?
5. What about job failures?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q21 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
