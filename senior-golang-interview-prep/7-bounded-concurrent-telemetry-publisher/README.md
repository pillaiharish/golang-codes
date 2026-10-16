[← Senior Go interview preparation index](../README.md) · [Personal notes](NOTES.md) · [Revision log](../REVISION-LOG.md)

> Study notes adapted from the supplied transcript. Technical examples are preserved for review; code blocks are illustrative unless separately validated.

> **Verification status:** The embedded publisher, test and benchmark snippets have not been assembled and executed as a Go package in this change. Pay particular attention to shutdown deadlines, blocked sinks, and queue accounting when validating the exercise.

# 7 — Bounded Concurrent Telemetry Publisher

## Capstone: design and defend the publisher

This is the capstone module of this learning sequence because it forces you to combine everything:

- scheduler behavior
- channels and backpressure
- synchronization
- context cancellation
- graceful shutdown
- memory awareness
- benchmarking
- production trade-offs

We are going to design a publisher that is close to a real Kafka telemetry pipeline.

---

## 7.1 Problem restatement

You have:

- many producers writing telemetry events into a bounded queue
- a queue capacity of 1,024
- one or more workers sending events to Kafka
- synchronous downstream sends
- under load, the queue fills and about 9% of events are dropped

You need to build a publisher that:

1. defines `TelemetryEvent` and a `Publisher` interface
2. implements a bounded channel-backed queue
3. supports:
   - nonblocking publish
   - optional context-aware blocking publish
4. runs a configurable number of workers
5. tracks counters:
   - accepted
   - queue-full dropped
   - successful sends
   - failed sends
6. supports graceful shutdown with a deadline
7. has unit tests
8. passes the race detector
9. benchmarks throughput at concurrency 1, 4, 16, and 64
10. can be defended architecturally

---

## 7.2 High-level design

The clean design is:

```text
Producers
   │
   ├─ TryPublish()       nonblocking
   ├─ Publish(ctx)       blocking with cancellation
   ▼
Bounded in-memory queue (chan TelemetryEvent, cap 1024)
   │
   ▼
Worker pool
   │
   ├─ Sink.Send(ctx, event)
   ▼
Kafka / fake sink / downstream publisher
```

Important design rule:

> The publisher should expose clear acceptance semantics and explicit failure counters.

That means callers should know whether:

- the event was accepted into the queue
- the event was dropped because the queue was full
- the event was rejected because shutdown started
- the event was eventually sent successfully
- the send failed downstream

This is exactly the kind of observability a senior engineer should build in from the start.

---

## 7.3 Reference implementation

Below is a compact but production-shaped implementation.

It is not the only correct design, but it demonstrates the important senior-level properties:

- safe shutdown
- no send on closed channel
- context-aware publish
- bounded queue
- counters
- worker pool
- testability via a `Sink` interface

---

### 7.3.1 Core types

```go
package telemetry

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrPublisherClosed  = errors.New("telemetry: publisher is closed")
	ErrPublisherClosing = errors.New("telemetry: publisher is closing")
)

type TelemetryEvent struct {
	ID        string
	Tenant    string
	Timestamp time.Time
	Payload   []byte
}

// Sink abstracts the downstream publisher, e.g. Kafka.
type Sink interface {
	Send(ctx context.Context, event TelemetryEvent) error
	Close(ctx context.Context) error
}

// Publisher is the public API we expose to producers.
type Publisher interface {
	TryPublish(event TelemetryEvent) bool
	Publish(ctx context.Context, event TelemetryEvent) error
	Shutdown(ctx context.Context) error
}
```

### Why `Sink` matters

The `Sink` interface is essential for testing.

In production it can wrap Kafka.

In tests it can be:

- instant
- delayed
- failing
- blocking until context cancellation

That allows you to test:

- normal delivery
- backpressure
- queue drops
- shutdown timeouts
- cancellation behavior

without needing a real Kafka cluster.

---

### 7.3.2 Counters

