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

transport asks:
did the packet arrive?

VRP asks:
is this mutation allowed to change state?

---

## 4. Why This Matters

Reliable delivery does not guarantee correct execution.

A packet can arrive twice.

A request can retry.

A stale authority can act late.

A path can change.

VRP treats all of these as expected runtime conditions.

---

## 5. Correct Framing

VRP should be described as:

execution correctness layer for unreliable networks

or:

continuity-first protocol model where session identity is independent of transport

---

## 6. Incorrect Framing

VRP should not be reduced to:

next-gen VPN  
endpoint rotation  

Those descriptions miss the core model.

---

## 7. Canonical Statement

Session identity is above transport.

Correctness is enforced at the commit layer.

Replay is not recovery.

Retry is not correctness.

Convergence must be deterministic.

---

## 8. Summary

VRP is not about making packets reliable.

VRP is about ensuring that unreliable packet behavior cannot corrupt execution.