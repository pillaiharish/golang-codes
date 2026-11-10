# Round 2 — Technical Screen / Assessment

**Reference only · company-neutral**

The second stage may be an **online assessment, live coding session, technical phone screen, or a combination**. For a mid-level-to-senior backend/platform role, prepare algorithms, Go implementation quality, Kubernetes fundamentals, distributed systems, and design judgment. These are **representative study prompts, not verified questions from any specific employer**.

## A. Online assessment: algorithms and problem-solving

Possible areas: arrays, strings, hash maps, stacks and queues, linked lists, trees, graphs, sliding windows, two pointers, binary search, sorting, recursion, basic dynamic programming, debugging, and time/space complexity. Practice easy-to-medium problems first; be able to reason aloud about edge cases and costs.

### Topic-based problem bank

| Area | Problems to practice |
|---|---|
| Arrays and strings | Two Sum; Merge Intervals; Longest Substring Without Repeating Characters; Product of Array Except Self; Rotate Array; Maximum Subarray; Group Anagrams |
| Hash maps / frequency | Top K Frequent Elements; Valid Anagram; First Non-Repeating Character; Frequency Sort |
| Sliding window | Minimum Size Subarray Sum; Longest Substring Without Repeating Characters; Sliding Window Maximum |
| Stacks and queues | Valid Parentheses; Min Stack; Implement Queue Using Stacks; Next Greater Element |
| Trees | Binary Tree Level Order Traversal; Lowest Common Ancestor; Validate Binary Search Tree; Maximum Depth of Binary Tree; Serialize/Deserialize Binary Tree |
| Graphs | Number of Islands; Clone Graph; Course Schedule; Network Delay Time; Shortest Path in an Undirected Graph |
| Practical concurrency | Worker pool; rate limiter; thread-safe cache; LRU cache; concurrent job scheduler; fan-in/fan-out; concurrent API result merge |

Graph traversal, shortest paths, and cycle detection are useful for reasoning about network topologies and task dependencies. Practice explaining the algorithm, complexity, and failure cases without relying on memorized code.

## B. Live coding: practical design prompts

Interviewers may assess **correctness, clean APIs, concurrency, cancellation, tests, edge cases, and trade-offs** rather than only whether a solution runs.

### 1. LRU cache

- Implement `Get` and `Put` with a hash map and doubly linked list; justify `O(1)` average operations.
- Discuss capacity zero, replacement of an existing key, eviction, and a thread-safe extension.

### 2. Rate limiter

- Compare token bucket, fixed window, and sliding window approaches.
- Address monotonic time, bursts, concurrent callers, and per-key fairness.
- Explain what must change when enforcement is distributed across processes.

### 3. Concurrent URL fetcher / web crawler

- Bound concurrency, pass context deadlines, collect per-request errors, and avoid leaks.
- Explain ownership of channel closure, result ordering, and shutdown.

### 4. Merge intervals / meeting scheduler

- Sort and merge overlapping intervals; test empty input, touching boundaries, and unsorted data.

### 5. Log aggregator / parser

- Count errors per minute or return top-N error types from a stream.
- Consider time buckets, malformed input, memory bounds, and streaming updates.

### 6. Task scheduler

- Schedule dependent tasks using a directed acyclic graph; detect cycles.
- Explain topological ordering and what happens when a dependency fails.

### 7. Distributed job queue simulation

- Run multiple workers, retry failed tasks, and support graceful shutdown.
- Discuss idempotency, duplicate delivery, bounded retries, delayed backoff, and dead-letter queues.

**Practice method:** Write a small API, tests for success/failure/cancellation, run the race detector where relevant, then explain one production-scale limitation.

## C. Go-specific technical questions

### Goroutines, channels, and `select`

1. What is a goroutine, and how does it differ from an OS thread?
2. How do buffered and unbuffered channels differ?
3. What causes a channel deadlock, and how would you debug it?
4. How do goroutine leaks occur, and how do you prevent them?
5. How does `select` behave when multiple cases are ready?
6. How do you implement timeouts and cancellation with `select` and `context`?

### Context, synchronization, and safe sharing

7. Why pass `context.Context` through a request call chain?
8. What is the difference between a deadline and explicit cancellation?
9. When should you use `sync.Mutex` instead of channels?
10. How does `sync.WaitGroup` coordinate goroutines? What mistakes cause misuse?
11. When is `sync.RWMutex` worthwhile, and when can it hurt performance?
12. What constitutes a data race? How would `go test -race` help?

### Interfaces and error handling