```go
type CounterSnapshot struct {
	Accepted          int64
	DroppedQueueFull  int64
	Sent              int64
	Failed            int64
	RejectedClosed    int64
	AbandonedShutdown int64
}

type Counters struct {
	accepted          atomic.Int64
	droppedQueueFull  atomic.Int64
	sent              atomic.Int64
	failed            atomic.Int64
	rejectedClosed    atomic.Int64
	abandonedShutdown atomic.Int64
}

func (c *Counters) Snapshot() CounterSnapshot {
	return CounterSnapshot{
		Accepted:          c.accepted.Load(),
		DroppedQueueFull:  c.droppedQueueFull.Load(),
		Sent:              c.sent.Load(),
		Failed:            c.failed.Load(),
		RejectedClosed:    c.rejectedClosed.Load(),
		AbandonedShutdown: c.abandonedShutdown.Load(),
	}
}
```

### Why include `AbandonedShutdown`?

Because there is a subtle difference between:

- dropped at ingestion because queue was full
- accepted but never sent because shutdown deadline expired

A senior design separates those two.

If you only expose `dropped`, you will not know whether loss is happening:

- before acceptance
- after acceptance during shutdown
- downstream due to send failure

That distinction matters for SLOs.

---

### 7.3.3 Publisher struct

```go
type Config struct {
	QueueCapacity int
	Workers       int
}

type BoundedPublisher struct {
	cfg     Config
	sink    Sink
	queue   chan TelemetryEvent
	counters Counters

	accepting atomic.Bool
	closing   chan struct{}
	closeOnce sync.Once

	shutdownOnce sync.Once
	shutdownErr  error

	publishMu   sync.RWMutex
	queueClosed bool

	workerCtx    context.Context
	workerCancel context.CancelFunc
	wg           sync.WaitGroup
}

func NewBoundedPublisher(cfg Config, sink Sink) *BoundedPublisher {
	if cfg.QueueCapacity <= 0 {
		cfg.QueueCapacity = 1024
	}
	if cfg.Workers < 0 {
		panic("telemetry: workers cannot be negative")
	}

	workerCtx, workerCancel := context.WithCancel(context.Background())

	p := &BoundedPublisher{
		cfg:          cfg,
		sink:         sink,
		queue:        make(chan TelemetryEvent, cfg.QueueCapacity),
		closing:      make(chan struct{}),
		workerCtx:    workerCtx,
		workerCancel: workerCancel,
	}

	p.accepting.Store(true)

	for i := 0; i < cfg.Workers; i++ {
		p.wg.Add(1)
		go p.worker(p.workerCtx)
	}

	return p
}
```

### Notes

- `Workers == 0` is allowed deliberately for deterministic tests.
- In production, you would normally configure at least one worker.
- The queue is bounded by design.
- The root worker context allows forced cancellation on shutdown timeout.

---

### 7.3.4 Nonblocking publish

```go
func (p *BoundedPublisher) TryPublish(event TelemetryEvent) bool {
	if !p.accepting.Load() {
		p.counters.rejectedClosed.Add(1)
		return false
	}

	p.publishMu.RLock()
	defer p.publishMu.RUnlock()

	if !p.accepting.Load() {
		p.counters.rejectedClosed.Add(1)
		return false
	}

	select {
	case <-p.closing:
		p.counters.rejectedClosed.Add(1)
		return false
	default:
	}

	select {
	case p.queue <- event:
		p.counters.accepted.Add(1)
		return true
	default:
		p.counters.droppedQueueFull.Add(1)
		return false
	}
}
```

### Behavior

- Returns immediately.
- Accepts if queue has space.
- Drops and increments `DroppedQueueFull` if queue is full.
- Rejects if shutdown has started.

This is the right default for an ingest API that must not block indefinitely.

---

### 7.3.5 Context-aware blocking publish

