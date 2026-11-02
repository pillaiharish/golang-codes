# Q28. Build a small pub/sub system

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../27-rate-limiter/README.md) · [Next question](../29-detect-goroutine-leak/README.md)

**Level 5 — core concurrency patterns.**

## Your task

Implement `Subscribe(topic)`, `Publish(topic, message)`, and `Unsubscribe`. Every active subscriber of a topic should receive each accepted message, subject to a documented slow-subscriber policy.

## Scenario questions — answer before opening the solution

1. Can multiple subscribers read one shared channel and each see every message?
2. What if a subscriber is slow?
3. What if unsubscribe races with publish?
4. Does this in-memory bus survive a process crash?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q28 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
