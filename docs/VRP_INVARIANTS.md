# VRP Invariants

---

## 1. Single Commit Invariant

A logical mutation may commit at most once.

---

## 2. Authority Invariant

Only current authority may produce valid commits.

---

## 3. Epoch Invariant

State MUST NOT be mutated by stale epochs.

---

## 4. Determinism Invariant

Given the same sequence of accepted mutations:

State MUST be identical.

---

## 5. Replay Safety Invariant

Replay MUST NOT alter committed state.

---

## 6. Transport Independence Invariant

Transport behavior MUST NOT affect correctness.

---

## 7. Continuity Invariant

Session identity MUST survive:

- path change
- packet loss
- duplication

---

## 8. Failure Containment Invariant

If correctness cannot be preserved:

Execution MUST stop.

---

## 9. Canonical Decision Invariant

For any mutation:

Only one decision is allowed.

---

## 10. Integrity Invariant

Invalid transitions MUST be rejected.

They MUST NOT be repaired.