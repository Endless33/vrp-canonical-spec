# VRP Security Boundary

This document defines what VRP trusts and what it does not.

---

## 1. Core Principle

VRP does not trust the network.

---

## 2. Untrusted Domain

The following are considered untrusted:

- packet delivery
- transport integrity
- ordering guarantees
- timing assumptions

The network may:

- duplicate packets
- drop packets
- reorder packets
- inject delayed inputs

---

## 3. Trusted Boundary

Correctness is enforced at:

- commit layer
- authority resolution
- epoch validation
- replay protection

---

## 4. Validation Rules

A mutation is valid only if:

- it belongs to the current epoch
- it is produced by the canonical authority
- it has not been committed before
- it passes replay constraints

---

## 5. Rejection Model

Invalid inputs MUST be rejected:

- duplicates
- stale epoch inputs
- non-authority inputs
- replayed mutations

---

## 6. No Implicit Trust

VRP does not assume:

- reliable delivery
- honest intermediaries
- correct ordering

---

## 7. Security Outcome

The system guarantees:

- no duplicate commits
- no state corruption
- deterministic validation
- bounded authority

---

## 8. Summary

Trust is not placed in the network.

Trust is enforced at execution boundaries.