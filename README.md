# VRP (Veil Routing Protocol)

VRP is a continuity-first execution model where session identity is independent of transport.

This repository defines the canonical specification of VRP.

---

## Core Properties

- Session identity does not depend on transport
- Transport failure does not imply session reset
- Duplicate inputs must not produce duplicate state transitions
- Correctness is enforced at the commit layer
- Replay is not treated as recovery

---

## What This Repository Contains

This is not a production implementation.

This is the canonical specification defining:

- Commit contract
- Authority resolution rules
- Epoch semantics
- Packet binding model
- Replay semantics
- System invariants

---

## Specification Documents

- docs/VRP_CANONICAL_MODEL.md
- docs/VRP_COMMIT_CONTRACT.md
- docs/VRP_AUTHORITY_AND_EPOCHS.md
- docs/VRP_PACKET_BINDING.md
- docs/VRP_REPLAY_SEMANTICS.md
- docs/VRP_INVARIANTS.md
- docs/VRP_NETWORK_CHAOS_CONTRACT.md
- docs/VRP_AUTHORITY_RACE_CONTRACT.md

---

## Executable Demos

These demos are not simulations.
They are executable forms of the specification.

---

### 1. Commit Contract

go run ./cmd/private_canonical_contract_demo

Proves:
- duplicate → rejected
- non-authority → rejected
- stale epoch → rejected
- valid → committed once

---

### 2. Network Chaos

go run ./cmd/private_network_chaos_contract_demo

Proves:
- duplicate packets do not corrupt state
- dropped packets do not produce inconsistency
- final state remains valid

---

### 3. Authority Race

go run ./cmd/private_authority_race_demo

Proves:
- competing authorities produce one canonical decision
- lower epoch is rejected
- same epoch conflict is deterministically resolved

---

### 4. Multi-Node Convergence

go run ./cmd/private_multi_node_convergence_demo

Proves:
- independent runtimes select the same winner
- decision does not depend on node instance

---

### 5. Disorder + Multi-Node Convergence

go run ./cmd/private_disorder_multi_node_convergence_demo

Proves:
- different input order does not affect result
- nodes converge to the same canonical decision
- disorder does not break determinism

---

## How to Run

git clone https://github.com/Endless33/vrp-canonical-spec
cd vrp-canonical-spec
go run ./cmd/private_canonical_contract_demo

---

## Key Statement

Correctness is not assumed from the network.

Correctness is enforced during execution.

---

## Status

Canonical specification in progress.

---

## Author

Vitalijus Riabovas
VRP / Jumping VPN
EOF