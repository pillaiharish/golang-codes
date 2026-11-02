# Q33. Debug goroutines with pprof

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../32-happens-before/README.md)

**Level 6 — production debugging.**

## Your task

Run a small local HTTP server with pprof enabled, deliberately block a goroutine on a channel or mutex, and inspect goroutine profiles/stacks to locate the blocking site.

## Scenario questions — answer before opening the solution

1. Does a large goroutine count prove a leak?
2. What is the difference between a goroutine profile and a CPU profile?
3. Why not expose pprof publicly?
4. Can pprof identify who should close a channel?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q33 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
