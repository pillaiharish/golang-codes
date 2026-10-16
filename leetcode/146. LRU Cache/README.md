# 146 — LRU Cache

## Before you solve: prerequisite basics

Study these in sequence and type the tiny Go example in each before reopening this solution:

1. [03 maps sets frequency and grouping](../0-problem-solving-building-blocks/03-maps-sets-frequency-and-grouping/README.md)
2. [11 structs pointers methods and linked lists](../0-problem-solving-building-blocks/11-structs-pointers-methods-and-linked-lists/README.md)
3. [16 stateful simulation and sequential comparison](../0-problem-solving-building-blocks/16-stateful-simulation-and-sequential-comparison/README.md)
4. [17 complexity amortization and input contracts](../0-problem-solving-building-blocks/17-complexity-amortization-and-input-contracts/README.md)

[Full mapping for all 23 LeetCode problems](../0-problem-solving-building-blocks/README.md#b-problem--exact-prerequisites-all-23-problem-folders) · [Practice log](../REVISION-LOG.md)

[← Index](../README.md) · [Building block: map + doubly linked list](../0-problem-solving-building-blocks/README.md#5-map--doubly-linked-list--fast-membership--fast-reorder) · [Go solution](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/lru-cache/)

## Recognition

**Clue:** “Get and Put in O(1) average time; evict least recently used entry when full.” A map alone finds an entry, but **cannot tell which one is least recently used**. A list alone records recency, but key lookup is O(n).

**Mental trigger:** **Map locates; doubly linked list orders.**

## Build the picture first

```text
HEAD <-> most recently used <-> ... <-> least recently used <-> TAIL
```

`head` and `tail` are **dummy/sentinel nodes**. They never represent data. Node after head is MRU; node before tail is LRU. Map stores `key -> *Node`.

## Invariant (the real solution)

Every real node appears **exactly once** between the sentinels and has **exactly one corresponding map entry**. Every Get hit, update, or insertion puts its node right after head. Eviction always removes tail.prev.

## Memorize four operations

1. **Remove node:** set `node.prev.next=node.next`, `node.next.prev=node.prev`.
2. **Add to head:** wire `node.prev=head`, `node.next=head.next`, then both neighboring pointers.
3. **Get:** missing → -1; found → detach, move to head, return value.
4. **Put:** existing → update + move to head; new → create + insert into map/list; if size exceeds capacity → detach tail.prev + delete map key.

### Dry run (capacity=2)

| Call | MRU → LRU | Return / effect |
|---|---|---|
| Put(1,1) | 1 | insert |
| Put(2,2) | 2,1 | insert |
| Get(1) | 1,2 | return 1 and move 1 |
| Put(3,3) | 3,1 | evict 2 |
| Get(2) | 3,1 | return -1 |
| Put(4,4) | 4,3 | evict 1 |

## Why the double links?

Given a pointer to a middle node, a **singly** linked list cannot find its previous node in O(1). The doubly linked list can detach a known node by rewiring two neighboring links, without traversing.

The dummy head/tail avoid special cases for zero, one, or many data nodes.

## Complexity and constraints

- Get and Put: expected O(1) from map operations, list updates O(1).
- O(capacity) auxiliary space for map entries and real nodes.
- This implementation treats `capacity <= 0` as zero and stores nothing (a defensive extension).
- **Not concurrency-safe.** If called across goroutines, choose explicit locking or another design; `map` and list pointer writes would race.
- Constructor returns a value containing pointers; don't copy active caches and mutate the copies as if they were independent objects.

## Common bugs

- Updating an existing key without moving its node to MRU.
- Rewiring only one direction of a doubly linked list.
- Forgetting to delete the evicted key from the map.
- Letting sentinel nodes be evicted when capacity is zero.
- Mistaking map lookup order for recency (Go maps are unordered).

## 30-second spoken answer

> I combine a map from key to list node with a doubly linked list ordered most-recent to least-recent. Every successful Get and Put moves a node to the front in constant time. If an insertion exceeds capacity, I remove the node before the dummy tail and delete its map entry. The map handles lookup, the list handles recency, so operations are expected O(1).

## Cold recall

1. Draw pointer rewiring for both `remove` and `addToHead`.
2. What invariant joins the map to the list?
3. Walk through Put(1), Put(2), Get(1), Put(3) for capacity 2.
4. What changes for capacity 0, updating an existing key, or concurrent use?

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—`.
