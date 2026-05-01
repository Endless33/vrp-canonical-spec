# VRP Network Chaos Contract

This document defines how VRP behaves under unreliable network conditions.

It extends the canonical commit contract to real network behavior.

---

## 1. Network Model

The network is assumed to be unreliable.

It may:

- duplicate packets
- drop packets
- reorder packets
- deliver packets through different paths

These behaviors are not exceptions.

They are normal operating conditions.

---

## 2. Core Guarantee

Network disorder MUST NOT corrupt committed state.

Regardless of network behavior:

- each logical mutation may commit at most once
- invalid or duplicated inputs MUST NOT alter state

---

## 3. Duplicate Handling

Duplicate packets:

- MAY be delivered multiple times
- MUST be evaluated independently
- MUST be rejected if mutation is already committed

---

## 4. Drop Handling

Dropped packets:

- MAY never reach the runtime
- MUST NOT trigger recovery logic at commit layer
- MUST NOT produce partial or inconsistent state

---

## 5. Reordering

Packets MAY arrive out of order.

The runtime:

- MUST evaluate packets against current state
- MUST reject stale or invalid mutations

---

## 6. Path Independence

Packets may arrive via different paths.

Path variation:

- MUST NOT affect correctness
- MUST NOT change commit outcome

---

## 7. Execution Rule

For any set of network conditions:

The final committed state MUST be identical
to the state produced under ideal delivery.

---

## 8. No Recovery Semantics

VRP does not repair incorrect execution.

Instead:

- invalid transitions are rejected
- only valid transitions are committed

---

## 9. Canonical Outcome

Given:

- duplicate delivery
- dropped packets
- reordering
- path changes

The runtime MUST produce:

- a single canonical state
- no duplicate commits
- no corrupted state

---

## 10. Reference Demo

See:

cmd/private_network_chaos_contract_demo/main.go

This demo demonstrates:

- duplicate rejection
- dropped packet handling
- consistent final state

---

## 11. Summary

The network may behave arbitrarily.

Correctness must not.

VRP enforces correctness at the commit boundary,
not at the transport layer.