```go
func (p *BoundedPublisher) Publish(ctx context.Context, event TelemetryEvent) error {
	if !p.accepting.Load() {
		p.counters.rejectedClosed.Add(1)
		return ErrPublisherClosed
	}

	p.publishMu.RLock()
	defer p.publishMu.RUnlock()

	if !p.accepting.Load() {
		p.counters.rejectedClosed.Add(1)
		return ErrPublisherClosed
	}

	select {
	case <-p.closing:
		p.counters.rejectedClosed.Add(1)
		return ErrPublisherClosing
	default:
	}

	select {
	case p.queue <- event:
		p.counters.accepted.Add(1)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-p.closing:
		p.counters.rejectedClosed.Add(1)
		return ErrPublisherClosing
	}
}
```

### Behavior

- Blocks until queue space is available, unless:
  - context is cancelled/timed out
  - publisher starts closing
- Useful when the caller can tolerate backpressure.
- Still safe during shutdown.

---

### 7.3.6 Worker loop

```go
func (p *BoundedPublisher) worker(ctx context.Context) {
	defer p.wg.Done()

	for {
		select {
		case event, ok := <-p.queue:
			if !ok {
				return
			}

			if err := p.sink.Send(ctx, event); err != nil {
				p.counters.failed.Add(1)
			} else {
				p.counters.sent.Add(1)
			}

		case <-ctx.Done():
			return
		}
	}
}
```

### Important property

The worker exits when either:

- queue is closed and drained, or
- worker context is cancelled due to shutdown deadline

This gives you both:

- graceful drain
- forced abort if deadline expires

---

### 7.3.7 Graceful shutdown

```go
func (p *BoundedPublisher) Shutdown(ctx context.Context) error {
	p.shutdownOnce.Do(func() {
		p.shutdownErr = p.shutdown(ctx)
	})
	return p.shutdownErr
}

func (p *BoundedPublisher) shutdown(ctx context.Context) error {
	// 1. Stop accepting new work.
	p.accepting.Store(false)
	p.closeOnce.Do(func() {
		close(p.closing)
	})

	// 2. Wait for in-flight publish calls to finish,
	// then close the queue safely.
	p.publishMu.Lock()
	if !p.queueClosed {
		close(p.queue)
		p.queueClosed = true
	}
	p.publishMu.Unlock()

	// 3. Wait for workers to drain, but respect the deadline.
	done := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(done)
	}()

	var waitErr error

	select {
	case <-done:
		// Workers drained normally.
	case <-ctx.Done():
		waitErr = ctx.Err()
		p.workerCancel()
		<-done
	}

	// 4. Count accepted-but-not-sent events left in the queue.
	if n := len(p.queue); n > 0 {
		p.counters.abandonedShutdown.Add(int64(n))
	}

	// 5. Close the downstream sink.
	sinkErr := p.sink.Close(ctx)

	return errors.Join(waitErr, sinkErr)
}
```

### Why this shutdown order matters

This is the part interviewers probe.

#### Step 1: stop accepting new events

If you do not stop producers first, you risk:

- send on closed channel
- panic during shutdown
- indefinite race between ingest and close

#### Step 2: close queue safely

The `publishMu` write lock ensures no publisher is currently inside a send operation when the queue is closed.

This avoids:

```text
panic: send on closed channel
```

#### Step 3: wait for drain with deadline

This gives workers a chance to process already-accepted events.

But it does not wait forever.

If the deadline expires:

- cancel worker context
- abandon remaining work
- count abandoned events

#### Step 4: close the sink

This is critical for Kafka-style clients.

Many Kafka producers have internal batching/buffering. If you exit without closing the producer, you may lose messages that were accepted by the client library but not yet acknowledged by the broker.

---

### 7.3.8 Observability helpers

```go
func (p *BoundedPublisher) Snapshot() CounterSnapshot {
	return p.counters.Snapshot()
}

func (p *BoundedPublisher) QueueLength() int {
	return len(p.queue)
}
```

In production, you would export these as metrics:

- `publisher_accepted_total`
- `publisher_dropped_queue_full_total`
- `publisher_sent_total`
- `publisher_failed_total`
- `publisher_rejected_closed_total`
- `publisher_abandoned_shutdown_total`
- `publisher_queue_length`

