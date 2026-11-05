# Q58. Benchmark hot-path error creation
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../57-circuit-breaker-policy/README.md) · [Next](../59-legacy-error-migration/README.md)

**Level:** Senior | **Concepts:** allocations, sentinel vs formatting, benchmarking.

## Your task — code first

1. Compare shared sentinel returns against fmt.Errorf with identifiers.
2. Run go test -bench=. -benchmem.
3. Explain why optimizing error allocation without profiling can be premature.

## Failure cases and acceptance tests

- [ ] Both strategies return valid errors.
- [ ] Benchmark functions isolate each path.

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

