# Q27. Implement a rate limiter

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../26-bridge-channel-of-channels/README.md) · [Next question](../28-pub-sub/README.md)

**Level 5 — core concurrency patterns.**

## Your task

Process 20 tasks at a pace of **at most five starts per second**. Implement a simple ticker-based pacer and compare it with `golang.org/x/time/rate`. Explain why a concurrency bound may also be necessary.

## Scenario questions — answer before opening the solution

1. Does a ticker alone cap active goroutines?
2. Does a semaphore guarantee five requests per second?
3. What happens when a downstream service slows?
4. What does a token bucket add?
5. Can you safely use this for per-tenant fairness?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q27 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
