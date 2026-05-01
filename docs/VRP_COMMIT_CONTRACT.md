# VRP Commit Contract

This document defines the canonical commit contract of VRP.

---

## 1. Commit Boundary

A mutation is considered valid only if it is accepted at the commit layer.

Delivery does not imply validity.

---

## 2. Commit Conditions

A mutation MAY commit if and only if:

- session_id matches active session
- authority is current
- epoch is current or valid forward transition
- mutation_id has not been committed before
- packet binding is valid
- replay window accepts the sequence

---

## 3. Idempotency Enforcement

Each mutation_id MUST commit at most once.

Duplicate mutation_id MUST be rejected.

---

## 4. Authority Requirement

Only the current authority MAY produce a valid commit.

Non-authoritative inputs MUST be rejected.

---

## 5. Epoch Constraint

A mutation MUST NOT commit if:

- it belongs to a stale epoch
- it attempts to override a committed state from a newer epoch

---

## 6. Deterministic Decision

For any mutation:

- multiple inputs may exist
- only one canonical decision is allowed

---

## 7. Rejection Rules

A mutation MUST be rejected if:

- duplicate
- stale
- non-authoritative
- invalid binding
- replay violation

---

## 8. No Retry at Commit Layer

Commit layer does not implement retry.

Retry is treated as duplicate execution attempt.

---

## 9. Finality

Once committed:

- mutation is final
- state cannot be rolled back via replay