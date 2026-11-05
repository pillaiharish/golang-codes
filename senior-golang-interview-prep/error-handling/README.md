# Senior Go error handling — learning-by-solving track

[Senior Go preparation](../README.md) · [Existing error examples](../../error-handling/README.md) · [Study plan](STUDY-PLAN.md) · [Revision notes](NOTES.md)

This track contains **60 numbered questions across six levels**, each with a question-only README and a sibling `solution/` containing a worked explanation, an independent Go module, runnable demonstration and failure-case unit tests.

## How to use it

1. Open a **question README**, write your own Go code, and test failure paths *before* opening `solution/`.
2. Compare the reference explanation, reproduce the code without copying, and explain its trade-offs aloud.
3. Run tests in that question's solution folder or use [verify.sh](verify.sh) for the whole track.
4. Record what still feels unclear in [NOTES.md](NOTES.md), then revisit the same question from memory.
5. Complete the [six capstone project specifications](projects/README.md) as you advance. These projects are intentionally left as exercises; the Q1–Q60 solutions are provided.

**Prerequisite:** Go 1.21+ (Go 1.20 introduced `errors.Join` and multi-`%w`; these examples also use Go 1.21's `log/slog`). Q42 and Q48 use external Go modules.

## Question index

### Level 1: Starter (Q1–Q10)

**Focus:** error interface, nil, errors.New, fmt.Errorf, defer, CLI and logging. **Capstone:** [CLI calculator](projects/1-cli-calculator/README.md).

| # | Question to solve | Worked solution and tests |
|---:|---|---|
| Q1 | [Implement the built-in error interface](1-error-interface/README.md) | [Explanation](1-error-interface/solution/README.md) · [main.go](1-error-interface/solution/main.go) · [tests](1-error-interface/solution/main_test.go) |
| Q2 | [Return and check an error from Divide](2-divide-return-error/README.md) | [Explanation](2-divide-return-error/solution/README.md) · [main.go](2-divide-return-error/solution/main.go) · [tests](2-divide-return-error/solution/main_test.go) |
| Q3 | [Understand the zero value of error](3-nil-error-zero-value/README.md) | [Explanation](3-nil-error-zero-value/solution/README.md) · [main.go](3-nil-error-zero-value/solution/main.go) · [tests](3-nil-error-zero-value/solution/main_test.go) |
| Q4 | [Compare errors.New and fmt.Errorf](4-new-vs-errorf/README.md) | [Explanation](4-new-vs-errorf/solution/README.md) · [main.go](4-new-vs-errorf/solution/main.go) · [tests](4-new-vs-errorf/solution/main_test.go) |
| Q5 | [Add context to parsing failures](5-parse-age-context/README.md) | [Explanation](5-parse-age-context/solution/README.md) · [main.go](5-parse-age-context/solution/main.go) · [tests](5-parse-age-context/solution/main_test.go) |
| Q6 | [Show why ignoring errors hides failure](6-do-not-ignore-errors/README.md) | [Explanation](6-do-not-ignore-errors/solution/README.md) · [main.go](6-do-not-ignore-errors/solution/main.go) · [tests](6-do-not-ignore-errors/solution/main_test.go) |
| Q7 | [Use defer and handle file cleanup](7-defer-file-close/README.md) | [Explanation](7-defer-file-close/solution/README.md) · [main.go](7-defer-file-close/solution/main.go) · [tests](7-defer-file-close/solution/main_test.go) |
| Q8 | [Return a nonzero CLI exit status](8-cli-error-exit/README.md) | [Explanation](8-cli-error-exit/solution/README.md) · [main.go](8-cli-error-exit/solution/main.go) · [tests](8-cli-error-exit/solution/main_test.go) |
| Q9 | [Return errors and log once at the boundary](9-return-vs-log-error/README.md) | [Explanation](9-return-vs-log-error/solution/README.md) · [main.go](9-return-vs-log-error/solution/main.go) · [tests](9-return-vs-log-error/solution/main_test.go) |
| Q10 | [Write composable Go error messages](10-composable-error-messages/README.md) | [Explanation](10-composable-error-messages/solution/README.md) · [main.go](10-composable-error-messages/solution/main.go) · [tests](10-composable-error-messages/solution/main_test.go) |

### Level 2: Beginner (Q11–Q20)

**Focus:** sentinels, typed errors, errors.Is/As, typed-nil, panic/recover. **Capstone:** [In-memory user store](projects/2-in-memory-store/README.md).

| # | Question to solve | Worked solution and tests |
|---:|---|---|
| Q11 | [Return a sentinel error from FindUser](11-sentinel-not-found/README.md) | [Explanation](11-sentinel-not-found/solution/README.md) · [main.go](11-sentinel-not-found/solution/main.go) · [tests](11-sentinel-not-found/solution/main_test.go) |
| Q12 | [Replace message matching with error identity](12-avoid-error-string-comparison/README.md) | [Explanation](12-avoid-error-string-comparison/solution/README.md) · [main.go](12-avoid-error-string-comparison/solution/main.go) · [tests](12-avoid-error-string-comparison/solution/main_test.go) |
| Q13 | [Create and inspect a ValidationError](13-typed-validation-error/README.md) | [Explanation](13-typed-validation-error/solution/README.md) · [main.go](13-typed-validation-error/solution/main.go) · [tests](13-typed-validation-error/solution/main_test.go) |
| Q14 | [Match wrapped sentinels with errors.Is](14-errors-is-wrapped-sentinel/README.md) | [Explanation](14-errors-is-wrapped-sentinel/solution/README.md) · [main.go](14-errors-is-wrapped-sentinel/solution/main.go) · [tests](14-errors-is-wrapped-sentinel/solution/main_test.go) |
| Q15 | [Extract typed errors with errors.As](15-errors-as-typed/README.md) | [Explanation](15-errors-as-typed/solution/README.md) · [main.go](15-errors-as-typed/solution/main.go) · [tests](15-errors-as-typed/solution/main_test.go) |
| Q16 | [Reproduce and fix a typed-nil error](16-typed-nil-error/README.md) | [Explanation](16-typed-nil-error/solution/README.md) · [main.go](16-typed-nil-error/solution/main.go) · [tests](16-typed-nil-error/solution/main_test.go) |
| Q17 | [Use panic only for programmer invariants](17-panic-only-for-invariants/README.md) | [Explanation](17-panic-only-for-invariants/solution/README.md) · [main.go](17-panic-only-for-invariants/solution/main.go) · [tests](17-panic-only-for-invariants/solution/main_test.go) |
| Q18 | [Recover only in the panicking goroutine](18-recover-deferred-boundary/README.md) | [Explanation](18-recover-deferred-boundary/solution/README.md) · [main.go](18-recover-deferred-boundary/solution/main.go) · [tests](18-recover-deferred-boundary/solution/main_test.go) |
| Q19 | [Implement a documented Must helper](19-must-pattern/README.md) | [Explanation](19-must-pattern/solution/README.md) · [main.go](19-must-pattern/solution/main.go) · [tests](19-must-pattern/solution/main_test.go) |
| Q20 | [Write table-driven tests for ParseAge](20-table-driven-error-tests/README.md) | [Explanation](20-table-driven-error-tests/solution/README.md) · [main.go](20-table-driven-error-tests/solution/main.go) · [tests](20-table-driven-error-tests/solution/main_test.go) |

### Level 3: Intermediate (Q21–Q30)

**Focus:** %w, Unwrap, deep chains, errors.Join, config classification. **Capstone:** [Configuration loader](projects/3-config-loader/README.md).

| # | Question to solve | Worked solution and tests |
|---:|---|---|
| Q21 | [Compare %v and %w on an error chain](21-percent-v-vs-percent-w/README.md) | [Explanation](21-percent-v-vs-percent-w/solution/README.md) · [main.go](21-percent-v-vs-percent-w/solution/main.go) · [tests](21-percent-v-vs-percent-w/solution/main_test.go) |
| Q22 | [Implement Error and Unwrap methods](22-custom-context-wrapper/README.md) | [Explanation](22-custom-context-wrapper/solution/README.md) · [main.go](22-custom-context-wrapper/solution/main.go) · [tests](22-custom-context-wrapper/solution/main_test.go) |
| Q23 | [Trace errors.Is through multiple wrappers](23-traverse-error-chain/README.md) | [Explanation](23-traverse-error-chain/solution/README.md) · [main.go](23-traverse-error-chain/solution/main.go) · [tests](23-traverse-error-chain/solution/main_test.go) |
| Q24 | [Extract DBError through a chain with errors.As](24-extract-db-error/README.md) | [Explanation](24-extract-db-error/solution/README.md) · [main.go](24-extract-db-error/solution/main.go) · [tests](24-extract-db-error/solution/main_test.go) |
| Q25 | [Preserve cause and category while reading config](25-contextual-config-read/README.md) | [Explanation](25-contextual-config-read/solution/README.md) · [main.go](25-contextual-config-read/solution/main.go) · [tests](25-contextual-config-read/solution/main_test.go) |
| Q26 | [Structure concise error messages across layers](26-layered-error-messages/README.md) | [Explanation](26-layered-error-messages/solution/README.md) · [main.go](26-layered-error-messages/solution/main.go) · [tests](26-layered-error-messages/solution/main_test.go) |
| Q27 | [Collect field errors with errors.Join](27-join-validation-errors/README.md) | [Explanation](27-join-validation-errors/solution/README.md) · [main.go](27-join-validation-errors/solution/main.go) · [tests](27-join-validation-errors/solution/main_test.go) |
| Q28 | [Verify errors.Join with nil inputs](28-join-nil-errors/README.md) | [Explanation](28-join-nil-errors/solution/README.md) · [main.go](28-join-nil-errors/solution/main.go) · [tests](28-join-nil-errors/solution/main_test.go) |
| Q29 | [Wrap multiple errors with fmt.Errorf](29-multiple-percent-w/README.md) | [Explanation](29-multiple-percent-w/solution/README.md) · [main.go](29-multiple-percent-w/solution/main.go) · [tests](29-multiple-percent-w/solution/main_test.go) |
| Q30 | [Debug and repair a broken error chain](30-debug-errors-is-failure/README.md) | [Explanation](30-debug-errors-is-failure/solution/README.md) · [main.go](30-debug-errors-is-failure/solution/main.go) · [tests](30-debug-errors-is-failure/solution/main_test.go) |

### Level 4: Upper intermediate (Q31–Q40)

**Focus:** testing, structured logging, context, third-party adapters, static analysis. **Capstone:** [File processor](projects/4-file-processor/README.md).

| # | Question to solve | Worked solution and tests |
|---:|---|---|
| Q31 | [Assert error identity and type in tests](31-test-errors-structurally/README.md) | [Explanation](31-test-errors-structurally/solution/README.md) · [main.go](31-test-errors-structurally/solution/main.go) · [tests](31-test-errors-structurally/solution/main_test.go) |
| Q32 | [Use table-driven validation error tests](32-table-tests-validation/README.md) | [Explanation](32-table-tests-validation/solution/README.md) · [main.go](32-table-tests-validation/solution/main.go) · [tests](32-table-tests-validation/solution/main_test.go) |
| Q33 | [Choose one error logging boundary](33-avoid-double-logging/README.md) | [Explanation](33-avoid-double-logging/solution/README.md) · [main.go](33-avoid-double-logging/solution/main.go) · [tests](33-avoid-double-logging/solution/main_test.go) |
| Q34 | [Log failures with structured slog fields](34-structured-slog-errors/README.md) | [Explanation](34-structured-slog-errors/solution/README.md) · [main.go](34-structured-slog-errors/solution/main.go) · [tests](34-structured-slog-errors/solution/main_test.go) |
| Q35 | [Return context cancellation and deadline errors](35-context-cancellation-errors/README.md) | [Explanation](35-context-cancellation-errors/solution/README.md) · [main.go](35-context-cancellation-errors/solution/main.go) · [tests](35-context-cancellation-errors/solution/main_test.go) |
| Q36 | [Wrap context errors without changing their policy](36-wrap-context-errors/README.md) | [Explanation](36-wrap-context-errors/solution/README.md) · [main.go](36-wrap-context-errors/solution/main.go) · [tests](36-wrap-context-errors/solution/main_test.go) |
| Q37 | [Translate third-party errors into domain errors](37-translate-third-party-errors/README.md) | [Explanation](37-translate-third-party-errors/solution/README.md) · [main.go](37-translate-third-party-errors/solution/main.go) · [tests](37-translate-third-party-errors/solution/main_test.go) |
| Q38 | [Avoid redundant context in error chains](38-avoid-overwrapping/README.md) | [Explanation](38-avoid-overwrapping/solution/README.md) · [main.go](38-avoid-overwrapping/solution/main.go) · [tests](38-avoid-overwrapping/solution/main_test.go) |
| Q39 | [Build a small application error taxonomy](39-classify-app-errors/README.md) | [Explanation](39-classify-app-errors/solution/README.md) · [main.go](39-classify-app-errors/solution/main.go) · [tests](39-classify-app-errors/solution/main_test.go) |
| Q40 | [Use linters to catch error mistakes](40-static-error-analysis/README.md) | [Explanation](40-static-error-analysis/solution/README.md) · [main.go](40-static-error-analysis/solution/main.go) · [tests](40-static-error-analysis/solution/main_test.go) |

### Level 5: Advanced (Q41–Q50)

**Focus:** goroutine propagation, errgroup, retries, HTTP/gRPC, SQL, metrics. **Capstone:** [Concurrent URL health checker](projects/5-url-health-checker/README.md).

| # | Question to solve | Worked solution and tests |
|---:|---|---|
| Q41 | [Collect errors from goroutines safely](41-goroutine-error-propagation/README.md) | [Explanation](41-goroutine-error-propagation/solution/README.md) · [main.go](41-goroutine-error-propagation/solution/main.go) · [tests](41-goroutine-error-propagation/solution/main_test.go) |
| Q42 | [Cancel sibling tasks using errgroup](42-errgroup-fail-fast-cancellation/README.md) | [Explanation](42-errgroup-fail-fast-cancellation/solution/README.md) · [main.go](42-errgroup-fail-fast-cancellation/solution/main.go) · [tests](42-errgroup-fail-fast-cancellation/solution/main_test.go) |
| Q43 | [Join validation errors from a worker pool](43-concurrent-multierror-validation/README.md) | [Explanation](43-concurrent-multierror-validation/solution/README.md) · [main.go](43-concurrent-multierror-validation/solution/main.go) · [tests](43-concurrent-multierror-validation/solution/main_test.go) |
| Q44 | [Distinguish request timeout from dependency failure](44-context-timeout/README.md) | [Explanation](44-context-timeout/solution/README.md) · [main.go](44-context-timeout/solution/main.go) · [tests](44-context-timeout/solution/main_test.go) |
| Q45 | [Implement bounded, cancellable retries](45-bounded-retry-backoff/README.md) | [Explanation](45-bounded-retry-backoff/solution/README.md) · [main.go](45-bounded-retry-backoff/solution/main.go) · [tests](45-bounded-retry-backoff/solution/main_test.go) |
| Q46 | [Classify retryability structurally](46-retryable-error-type/README.md) | [Explanation](46-retryable-error-type/solution/README.md) · [main.go](46-retryable-error-type/solution/main.go) · [tests](46-retryable-error-type/solution/main_test.go) |
| Q47 | [Map domain errors to safe HTTP JSON](47-http-domain-error-responses/README.md) | [Explanation](47-http-domain-error-responses/solution/README.md) · [main.go](47-http-domain-error-responses/solution/main.go) · [tests](47-http-domain-error-responses/solution/main_test.go) |
| Q48 | [Return idiomatic gRPC status codes](48-grpc-status-errors/README.md) | [Explanation](48-grpc-status-errors/solution/README.md) · [main.go](48-grpc-status-errors/solution/main.go) · [tests](48-grpc-status-errors/solution/main_test.go) |
| Q49 | [Translate sql.ErrNoRows at repository boundary](49-translate-sql-errors/README.md) | [Explanation](49-translate-sql-errors/solution/README.md) · [main.go](49-translate-sql-errors/solution/main.go) · [tests](49-translate-sql-errors/solution/main_test.go) |
| Q50 | [Count errors without high-cardinality labels](50-observable-error-metrics/README.md) | [Explanation](50-observable-error-metrics/solution/README.md) · [main.go](50-observable-error-metrics/solution/main.go) · [tests](50-observable-error-metrics/solution/main_test.go) |

### Level 6: Senior (Q51–Q60)

**Focus:** team standards, taxonomy, compatibility, SLOs, incident response and migrations. **Capstone:** [Order/payment service](projects/6-payment-order-service/README.md).

| # | Question to solve | Worked solution and tests |
|---:|---|---|
| Q51 | [Define and enforce a team error policy](51-team-error-policy/README.md) | [Explanation](51-team-error-policy/solution/README.md) · [main.go](51-team-error-policy/solution/main.go) · [tests](51-team-error-policy/solution/main_test.go) |
| Q52 | [Design a domain error taxonomy for commerce](52-commerce-error-taxonomy/README.md) | [Explanation](52-commerce-error-taxonomy/solution/README.md) · [main.go](52-commerce-error-taxonomy/solution/main.go) · [tests](52-commerce-error-taxonomy/solution/main_test.go) |
| Q53 | [Evolve sentinels without breaking errors.Is](53-backward-compatible-error-types/README.md) | [Explanation](53-backward-compatible-error-types/solution/README.md) · [main.go](53-backward-compatible-error-types/solution/main.go) · [tests](53-backward-compatible-error-types/solution/main_test.go) |
| Q54 | [Design safe errors across service boundaries](54-cross-service-error-contract/README.md) | [Explanation](54-cross-service-error-contract/solution/README.md) · [main.go](54-cross-service-error-contract/solution/main.go) · [tests](54-cross-service-error-contract/solution/main_test.go) |
| Q55 | [Decide what counts against availability SLOs](55-slo-error-budget-accounting/README.md) | [Explanation](55-slo-error-budget-accounting/solution/README.md) · [main.go](55-slo-error-budget-accounting/solution/main.go) · [tests](55-slo-error-budget-accounting/solution/main_test.go) |
| Q56 | [Triage a downstream failure without log flooding](56-debug-error-storm/README.md) | [Explanation](56-debug-error-storm/solution/README.md) · [main.go](56-debug-error-storm/solution/main.go) · [tests](56-debug-error-storm/solution/main_test.go) |
| Q57 | [Implement a minimal circuit breaker state machine](57-circuit-breaker-policy/README.md) | [Explanation](57-circuit-breaker-policy/solution/README.md) · [main.go](57-circuit-breaker-policy/solution/main.go) · [tests](57-circuit-breaker-policy/solution/main_test.go) |
| Q58 | [Benchmark hot-path error creation](58-benchmark-error-cost/README.md) | [Explanation](58-benchmark-error-cost/solution/README.md) · [main.go](58-benchmark-error-cost/solution/main.go) · [tests](58-benchmark-error-cost/solution/main_test.go) |
| Q59 | [Migrate legacy string-based errors safely](59-legacy-error-migration/README.md) | [Explanation](59-legacy-error-migration/solution/README.md) · [main.go](59-legacy-error-migration/solution/main.go) · [tests](59-legacy-error-migration/solution/main_test.go) |
| Q60 | [Review error policy before implementation](60-design-review-error-checklist/README.md) | [Explanation](60-design-review-error-checklist/solution/README.md) · [main.go](60-design-review-error-checklist/solution/main.go) · [tests](60-design-review-error-checklist/solution/main_test.go) |

## Additional practice and revision

- [Six capstone projects](projects/README.md) — staged deliverables and failure-injection cases.
- [Study plan and local validation](STUDY-PLAN.md) — six-week pacing and running separate modules.
- [Common anti-patterns and interview scenarios](PITFALLS-AND-SCENARIOS.md) — debug broken chains, typed nil, retry storms and leak prevention.
- [Senior readiness checklist](ASSESSMENT.md) — error policy, API contract, reliability and review rubric.
- [Revision notes](NOTES.md) — blank confidence and evidence tracker; no exercises are marked as completed.
- [Older runnable error examples](../../error-handling/README.md) — preserved without changes.

The code is for study, not a claim that every design is production-ready. The breaker, metric counter and status adapters are intentionally small so the important error semantics remain visible.

