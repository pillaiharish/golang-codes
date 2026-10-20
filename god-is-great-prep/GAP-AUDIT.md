# Layer 1 / Layer 2 gap audit

[Prep home](README.md) · [Layer 1](layer1/README.md) · [Layer 2](layer2/README.md)

Assessment of the **existing two long README documents** before this PR. This is a syllabus coverage audit, not a claim that every old example is correct or tested.

| Area | Existing baseline | What this PR changes |
|---|---|---|
| Layer 1 18 beginner topics | Covered extensively in one long README | Keeps original verbatim as `reference.md`, adds navigable 18-topic index and exit checkpoint |
| Layer 2 arrays/slices/maps/structs | Examples and narrative present | Adds operation checklist, ownership/complexity notes, run-tested reference APIs |
| Layer 2 pointers/methods/interfaces | Basic syntax and examples present | Adds fresh runnable examples and tests; reserves method-set traps for Layer 3 |
| Layer 2 strings/runes | Covered, including `strings.Builder` | Adds code-point reversal and Unicode caveat |
| Layer 2 singly/doubly linked lists | Definitions, insertion/traversal were present | Adds executable delete/reverse, head/tail and move-to-front examples |
| Layer 2 stack/queue/set | Conceptual snippets were present | Adds reusable types and tests including empty and retention/reset behavior |
| Layer 2 multidimensional slices, sorting/search | Covered as introductions | Adds jagged matrix validation, transpose and safe clone-before-sort |
| Multiset and frequency maps | Not a distinct study unit | Adds one as an explicit module |
| Testing and time/space complexity | Mentioned, but not a consistent exit gate | Adds per-topic test/complexity/recall requirements and a self-contained Go module |
| Permanent retention workflow | Not codified in folder root | Adds study mandate, 1/3/7/14/30-day recall schedule, module checkpoints |
| Discoverability | No folder-root index | Adds curriculum home and link from repository root |

## Corrections / qualifications

- **Slice-based queue:** Removing the first element by `q = q[1:]` is O(1); it does not itself copy the remaining elements. Risks are backing-storage retention and repeated growth; use compaction, reset or a ring buffer when appropriate.
- **Map iteration:** order is unspecified and must never be depended upon (not an invariant “random order”).
- **Copy semantics:** arrays and structs copy values, but copied fields may still reference nested slices/maps/pointers; `copy` is not recursive.
- **Unicode:** a rune is a Unicode code point; reversing runes is not always reversing human-visible grapheme clusters.
- **Pointer increment syntax:** `*p++` is parsed as `(*p)++`, but explicit `(*p)++` is clearer in teaching materials.
- **Complexity:** map access is O(1) *on average*, queue operations can be *amortized* O(1), and linked-list tail insertion is only O(1) when a tail is held.

## Explicitly deferred

Not forced into beginner material: channels, goroutine scheduler, `context`, GC/escape analysis, race detector, Kafka/K8s/Redis/ClickHouse, REST servers, LRU, advanced heaps/trees/graphs, DSA pattern interviews, system design. These belong to Layers 3–8 under the larger plan.

## Suggested next expansion after this PR

For the next module, implement one operation per file in Layer 4 (insert/delete/move/reverse) using tests; then create small traps with expected wrong output in Layer 3. Link each exercise to its Layer 2 prerequisite instead of copying full lessons again.
