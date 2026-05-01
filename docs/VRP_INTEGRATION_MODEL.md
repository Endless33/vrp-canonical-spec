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

transport input
→ VRP validation
→ authority / epoch check
→ replay / duplicate check
→ commit boundary
→ application mutation

---

## 4. Core Rule

Delivery does not imply validity.

A request may arrive successfully and still be rejected.

A mutation is valid only if it passes the commit boundary.

---

## 5. Integration Modes

### Library mode

Application calls VRP before applying mutation.

if vrp.Accept(input):
    apply mutation
else:
    reject

### Sidecar mode

VRP runs beside the application and filters mutation requests.

### Gateway mode

VRP sits before a service boundary and enforces execution correctness before forwarding.

---

## 6. What VRP Provides

VRP provides:

- duplicate mutation rejection
- stale epoch rejection
- non-authority rejection
- deterministic authority resolution
- convergence under disorder
- commit-level correctness

---

## 7. What VRP Does Not Assume

VRP does not assume:

- reliable transport
- stable paths
- ordered delivery
- honest timing
- retry correctness

---

## 8. Operational Outcome

With VRP:

- duplicate inputs do not double-execute
- stale decisions cannot mutate state
- authority races converge to one winner
- delivery order does not define correctness
- session identity survives transport instability

---

## 9. Summary

VRP is not a retry system.

VRP is not a recovery wrapper.

VRP is an execution correctness boundary for unreliable networks.