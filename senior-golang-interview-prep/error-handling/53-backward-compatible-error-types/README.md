# Q53. Evolve sentinels without breaking errors.Is
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../52-commerce-error-taxonomy/README.md) · [Next](../54-cross-service-error-contract/README.md)

**Level:** Senior | **Concepts:** custom Is, compatibility, equality contracts.

## Your task — code first

1. Keep ErrNotFound exported and add NotFoundError{Resource,ID}.
2. Implement Is(target error) bool for compatibility.
3. Explain why callers using err==ErrNotFound require an explicit migration or retaining exact sentinel return.

## Failure cases and acceptance tests

- [ ] errors.Is(newError, ErrNotFound) succeeds.
- [ ] Direct == does not magically remain compatible.

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

