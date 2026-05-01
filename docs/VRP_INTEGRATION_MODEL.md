# VRP Integration Model

VRP is an execution correctness layer for unreliable networks.

It is not limited to VPN tunneling.

VRP can be integrated wherever a system must preserve correctness while transport, delivery order, or network paths are unstable.

---

## 1. What VRP Wraps

VRP wraps execution boundaries.

Examples:

- API request execution
- financial transaction execution
- distributed job execution
- state mutation admission
- session migration
- authority handoff

VRP does not replace the application.

It controls whether a state-changing operation is allowed to commit.

---

## 2. Integration Position

VRP sits between:

- transport delivery
- application state mutation

Transport may deliver inputs.

Application may request state changes.

VRP decides whether execution is valid.

---

## 3. Minimal Integration Flow

```text
transport input
    ↓
VRP validation
    ↓
authority / epoch check
    ↓
replay / duplicate check
    ↓
commit boundary
    ↓
application mutation

---

4. Core Rule
Delivery does not imply validity.
A request may arrive successfully and still be rejected.
A mutation is valid only if it passes the commit boundary.

---

5. Integration Modes
Library mode
Application calls VRP before applying mutation.

if vrp.Accept(input):
    apply mutation
else:
    reject

Sidecar mode
VRP runs beside the application and filters mutation requests.
Gateway mode
VRP sits before a service boundary and enforces execution correctness before forwarding.

---

6. What VRP Provides
VRP provides:
duplicate mutation rejection
stale epoch rejection
non-authority rejection
deterministic authority resolution
convergence under disorder
commit-level correctness

---

7. What VRP Does Not Assume
VRP does not assume:
reliable transport
stable paths
ordered delivery
honest timing
retry correctness

---

8. Operational Outcome
With VRP:
duplicate inputs do not double-execute
stale decisions cannot mutate state
authority races converge to one winner
delivery order does not define correctness
session identity survives transport instability

---

9. Summary
VRP is not a retry system.
VRP is not a recovery wrapper.
VRP is an execution correctness boundary for unreliable networks.

---

## 2. `docs/VRP_USE_CASES.md`

```markdown
# VRP Use Cases

This document describes practical scenarios where VRP semantics apply.

VRP is designed for systems where network disorder must not corrupt execution.

---

## 1. Bank Transfer

### Problem

A payment request is sent.

The response is lost.

The client retries.

The system receives the same logical mutation twice.

Without a commit boundary, this may produce double execution.

---

### VRP Behavior

```text
payment-001 → ACCEPTED
payment-001 retry → REJECTED_DUPLICATE

---

Guarantee

one logical transfer → one committed state transition

---

2. Distributed Job Execution
Problem
A job scheduler sends work to a node.
The network duplicates the request.
Two workers attempt to process the same job.
Without deterministic commit admission, the job may execute twice.

---

VRP Behavior

job-778 from node-a → ACCEPTED
job-778 from node-b → REJECTED_DUPLICATE

Only one canonical execution is allowed.

---

Guarantee

one logical job → one canonical execution

---

3. API Request Under Retry
Problem
An API request times out.
The client retries.
The server receives multiple equivalent mutation attempts.
Transport success does not prove execution correctness.

---

VRP Behavior

request-42 → ACCEPTED
request-42 retry → REJECTED_DUPLICATE

Retry is not treated as recovery.
It is treated as a repeated execution attempt.

---

Guarantee

retry cannot create a second valid mutation

---

4. Authority Race
Problem
Two authorities attempt to decide the same mutation.
Network timing makes both appear valid locally.
Without deterministic resolution, state may diverge.

---

VRP Behavior

node-a epoch=2 → REJECTED_LOWER_EPOCH
node-b epoch=3 → ACCEPTED_WINNER
node-c epoch=3 lower priority → REJECTED_LOWER_PRIORITY

---

Guarantee

multiple decision sources → one canonical winner

---

5. Path Migration
Problem
A session moves from one path to another.
Packets from the old path may still arrive late.
Without session-level validation, stale inputs may mutate state.

---

VRP Behavior

old path input → evaluated
stale or duplicate input → rejected
new path input → evaluated under current epoch

---

Guarantee

path change does not reset session identity

---

6. Summary
VRP applies wherever the system must preserve correctness under:
retries
packet duplication
packet loss
out-of-order delivery
authority races
path migration
stale inputs
VRP does not make the network reliable.
VRP makes execution independent of network reliability.

---

## 3. `docs/VRP_POSITIONING.md`

```markdown
# VRP Positioning

VRP is an execution correctness layer for unreliable networks.

It is not defined as a VPN product.

VPN is one possible transport surface.

VRP defines how sessions, mutations, authority, and commit correctness behave when transport is unstable.

---

## 1. What VRP Is

VRP is:

- a continuity-first execution model
- a commit-boundary protocol
- a deterministic authority resolution model
- a replay-safe mutation admission layer
- a convergence model under disorder

---

## 2. What VRP Is Not

VRP is not:

- a retry framework
- a reconnect system
- a tunnel wrapper
- a load balancer
- a packet router only
- a best-effort recovery mechanism

---

## 3. Core Distinction

Traditional transport systems focus on delivery.

VRP focuses on execution correctness.

```text
transport asks:
did the packet arrive?

VRP asks:
is this mutation allowed to change state?

---

4. Why This Matters
Reliable delivery does not guarantee correct execution.
A packet can arrive twice.
A request can retry.
A stale authority can speak late.
A path can change.
VRP treats all of these as expected runtime conditions.

---

5. Correct Framing
VRP should be described as:

an execution correctness layer for unreliable networks

or:

a continuity-first protocol model where session identity is independent of transport

---

6. Incorrect Framing
VRP should not be reduced to:

next-gen VPN

or:

VPN endpoint rotation

Those descriptions miss the core model.

---

7. Canonical Statement
Session identity is above transport.
Correctness is enforced at the commit layer.
Replay is not recovery.
Retry is not correctness.
Convergence must be deterministic.

---

8. Summary
VRP is not about making packets reliable.
VRP is about ensuring that unreliable packet behavior cannot corrupt execution.

---