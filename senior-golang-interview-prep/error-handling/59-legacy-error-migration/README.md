# Q59. Migrate legacy string-based errors safely
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../58-benchmark-error-cost/README.md) · [Next](../60-design-review-error-checklist/README.md)

**Level:** Senior | **Concepts:** incremental change, regression tests, typed contracts.

## Your task — code first

1. Find a legacy function that emits a free-form error string.
2. Add a stable sentinel in a new adapter without requiring callers to parse text.
3. Write a phased migration plan that preserves existing external behavior.

## Failure cases and acceptance tests

- [ ] New callers can use errors.Is.
- [ ] Old displayed message remains suitable during transition.

## Explain before revealing

- What is the behavior on both success and failure?
- Which error identity or typed details can the caller inspect?
- Where does wrapping, logging, cancellation or retry policy belong?
- Which apparently reasonable implementation could conceal a failure?

## Practice workflow

1. Implement independently without opening the answer.
2. Write tests for expected failures and successes.
3. Compare with the [worked solution](solution/README.md), then refactor.
4. Run tests and static checks; record evidence in [revision notes](../NOTES.md).

**Mastery is earned by reproducing the solution and explaining its trade-offs**, not by reading the reference.

