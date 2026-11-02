# Q28. Build a small pub/sub system — worked solution

[Back to Q28 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q25: tee](../../25-tee-broadcast-two-outputs/README.md) · [Revision notes](../../NOTES.md)

**Level 5 — core concurrency patterns.**

## Task

Implement `Subscribe(topic)`, `Publish(topic, message)`, and `Unsubscribe`. Every active subscriber of a topic should receive each accepted message, subject to a documented slow-subscriber policy.

## Reference approach

- Own a mapping `topic -> subscriber set`; protect the mapping with a mutex or serialize all operations through one coordinator goroutine.
- Use a **single ownership policy** for subscriber channel closing. Do not close from one goroutine while another may still send.
- Decide whether a slow subscriber blocks all publishers, gets bounded buffering, drops with metrics, or is disconnected.
- Return a cancellation/unsubscribe function for each subscription and define what happens to in-flight publishes.
- Test concurrent subscribe/unsubscribe/publish and ensure no send-on-closed panic or leaked subscriber goroutine.

Simplified API sketch (not a complete implementation):

```go
type Message struct { Topic, Body string }
type Subscription struct { C <-chan Message; Cancel func() }
type Bus interface {
    Subscribe(topic string) Subscription
    Publish(ctx context.Context, m Message) error
}
```

The hardest part is not the map: it is delivery semantics, backpressure and safe channel closure under concurrent subscription changes.

## Scenario Q&A

| Question | Answer |
|---|---|
| Can multiple subscribers read one shared channel and each see every message? | No. They compete for messages. Each subscriber needs its own delivery path or broker-managed fan-out. |
| What if a subscriber is slow? | Define a bound and a policy: block, drop, disconnect, or replay later from durable storage. |
| What if unsubscribe races with publish? | Use shared ownership/serialization so no publisher can send to a channel after it closes. |
| Does this in-memory bus survive a process crash? | No. Durable pub/sub needs persistence and replay semantics. |

**Recall:** Compare a local pub/sub with Kafka topics, partitions, consumer groups, at-most-once versus at-least-once, and delivery acknowledgment.

Reference: [Go memory model](https://go.dev/ref/mem).
