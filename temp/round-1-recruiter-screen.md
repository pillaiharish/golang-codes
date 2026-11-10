# Round 1 — Recruiter Screen

**Reference only · company-neutral**

This round is often a roughly **30-minute conversation** about role fit, communication, motivation, experience, and logistics. Exact interview formats vary. Prepare two categories: **questions you may be asked** and **questions you should ask**.

## A. Questions you may be asked

### 1. Tell me about yourself

Prepare a **60–90-second introduction** with:

- Current role and approximate total experience.
- Primary language: Go, with relevant Python or automation exposure.
- Kubernetes deployment and debugging experience.
- Scalable/backend and distributed systems examples.
- Why the role is a logical next step.

**Example (adapt to your actual experience):**

> I have around four years of experience spanning software development and test automation, with Go as my primary language. I work on DNS-oriented backend and platform components, including policy enforcement and reliability-related functionality. I have deployed and debugged applications on Kubernetes, and I have worked with automation and distributed service behavior. I want to deepen my platform-engineering experience in Kubernetes lifecycle management, controllers, and large-scale automation. This role brings together Go, networking, reliable systems, and the areas in which I want to grow.

**Honesty check:** Deploying/debugging workloads does **not** automatically mean you have built operators, provisioned clusters, or owned cluster upgrades. Describe those areas as learning goals unless you have done them.

### 2. Why are you interested in this organization?

Replace company-specific praise with questions you can substantiate:

- Is the team building cloud-native platform infrastructure?
- Does the product involve network automation, orchestration, or service-provider scale?
- Does the role offer ownership of reliable, production-grade Go services?
- What about the engineering problems genuinely interests you?

**Example:**

> I'm interested in the team's work on cloud-native network automation. I enjoy building reliable infrastructure components, and I want to work more deeply on Kubernetes platform engineering and distributed control systems. The engineering problems described for this role connect well with my Go and networking background.

### 3. Why this platform-engineering role?

Connect Go, network services, automation, reliability, and your interest in Kubernetes internals.

> I like systems where correctness, observability, and operational reliability matter. I've worked on backend/networking components and want to extend that experience toward platform automation, Kubernetes controllers, and infrastructure lifecycle engineering.

### 4. What is your Go experience?

Have one real example ready for each applicable area:

- Backend services, DNS/networking components, and APIs.
- Goroutines, channels, `context`, synchronization, and cancellation.
- Interfaces, error handling, testing, and clean package boundaries.
- Health checks, performance, production debugging, and automation.
- REST/gRPC, CLI tools, or Kubernetes clients **only if actually used**.

**Example:**

> Go is my primary programming language. I've used it for network-facing backend components and reliability features. I can discuss concurrency, context cancellation, interfaces, error handling, testability, and the trade-offs I made in production code.

### 5. What is your Kubernetes experience?

Know how to discuss these topics, separating hands-on experience from study:

- Deployments, StatefulSets, DaemonSets, Pods, Services, and Ingress.
- ConfigMaps, Secrets, probes, namespaces, RBAC, and NetworkPolicies.
- Helm or Kustomize if used; troubleshooting, logs, events, and rollout history.
- Cluster management, provisioning, upgrades, and configuration automation.
- CRDs, controllers/operators, and the reconciliation model.

**Example grounded in existing experience:**

> My hands-on Kubernetes experience is deploying applications and debugging workloads. I can investigate pod status, logs, events, configuration, and connectivity issues. I am strengthening my knowledge of control-plane internals, controllers, cluster upgrades, and provisioning because those are important responsibilities of this role.

**Follow-up:** Prepare one concrete workload-debugging story: symptom → commands/evidence → root cause → fix → prevention.

### 6. Have you worked on distributed systems?

Choose a concrete component and explain **retries, timeouts, concurrency, health checks, partial failure, and backpressure** where applicable. Avoid reciting patterns you have not implemented.

> I've worked on network-facing services where upstream health, timeouts, retry behavior, and failures matter. I can walk through a production example, explain the failure mode, and describe the trade-offs in handling it.

### 7. What is your network-automation experience?

Useful themes:

- DNS/networking fundamentals and protocol-aware services.
- REST/gRPC APIs and configuration management.
- Automation pipelines, observability, and telemetry.
- Desired versus observed state and closed-loop remediation.

> My core experience is Go-based network services and automation. I have not necessarily built a full network-device orchestrator, but the underlying reliability, API, and configuration concepts are familiar. I am interested in applying them to closed-loop platform automation.

### 8. Are you comfortable with the location and work arrangement?

Confirm the expected office location, hybrid/onsite cadence, relocation/travel requirements, and your actual flexibility. Do not promise a schedule you cannot maintain.

### 9. What are your current compensation, expectations, and notice period?

Have exact, current figures ready privately. Do **not** put personal compensation or notice-period details in public notes.

**Possible response:**

> I'm evaluating the scope and level of the role alongside compensation. Could you share the budgeted range and compensation mix so we can establish whether our expectations align?

Answer direct factual questions accurately and distinguish fixed, variable, equity, and other components.

### 10. Are you interviewing elsewhere?

Be honest and concise.

> I'm exploring a small number of relevant backend/platform roles. This team's focus on Go, Kubernetes, and distributed infrastructure is particularly aligned with the work I want to do.

### 11. Why should we move your profile forward?

**Suggested response (adapt, do not memorize mechanically):**

> My experience combines Go software development, networking, and building reliable backend components. I've deployed and troubleshot applications on Kubernetes, and I understand how production services must handle failures and remain observable. I'm clear about the platform areas I still need to learn, and I'm motivated to take on deeper Kubernetes lifecycle and automation work. I believe that gives me a solid base for this role.

## B. Questions you should ask the recruiter

1. How large is the platform-engineering team, and where does it sit in the organization?
2. Is the role primarily building Kubernetes controllers/operators, or operating and scaling clusters?
3. What does the full interview process look like after this conversation?
4. What level is this position mapped to: mid-level or senior?
5. Is the team building new platform capabilities or extending an established system?
6. What is the work arrangement: hybrid, remote, or onsite?
7. What is the target hiring timeline?
8. What is the budgeted compensation range and structure?
9. Does the role include on-call responsibility or incident response?
10. What are the next steps and expected feedback timeline?

Choose **three to five** based on what the recruiter has already covered. A good priority order is role scope, level, process, work arrangement, and compensation range.

## C. Quick preparation checklist

- [ ] Practice the introduction within 90 seconds.
- [ ] Prepare two evidence-backed project stories aligned with Go/platform work.
- [ ] Prepare one real Kubernetes troubleshooting example.
- [ ] Explain why this role fits your career direction.
- [ ] Know exact notice-period and compensation facts without publishing them here.
- [ ] Choose three to five useful recruiter questions.

**Guiding principle:** Represent hands-on work accurately; label concepts you are learning rather than presenting them as production ownership.