For a telemetry system, queue length and drop rate are first-class signals.

---

## 7.4 Unit tests

These tests are deliberately simple but meaningful.

They cover:

- queue-full drop
- blocking publish timeout
- successful delivery
- shutdown rejection
- shutdown deadline abandonment

---

### 7.4.1 Fake sink for tests

```go
package telemetry

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type fakeSink struct {
	mu                  sync.Mutex
	sent                int
	delay               time.Duration
	blockUntilCancelled bool
	closed              atomic.Bool
}

func (f *fakeSink) Send(ctx context.Context, event TelemetryEvent) error {
	if f.blockUntilCancelled {
		<-ctx.Done()
		return ctx.Err()
	}

	if f.delay > 0 {
		select {
		case <-time.After(f.delay):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	f.mu.Lock()
	f.sent++
	f.mu.Unlock()

	return nil
}

func (f *fakeSink) Close(ctx context.Context) error {
	f.closed.Store(true)
	return nil
}

func (f *fakeSink) Sent() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.sent
}
```

---

### 7.4.2 Test: nonblocking publish drops when full

```go
package telemetry

import (
	"context"
	"testing"
)

func TestTryPublishDropsWhenQueueIsFull(t *testing.T) {
	sink := &fakeSink{}

	p := NewBoundedPublisher(Config{
		QueueCapacity: 2,
		Workers:       0,
	}, sink)

	event := TelemetryEvent{ID: "e1"}

	if !p.TryPublish(event) {
		t.Fatal("expected first publish to succeed")
	}
	if !p.TryPublish(event) {
		t.Fatal("expected second publish to succeed")
	}
	if p.TryPublish(event) {
		t.Fatal("expected third publish to drop")
	}

	snap := p.Snapshot()

	if snap.Accepted != 2 {
		t.Fatalf("expected accepted=2, got %d", snap.Accepted)
	}
	if snap.DroppedQueueFull != 1 {
		t.Fatalf("expected dropped=1, got %d", snap.DroppedQueueFull)
	}

	if err := p.Shutdown(context.Background()); err != nil {
		t.Fatalf("shutdown failed: %v", err)
	}
}
```

### Why `Workers: 0` here?

Because we want a deterministic full queue.

If workers are running, they may drain events and make the test timing-dependent.

---

### 7.4.3 Test: blocking publish honors context timeout

```go
package telemetry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestBlockingPublishHonorsContextTimeout(t *testing.T) {
	sink := &fakeSink{}

	p := NewBoundedPublisher(Config{
		QueueCapacity: 1,
		Workers:       0,
	}, sink)

	event := TelemetryEvent{ID: "e1"}

	if !p.TryPublish(event) {
		t.Fatal("expected first publish to succeed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := p.Publish(ctx, event)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected DeadlineExceeded, got %v", err)
	}

	if err := p.Shutdown(context.Background()); err != nil {
		t.Fatalf("shutdown failed: %v", err)
	}
}
```

---

### 7.4.4 Test: shutdown rejects new publishes

```go
package telemetry

import (
	"context"
	"errors"
	"testing"
)

func TestShutdownRejectsNewPublishes(t *testing.T) {
	sink := &fakeSink{}

	p := NewBoundedPublisher(Config{
		QueueCapacity: 4,
		Workers:       1,
	}, sink)

	if err := p.Shutdown(context.Background()); err != nil {
		t.Fatalf("shutdown failed: %v", err)
	}

	err := p.Publish(context.Background(), TelemetryEvent{ID: "late"})
	if !errors.Is(err, ErrPublisherClosed) {
		t.Fatalf("expected ErrPublisherClosed, got %v", err)
	}

	if p.TryPublish(TelemetryEvent{ID: "late"}) {
		t.Fatal("expected TryPublish to fail after shutdown")
	}
}
```

---

### 7.4.5 Test: events are delivered through workers

