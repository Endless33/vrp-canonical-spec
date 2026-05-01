# VRP Canonical Model

VRP (Veil Routing Protocol) defines a continuity-first execution model where session identity is independent of transport.

This document defines the canonical model of execution.

---

## 1. Core Principle

A session is not bound to any transport.

Transport is a delivery mechanism.
Session identity is a runtime property.

Failure of transport MUST NOT imply loss of session identity.

---

## 2. Execution Model

VRP operates under unreliable network conditions.

The network may:

- duplicate packets
- drop packets
- reorder packets
- change delivery paths

These conditions are treated as normal.

Correctness MUST NOT depend on transport behavior.

---

## 3. Commit-Centric Correctness

Correctness is enforced at the commit boundary.

A mutation is not considered valid because it was delivered.

A mutation is valid only if it is accepted at commit.

---

## 4. Deterministic State Evolution

For any logical mutation:

- it may be delivered multiple times
- it may arrive through different paths
- it may be delayed or reordered

Despite this:

Only one canonical commit is allowed.

---

## 5. No Retry Semantics

Retry is not part of the correctness model.

Retry duplicates execution attempts.

VRP does not rely on retry for recovery.

Instead:

- inputs may be re-evaluated
- but state mutation is strictly controlled

---

## 6. Authority and Decision Control

At any moment, only one authority may produce a canonical decision for a session.

Multiple inputs do not imply multiple valid outcomes.

Authority determines which mutation becomes canonical.

---

## 7. Epoch-Based Continuity

Execution is divided into epochs.

Epochs:

- are strictly monotonic
- define authority boundaries
- prevent stale decisions from mutating state

Old epoch inputs may be observed but MUST NOT mutate state.

---

## 8. Packet Independence

Packets are not trusted as sources of truth.

They are evaluated against:

- session context
- authority
- epoch
- commit rules

A valid packet may still be rejected.

---

## 9. Replay and Duplication

Duplicate delivery is expected.

Replay is not recovery.

A previously committed mutation MUST NOT be committed again.

Replay MUST be rejected at the commit layer.

---

## 10. Failure Handling

Failure is treated as a runtime event.

If correctness cannot be preserved:

Execution MUST NOT continue.

VRP does not attempt to repair invalid state transitions.

Invalid transitions are rejected.

---

## 11. Canonical Guarantee

For any session:

- state evolves deterministically
- duplicate inputs do not corrupt state
- transport behavior does not affect correctness

---

## 12. Summary

VRP does not attempt to make the network reliable.

VRP enforces correctness despite unreliable conditions.

Correctness is a property of execution,
not of transport.