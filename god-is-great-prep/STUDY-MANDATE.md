# The Go Study Mandate — Layer 1 and Layer 2

[Home](README.md)

This is the **quality bar for every new lesson** in this folder. A heading or pasted code block does not count as understanding.

## Every topic must provide

1. **Definition and when to use it.** Explain in plain English, and compare the nearest alternative.
2. **One tiny valid Go example.** Show input, expected output/return, and explain the important lines.
3. **Operations.** Where meaningful: create, read, update, insert, delete, traverse, search, copy, compare, sort/reorder, reverse.
4. **Costs and ownership.** State time/space complexity and whether changes affect original data. Explicitly say “not applicable” for operations that do not belong to the type.
5. **At least two boundaries.** Empty/nil, duplicate, out-of-bounds, Unicode, fixed-size, missing keys or invalid input as applicable.
6. **Proof.** Compile/run a sample and test one success and one edge case.
7. **Memory prompt.** Ask yourself to reproduce the example with notes closed, explain the result, and name a failure mode.
8. **Revisit.** Recall next day, days 3/7/14/30; if you fail, revise the concept and repeat.

## Curriculum boundaries

| Layer | Learn here | Save for later |
|---|---|---|
| 1 | Syntax, functions, variables, flow control, simple errors, basic toolchain, *syntax only* for common data types | Internal representations, benchmarks, large codebases |
| 2 | Go collections, composition, pointers, methods, primitive structures, operations, complexity, basic tests | Deep traps (Layer 3), reusable operation catalog (Layer 4), goroutines/GC/Kafka/K8s (Layer 5+) |
| 3 | Edge cases, language semantics, aliasing, nil-interface, method sets, race/panic scenarios | Large production design |
| 4 | Reusable operation templates; drill insertion/deletion/reversals until fluent | Random interview grinding |
| 5+ | Runtime, concurrency, production backend Go, algorithms, distributed systems, design | — |

## Go-specific correctness checklist

- Use `gofmt`; avoid intentionally unbuildable code except clearly marked **wrong examples**.
- Declare package/imports for runnable programs. Independent snippets can live in docs, but the tested package must compile.
- Handle missing keys with comma-ok when zero is ambiguous; don't assume map iteration order.
- For slices, distinguish `len`, `cap` and shared backing storage. `append` must use its returned slice.
- Avoid stale references when removing nodes; define whether an index is valid and what the API does if it is not.
- Beware slice queues that retain consumed elements; a head-index queue needs compaction/reset or a ring buffer.
- For a rune demo, note **runes are Unicode code points, not always grapheme clusters**.
- Do not state maps are thread-safe or that slice deletion changes storage capacity.
- Explicitly distinguish arrays copied by value from maps and slices that can refer to shared state.
- Describe complexity in the correct implementation model; don't promise worst-case O(1) for hash maps.
- Keep basic `interface` satisfaction separate from deep method-set and nil-interface traps.

## Exit criteria

**Layer 1:** fresh CLI example + basic helper test + a two-minute verbal explanation without notes.

**Layer 2:** each topic has a worked snippet; exercises run; you can implement stack, queue, set and linked-list primitives, including empty/singleton/duplicate cases. Explain operation complexity and why a chosen type fits.

**Don't fake completion:** track `Learning → Can reproduce → Can debug → Recalled at 7d/30d` per module. A passed test now is not evidence of retention next month.