```go
package telemetry

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestPublishDeliversAllEvents(t *testing.T) {
	sink := &fakeSink{}

	p := NewBoundedPublisher(Config{
		QueueCapacity: 64,
		Workers:       4,
	}, sink)

	const total = 500

	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			event := TelemetryEvent{
				ID: strconv.Itoa(i),
			}

			if err := p.Publish(context.Background(), event); err != nil {
				t.Errorf("publish failed: %v", err)
			}
		}(i)
	}

	wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := p.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown failed: %v", err)
	}

	snap := p.Snapshot()

	if snap.Accepted != total {
		t.Fatalf("expected accepted=%d, got %d", total, snap.Accepted)
	}
	if snap.Sent != total {
		t.Fatalf("expected sent=%d, got %d", total, snap.Sent)
	}
	if snap.Failed != 0 {
		t.Fatalf("expected failed=0, got %d", snap.Failed)
	}
	if snap.DroppedQueueFull != 0 {
		t.Fatalf("expected dropped=0, got %d", snap.DroppedQueueFull)
	}

	if sink.Sent() != total {
		t.Fatalf("expected sink sent=%d, got %d", total, sink.Sent())
	}
}
```

This is the key correctness test.

Run it with:

```bash
go test -race ./...
```

---

### 7.4.6 Test: shutdown deadline abandons stuck work

```go
package telemetry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestShutdownDeadlineAbandonsStuckWork(t *testing.T) {
	sink := &fakeSink{
		blockUntilCancelled: true,
	}

	p := NewBoundedPublisher(Config{
		QueueCapacity: 10,
		Workers:       2,
	}, sink)

	for i := 0; i < 10; i++ {
		if !p.TryPublish(TelemetryEvent{ID: "stuck"}) {
			t.Fatalf("expected publish %d to succeed", i)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	err := p.Shutdown(ctx)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected DeadlineExceeded, got %v", err)
	}

	snap := p.Snapshot()

	if snap.Failed+snap.AbandonedShutdown == 0 {
		t.Fatal("expected some failed or abandoned events during forced shutdown")
	}
}
```

### Why this test matters

It proves the publisher does not hang forever when the downstream sink is stuck.

This is exactly the kind of behavior you need during pod termination.

---

## 7.5 Running tests with the race detector

Use:

```bash
go test -race -count=5 ./...
```

For more stress:

```bash
go test -race -count=20 -cpu=1,4,8 ./...
```

A senior-level answer should mention:

> The race detector only proves that exercised interleavings are race-free. It is necessary but not sufficient.

---

## 7.6 Benchmarking throughput at concurrency 1, 4, 16, 64

The benchmark below measures producer throughput through the bounded publisher using a fast fake sink.

```go
package telemetry

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func BenchmarkBoundedPublisher(b *testing.B) {
	concurrencies := []int{1, 4, 16, 64}

	for _, conc := range concurrencies {
		b.Run(fmt.Sprintf("producers=%d", conc), func(b *testing.B) {
			sink := &fakeSink{}

			p := NewBoundedPublisher(Config{
				QueueCapacity: 1024,
				Workers:       16,
			}, sink)

			event := TelemetryEvent{
				ID:      "bench",
				Tenant:  "tenant-1",
				Payload: make([]byte, 256),
			}

			b.SetParallelism(conc)
			b.ReportAllocs()
			b.SetBytes(int64(len(event.Payload)))

			b.RunParallel(func(pb *testing.PB) {
				ctx := context.Background()

				for pb.Next() {
					if err := p.Publish(ctx, event); err != nil {
						b.Error(err)
						return
					}
				}
			})

			b.StopTimer()

			shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := p.Shutdown(shutdownCtx); err != nil {
				b.Error(err)
			}
		})
	}
}
```

Run:

```bash
go test -run '^$' -bench=. -benchmem ./...
```

---

## 7.7 What the benchmark is actually measuring

This benchmark measures:

