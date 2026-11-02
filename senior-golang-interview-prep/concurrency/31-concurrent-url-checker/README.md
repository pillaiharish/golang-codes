# Q31. Build a concurrent URL checker

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../30-graceful-shutdown/README.md) · [Next question](../32-happens-before/README.md)

**Level 6 — interview-sized capstone.**

## Your task

Given a list of URLs, check them concurrently, limit active HTTP requests to N, use a per-request timeout, collect one result per input, report errors, and support cancellation. Preserve input identity even when completions arrive out of order.

## Scenario questions — answer before opening the solution

1. A URL hangs forever; what bounds it?
2. One request fails; cancel all?
3. Can shared `results[index]` writes race?
4. How does the caller know which URL produced a result?
5. What if URLs come from untrusted input?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q31 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