13. How do interfaces and method sets work? What is the typed-`nil` trap?
14. Explain type assertions, type switches, and `any`.
15. What is the `error` interface? When do you wrap with `%w`?
16. How do `errors.Is`, `errors.As`, and `errors.Join` differ?
17. When should code use `panic` and `recover` rather than returning errors?

### Slices, maps, and performance

18. What does a slice header contain? What happens on `append`?
19. Why must ordinary maps not be read and written concurrently without synchronization?
20. Compare pointer and value receivers; explain escape analysis at a high level.
21. What causes excessive allocations or garbage collection pressure?
22. How would you inspect CPU, heap, and goroutine profiles using `pprof`?

Related permanent references: [Go revision modules](../senior-golang-interview-prep/README.md), [concurrency tasks](../senior-golang-interview-prep/concurrency/README.md), and [error handling exercises](../error-handling/README.md).

## D. Kubernetes fundamentals and troubleshooting

### Core question checklist

1. What happens when you run `kubectl apply -f deployment.yaml`?
2. Explain the control plane and the responsibilities of the API server, scheduler, and controller manager.
3. What does etcd store and why is it operationally important?
4. How is a node chosen for an unscheduled Pod?
5. What is a controller, and how does reconciliation work?
6. When would you choose a Deployment, StatefulSet, or DaemonSet?
7. Compare ClusterIP, NodePort, and LoadBalancer Services.
8. What does Ingress do, and what component implements it?
9. How do NetworkPolicies regulate traffic, and what must the network plugin support?
10. What is RBAC? How would you diagnose a forbidden API request?
11. How do liveness, readiness, and startup probes differ?
12. Compare ConfigMaps and Secrets; what protections do Secrets require?
13. How do you perform low-downtime rolling deployments?
14. How would you troubleshoot `CrashLoopBackOff`?
15. How would you troubleshoot `ImagePullBackOff`?
16. What are CRDs and custom resources?
17. What does a Kubernetes operator add beyond a workload deployment?
18. What does a reconciliation loop compare, and why must it be idempotent?
19. How do you safely roll out control-plane and worker-node upgrades?
20. How would you isolate tenants in a shared cluster?

### Deep dive: what happens when you create a Pod?

1. The client submits a request to the API server.
2. Admission and validation are applied; the accepted object is persisted in etcd.
3. The scheduler observes an unscheduled Pod and selects an eligible node.
4. The binding is recorded; the node's kubelet observes the assigned Pod.
5. The kubelet works with the container runtime and networking/storage components to start containers.
6. The kubelet reports Pod/container state through the API server.

Be able to distinguish **desired state**, **observed state**, and the asynchronous controllers that converge them.

### Deep dive: how would you build a Kubernetes operator?

- Define a CRD schema and the custom resource's desired state.
- Watch relevant objects with a controller (typically via informers/caches and a work queue).
- Reconcile desired versus observed state and make **idempotent** changes.
- Handle conflicts, retries/backoff, transient failures, and finalizers where needed.
- Update status conditions and expose useful logs, metrics, and events.
- Test eventual convergence, deletions, and failure scenarios.

**Learning boundary:** Know how to explain the design even if you have not yet implemented an operator. Do not present conceptual knowledge as production experience.

### Deep dive: how would you automate cluster upgrades?

- Verify supported upgrade paths, component version skew, compatibility, and backups.
- Upgrade control-plane components through the supported mechanism.
- Cordon/drain nodes with awareness of PodDisruptionBudgets and local data.
- Upgrade nodes in controlled batches; check workload availability and health.
- Validate networking, storage, DNS, controllers, metrics, and application-level SLOs.
- Plan remediation/rollback carefully: rolling back control-plane or etcd state is **not always a simple reversal**.

### Deep dive: how would you secure a multi-tenant cluster?

- Namespaces and tightly scoped RBAC.
- NetworkPolicies plus supported enforcement in the CNI.
- Pod Security Standards and admission policies.
- Secret management and encryption/access controls.
- ResourceQuotas, LimitRanges, and tenant isolation at runtime and node levels.
- Auditability, monitoring, and appropriate cluster boundaries for stronger isolation.

### Practical troubleshooting exercise

For a failing workload, practice an ordered investigation:

1. `kubectl get pods -A -o wide` and `kubectl describe pod ...`.
2. Inspect `kubectl logs ... --previous` where crash loops are involved.
3. Review events, image-pull credentials, requests/limits, probes, environment, and ConfigMaps/Secrets.
4. Check Service endpoints, DNS, NetworkPolicies, and node readiness where relevant.
5. Explain the evidence, the fix, and how recurrence would be detected.

## E. Distributed-systems questions