- producer-side concurrency
- queue acceptance behavior
- worker drain behavior
- synchronization overhead
- channel scheduling overhead
- shutdown drain cost

It does **not** measure real Kafka behavior unless the sink is a real Kafka client.

That is an important distinction.

---

## 7.8 Expected benchmark behavior

With a fake instant sink:

- throughput should increase as producer concurrency increases, up to a point
- eventually contention appears:
  - channel queueing
  - worker saturation
  - scheduler overhead
  - lock overhead from publish/close protection
- if workers are too few, producers may block due to backpressure

With a real Kafka sink:

- send latency dominates
- throughput depends heavily on:
  - batching
  - acks configuration
  - partition count
  - broker capacity
  - network latency
  - compression
  - serialization cost

This leads directly into the architecture challenges.

---

## 7.9 Architectural challenge questions

These are the questions that separate a senior answer from a basic implementation.

---

## Challenge 1: If Kafka itself is the bottleneck, will adding more goroutines really help?

### Short answer

Usually no.

### Senior answer

Adding more goroutines helps only if the bottleneck is local:

- insufficient worker concurrency
- excessive serialization
- local queueing
- CPU-bound encoding/compression
- underutilized network connections

If the bottleneck is Kafka itself, adding more goroutines may not improve throughput. It can even make things worse by increasing:

- broker load
- producer memory usage
- request contention
- batching inefficiency
- tail latency

You need evidence:

- Kafka producer send latency
- broker CPU/network
- partition count
- producer batch size
- request queueing
- local worker utilization

If Kafka is saturated, the correct fixes are usually:

- increase partitions, if ordering allows
- increase batch size
- enable compression
- tune `acks`
- add brokers
- reduce payload size
- sample/drop lower-value telemetry
- move to asynchronous publishing with durable buffering

A senior engineer does not say “add workers.” They say:

> “I would profile the pipeline and determine whether the constraint is local concurrency, serialization, network, or broker capacity.”

---

## Challenge 2: What guarantees can an in-memory queue provide if the pod crashes?

### Short answer

None.

### Senior answer

An in-memory queue provides no durability.

If the pod crashes after accepting events but before sending them to Kafka, those events are lost.

This means your delivery guarantee depends on when you acknowledge the event.

### If you acknowledge at enqueue time

You have effectively promised:

> “We accepted the event.”

But you have not promised:

> “The event will be delivered.”

If the pod crashes, accepted events can disappear.

This may be acceptable for lossy telemetry, but it must be explicit.

### If you need stronger guarantees

You need one of:

1. acknowledge only after Kafka success
2. persist events to disk before acknowledging
3. use an outbox/WAL pattern
4. rely on client-side retry from the original producer

This is not a coding detail. It is a product-level correctness decision.

---

## Challenge 3: When is dropping 9% of observability data acceptable, and when is it a correctness violation?

### Short answer

It depends on the class of data and the SLO.

### Acceptable when

The data is:

- high-volume telemetry
- sampled traces
- non-critical metrics
- debug-level events
- redundant signals
- approximate analytics

For example:

- CPU samples
- request latency histograms
- trace spans where sampling is expected
- diagnostic logs

In those cases, 9% loss may be acceptable if:

- it is documented
- it is monitored
- it does not break alerting
- it does not hide incidents

### Not acceptable when

The data is:

- billing events
- audit logs
- security events
- compliance records
- incident-critical traces
- usage metering
- contractual delivery data

In those cases, dropping 9% is not “lossy telemetry.”

It is a correctness violation.

### Senior framing

> “Loss is acceptable only when the system is explicitly designed as best-effort and the business impact of loss is understood. If the data drives billing, security, or compliance, loss is not an operational metric — it is a defect.”

---

## Challenge 4: What changes if we need at-least-once delivery instead of best-effort delivery?

This changes the design substantially.

### Best-effort design

Current design:

- accept into memory if possible
- drop if full
- send asynchronously
- loss possible

This optimizes for:

- availability
- low latency
- bounded memory

### At-least-once design

