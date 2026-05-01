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

This is not an implementation.

This is the canonical specification defining:

- Commit contract
- Authority resolution rules
- Epoch semantics
- Packet binding model
- Replay semantics
- System invariants

---

## Specification Documents

- `docs/VRP_CANONICAL_MODEL.md`
- `docs/VRP_COMMIT_CONTRACT.md`
- `docs/VRP_AUTHORITY_AND_EPOCHS.md`
- `docs/VRP_PACKET_BINDING.md`
- `docs/VRP_REPLAY_SEMANTICS.md`
- `docs/VRP_INVARIANTS.md`

---

## Quick Demo (Commit Contract)

This repository includes a minimal executable demo of the VRP commit contract.

It demonstrates:

- duplicate mutation rejection
- non-authority rejection
- stale epoch rejection
- single canonical commit

---

### Run locally

```bash
git clone https://github.com/Endless33/vrp-canonical-spec
cd vrp-canonical-spec
go mod init vrp-canonical-spec
go run ./cmd/private_canonical_contract_demo

## Key Statement

Correctness is not assumed from the network.

Correctness is enforced during execution.

---

## Status

Canonical specification (in progress)

---

## Author

VRP / Jumping VPN
