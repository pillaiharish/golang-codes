# Round 1 — Recruiter Screen

**Reference only · company-neutral · privacy-safe template**

**Privacy note:** This public document intentionally uses placeholders rather than anyone's work history, experience duration, employer, or other personal profile details. Fill in the examples only in a private interview-preparation copy; do not commit completed answers to a public repository.

This round is often a roughly **30-minute conversation** about role fit, communication, motivation, experience, and logistics. Exact interview formats vary. Prepare two categories: **questions you may be asked** and **questions you should ask**.

## A. Questions you may be asked

### 1. Tell me about yourself

Prepare a **60–90-second introduction** with:

- Your current role and relevant experience (use placeholders in public templates).
- Languages relevant to the role (for example, Go or Python).
- Kubernetes work, if any, described at your actual level of involvement.
- Backend or distributed-systems projects you can substantiate.
- Why the role is a logical next step.

**Fill-in template (replace placeholders privately before an interview):**

> I have [X years] of experience in [relevant engineering areas]. My strongest language is [language], and I have worked on [general category of systems or projects]. My hands-on experience with [platform or infrastructure technology] includes [tasks you can verify]. I'm interested in growing toward [relevant responsibilities], and this role connects that goal with [the team's engineering problems].

**Honesty check:** Deploying/debugging workloads does **not** automatically mean you have built operators, provisioned clusters, or owned cluster upgrades. Describe those areas as learning goals unless you have done them.

### 2. Why are you interested in this organization?

Replace company-specific praise with questions you can substantiate:

- Is the team building cloud-native platform infrastructure?
- Does the product involve network automation, orchestration, or service-provider scale?
- Does the role offer ownership of reliable, production-grade Go services?
- What about the engineering problems genuinely interests you?

**Example:**

> I'm interested in [specific technical problem the organization works on]. I enjoy [type of engineering work], and the role would let me build on [skills I have actually used] while developing [skills I want to learn].

### 3. Why this platform-engineering role?

Connect your relevant programming experience, system-design work, operational experience, and genuine interest in the role's platform responsibilities.

> I like working on systems where [engineering qualities that matter to you] are important. My work on [substantiated project category] motivates me to learn more about [target responsibilities].

### 4. What is your Go experience?

Have one real example ready for each applicable area:

- Relevant backend services, APIs, and networking components (only if applicable).
- Goroutines, channels, `context`, synchronization, and cancellation.
- Interfaces, error handling, testing, and clean package boundaries.
- Health checks, performance, production debugging, and automation.
- REST/gRPC, CLI tools, or Kubernetes clients **only if actually used**.

**Example:**

> I've used Go for [verified project types]. One example is [project], where I applied [specific language feature] and had to consider [trade-off]. I can explain the design and how I tested the result.

### 5. What is your Kubernetes experience?

Know how to discuss these topics, separating hands-on experience from study:

- Deployments, StatefulSets, DaemonSets, Pods, Services, and Ingress.
- ConfigMaps, Secrets, probes, namespaces, RBAC, and NetworkPolicies.
- Helm or Kustomize if used; troubleshooting, logs, events, and rollout history.
- Cluster management, provisioning, upgrades, and configuration automation.
- CRDs, controllers/operators, and the reconciliation model.

**Experience-neutral template (fill only with verified work):**

> My hands-on Kubernetes experience includes [tasks actually performed]. For example, I investigated [issue] using [tools or evidence] and reached [outcome]. I'm currently learning [areas relevant to the role that I have not yet owned].

**Follow-up:** Prepare one concrete workload-debugging story: symptom → commands/evidence → root cause → fix → prevention.

### 6. Have you worked on distributed systems?

Choose a concrete component and explain **retries, timeouts, concurrency, health checks, partial failure, and backpressure** where applicable. Avoid reciting patterns you have not implemented.

> I've worked on [a system I can discuss] where [one relevant failure mode] mattered. I can describe the observed behavior, my contribution, and why we chose [mitigation or trade-off].

### 7. What is your network-automation experience?

Useful themes:

- DNS/networking fundamentals and protocol-aware services.
- REST/gRPC APIs and configuration management.
- Automation pipelines, observability, and telemetry.
- Desired versus observed state and closed-loop remediation.

> My relevant experience includes [projects and technologies actually used]. I can explain how [a verified project] relates to [target platform concept], while being clear about the parts I have not worked on directly.

### 8. Are you comfortable with the location and work arrangement?

Confirm the expected office location, hybrid/onsite cadence, relocation/travel requirements, and your actual flexibility. Do not promise a schedule you cannot maintain.

### 9. What are your current compensation, expectations, and notice period?

Have exact, current figures ready privately. Do **not** put personal compensation or notice-period details in public notes.

**Possible response:**

> I'm evaluating the scope and level of the role alongside compensation. Could you share the budgeted range and compensation mix so we can establish whether our expectations align?

Answer direct factual questions accurately and distinguish fixed, variable, equity, and other components.

### 10. Are you interviewing elsewhere?

Be honest and concise.

> I'm exploring opportunities that match [my genuine interests and skills]. This team's work on [verified team responsibility] is especially relevant because [specific reason].

### 11. Why should we move your profile forward?

**Suggested response (adapt, do not memorize mechanically):**

> My experience includes [relevant skills] and [specific evidence-backed project]. I contributed [action] and observed [result]. I understand the core responsibilities of this role, and I can be precise about what I've done versus what I'm still learning.

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