1. How do you make an API operation idempotent?
2. When is retry safe, and when can it duplicate work?
3. What is exponential backoff with jitter?
4. What does a circuit breaker protect against?
5. What is eventual consistency, and when is it acceptable?
6. What is leader election? What happens on failover?
7. How does a distributed queue handle delivery and acknowledgment?
8. How do you handle duplicate or out-of-order events?
9. How do REST and gRPC differ in contracts, streaming, and operational behavior?
10. What is Protocol Buffers serialization and schema evolution?
11. How would you design an upstream health-checking system?
12. How do you handle partial failure across services?
13. Which logs, metrics, and traces would you collect for diagnosis?
14. How do you design for horizontal scalability?
15. What is backpressure, and when should you reject, delay, or drop work?

**Strong answer pattern:** Define the failure scenario, state assumptions, choose a mechanism, explain trade-offs, then identify an observable signal/test.

## F. Mini system-design and low-level design

Potential design prompts:

1. Design a Kubernetes controller that detects and remediates unhealthy workloads (discuss why simply restarting them may fight existing controllers).
2. Design a configuration synchronization system for network devices.
3. Design a distributed task scheduler.
4. Design a telemetry collection pipeline.
5. Design an API gateway rate limiter.
6. Design a job queue with retries and a dead-letter queue.
7. Design simple service discovery.
8. Design a closed-loop automation system that detects a fault and safely remediates it.

### Example: controller for custom resources

Explain the CRD, informers/cache, work queue, reconciliation, desired/actual state, idempotency, ownership, status updates, rate-limited retries, metrics, and tests. Include handling for events delivered repeatedly or out of order.

### Example: closed-loop automation

1. Collect telemetry and metrics with timestamps and source identity.
2. Detect a condition using rule/policy or anomaly logic.
3. Evaluate impact, authorization, safety thresholds, and action eligibility.
4. Trigger an idempotent remediation through a controlled API.
5. Observe resulting state and verify that the condition cleared.
6. Emit logs, metrics, audit evidence, and alerts; stop or escalate on repeated failure.

Consider deduplication, rate limits, rollbacks, failure domains, and preventing remediation loops.

## G. Questions to ask the technical interviewer

1. Does the team build operators/controllers, cluster lifecycle automation, or shared platform services?
2. What is the biggest scalability challenge today?
3. How mature is the cluster automation stack?
4. Does the team own provisioning, upgrades, security, and observability end-to-end?
5. What does success look like in the first three to six months?
6. How much engineering work is greenfield versus maintaining and evolving systems?
7. Is Go or Python the main implementation language, and how are they divided?
8. How does the team handle production incidents and on-call?
9. Where does agentic or closed-loop automation fit into the current architecture?
10. What are the largest technical risks or bottlenecks the team wants to solve?

Choose two to four questions that build on what you learned during the interview.

## H. Compact preparation plans

### Recruiter conversation

- 90-second introduction, two relevant project stories, and why the role fits.
- One Kubernetes deployment/troubleshooting story, not invented lifecycle ownership.
- Clear logistics (work mode, notice period, compensation) and recruiter questions.

### Technical screen

- Work through 20–30 appropriately selected easy-to-medium DSA problems **if time permits**.
- Practice Go worker pools, safe state sharing, cancellation, and tests.
- Review Kubernetes API flow, control plane, workload debugging, and reconciliation.
- Prepare one real scalability, one reliability, and one troubleshooting story.
- Explain retries, idempotency, backpressure, and partial failures with concrete examples.

## I. Must-practice shortlist for limited time

| Track | Highest priority |
|---|---|
| Coding | LRU Cache; Rate Limiter; Merge Intervals; Top K Frequent Elements; Course Schedule; Number of Islands; Longest Substring Without Repeating Characters; concurrent worker pool |
| Go | Worker pool; fan-in/fan-out; context cancellation; mutex versus channels; thread-safe cache |
| Kubernetes | `kubectl apply` path; operator/reconciliation loop; CrashLoopBackOff; safe cluster upgrade; multi-tenant security |
| Distributed systems | Backoff/retries; idempotency; event-driven processing; gRPC versus REST; health checks and self-healing |

## J. How to demonstrate technical maturity

For each problem, be ready to:

1. Clarify requirements and constraints before coding.
2. Describe your simple correct solution before optimizing.
3. Handle empty inputs, duplicate work, errors, cancellation, and shutdown where relevant.
4. Test normal behavior and failure paths; use `go test -race` for concurrent Go code.
5. State complexity and operational trade-offs.
6. Separate **what you have done at work** from **what you have learned or practiced**.

**Bottom line:** For the recruiter screen emphasize communication, role fit, and honest scope of experience; for technical rounds emphasize implementation, Go, Kubernetes, distributed systems, and operational reasoning.