At-least-once requires:

1. no silent loss after acceptance
2. retries on transient failure
3. idempotency or deduplication downstream
4. durable buffering or ACK-after-send
5. explicit backpressure when buffer is full

Possible designs:

#### Option A: Acknowledge after Kafka success

```text
Producer -> Publisher -> Kafka -> ACK -> caller ACK
```

Pros:

- stronger guarantee

Cons:

- higher latency
- caller sees downstream failures
- caller may retry, creating duplicates

#### Option B: Durable outbox/WAL

```text
Producer -> local durable outbox -> background publisher -> Kafka
```

Pros:

- survives pod crash
- decouples caller latency from Kafka
- replay possible

Cons:

- more complexity
- disk I/O
- replay/duplicate handling
- ordering concerns

#### Option C: Client-side retry

The original producer retries until the service confirms durable acceptance.

Pros:

- moves durability to edge

Cons:

- requires client cooperation
- duplicates possible
- backpressure still needed

### Important consequence

At-least-once usually implies duplicates.

So consumers must be:

- idempotent
- deduplicating
- tolerant of retries

If downstream cannot handle duplicates, at-least-once may still be operationally unsafe.

---

## Challenge 5: Which measurements would prove that the redesigned publisher is better?

This is the most important senior question.

You prove improvement with evidence, not adjectives.

### 1. Throughput

Measure:

- events/sec accepted
- events/sec sent
- bytes/sec

At different producer concurrency levels:

- 1
- 4
- 16
- 64

### 2. Loss behavior

Measure:

- accepted total
- dropped queue-full total
- failed sends total
- abandoned during shutdown total

You want to know not just how much was lost, but **where** it was lost.

### 3. Latency

Measure:

- publish latency p50/p95/p99
- queue wait time
- sink send latency p95/p99
- shutdown drain duration

### 4. Saturation

Measure:

- CPU usage
- goroutine count
- queue depth over time
- GC CPU
- allocation rate

### 5. Stability under failure

Test:

- Kafka slowdown
- Kafka outage
- pod shutdown during load
- burst traffic
- context cancellation
- deadline expiration

### 6. Correctness

Verify:

- accepted = sent + failed + abandoned, within accounting rules
- no panic during shutdown
- no send on closed channel
- no goroutine leaks
- race detector clean

### Strong interview answer

> “I would compare before/after benchmarks at multiple concurrency levels, production p99 publish latency, queue depth, drop rate, shutdown loss, and CPU/GC overhead. The redesigned publisher is better only if it improves the relevant SLOs without introducing unacceptable loss or resource pressure.”

---

## 7.10 Possible extensions to discuss in an interview

Once you have the base implementation, a senior interviewer may push you on improvements.

## Extension 1: Batching

Instead of sending one event at a time:

```go
sink.Send(ctx, event)
```

send batches:

```go
sink.SendBatch(ctx, batch)
```

This often improves Kafka throughput significantly.

Trade-offs:

- higher latency
- more complex shutdown
- more complex partial-failure handling

---

## Extension 2: Drop policy

Currently, when the queue is full, new events are dropped.

Other policies:

- drop newest
- drop oldest
- drop by tenant
- drop low-priority events
- spill to disk

Each has different correctness implications.

---

## Extension 3: Priority queues

You may want:

- error events
- security events
- incident-related telemetry

to have priority over normal metrics.

This requires more than one channel or a custom scheduler.

---

## Extension 4: Multi-tenant fairness

If one tenant floods the queue, other tenants may starve.

Possible fixes:

- per-tenant quotas
- weighted fairness
- separate queues
- token buckets

This is very relevant for observability platforms.

---

## Extension 5: Disk spillover

If memory queue is full, spill to disk instead of dropping.

Trade-offs:

- better durability
- more complexity
- disk latency
- replay ordering
- crash recovery

This is often the bridge between best-effort and at-least-once.

---

## 7.11 How to defend this implementation in an interview

If asked to defend the design, I would structure the answer like this:

## 1. Correctness

> “The publisher avoids send-on-closed-channel panics by stopping acceptance and using a publish lock during shutdown. It tracks accepted, dropped, sent, failed, rejected, and abandoned events separately, so loss is observable and attributable.”

## 2. Lifecycle

> “Shutdown is deadline-bound. It first rejects new work, closes the queue safely, waits for workers to drain, cancels remaining workers if the deadline expires, and then closes the sink.”

## 3. Backpressure

> “The bounded queue gives explicit backpressure. `TryPublish` provides a nonblocking lossy path for latency-sensitive APIs, while `Publish(ctx)` provides a blocking path for callers that can tolerate backpressure.”

## 4. Observability

> “The counters and queue length are exportable metrics. In production, I would alert on drop rate, failed sends, queue depth, and shutdown abandonment.”

## 5. Verification

> “I would verify the design with unit tests, race detector runs, benchmarks at multiple concurrency levels, and failure injection tests such as slow sinks, cancelled contexts, and forced shutdown deadlines.”

That answer is strong because it covers:

- semantics
- safety
- operations
- verification

---

## 7.12 Common weaknesses in candidate implementations

Interviewers will look for these.

## Weakness 1: Closing the channel without coordinating producers

This causes:

```text
panic: send on closed channel
```

Senior fix:

- stop producers first, or
- guard publish/close with synchronization, or
- avoid closing data channels while producers are active

---

## Weakness 2: No shutdown deadline

This causes shutdown to hang if Kafka is slow.

Senior fix:

- bounded drain
- context cancellation
- forced exit path

---

## Weakness 3: No distinction between dropped and failed events

This makes loss analysis impossible.

Senior fix:

- separate counters
- separate metrics
- separate alerts

---

## Weakness 4: Assuming more workers always help

Senior fix:

- benchmark
- profile
- identify bottleneck
- tune workers based on evidence

---

## Weakness 5: Treating in-memory queue as durable

Senior fix:

- state the guarantee explicitly
- document loss window
- add durable buffering if needed

---

## 7.13 What I would say if asked, “Is this production-ready?”

A good senior answer is honest:

> “This is a solid bounded publisher for best-effort telemetry. It has safe shutdown, backpressure, observability, and test coverage. However, it is not durable across pod crashes. If we need at-least-once delivery, we must change the acknowledgment model and add durable buffering or publish-after-success semantics. I would also add batching, metrics, and load tests before considering it production-ready for high-throughput Kafka traffic.”

That answer is better than claiming the first implementation is perfect.

---

## 7.14 Suggested commands for this module

```bash
# run tests
go test ./...

# run tests with race detector
go test -race ./...

# stress tests
go test -race -count=10 ./...

# run benchmarks
go test -run '^$' -bench=. -benchmem ./...

# run only publisher benchmarks
go test -run '^$' -bench=BoundedPublisher -benchmem ./...

# inspect escape analysis if optimizing allocations
go build -gcflags="-m" ./...

# CPU profile during benchmark
go test -run '^$' -bench=. -cpuprofile cpu.prof ./...

# memory profile during benchmark
go test -run '^$' -bench=. -memprofile mem.prof ./...

# open profiles
go tool pprof -http=:8080 cpu.prof
go tool pprof -http=:8080 mem.prof
```

---

---

## Revision checkpoint

Try answering from memory *before* opening the relevant section. Mark the boxes only after you can explain or demonstrate the result.

- [ ] Explain queue acceptance vs downstream delivery and the counters for each state.
- [ ] Review the code for shutdown races, forced cancellation, and hanging sinks.
- [ ] Justify nonblocking and context-aware blocking publication policies.
- [ ] Run tests, race detection, and benchmarks at concurrency 1, 4, 16, and 64.
- [ ] Defend the transition from best-effort in-memory queuing to at-least-once delivery.

**After review:** record your confidence, uncertainties, and evidence in [NOTES.md](NOTES.md) and add one entry to the [revision log](../REVISION-LOG.md).
