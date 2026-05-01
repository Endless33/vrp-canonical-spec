# VRP Glossary

This document defines core VRP terminology.

---

## Session

A logical execution context independent of transport.

---

## Mutation

A state-changing operation applied to a session.

---

## Commit

The act of accepting a mutation as valid and applying it to state.

---

## Commit Boundary

The layer where correctness is enforced.

---

## Authority

The entity allowed to produce valid decisions for a session.

---

## Epoch

A monotonic version defining authority validity.

Higher epochs override lower ones.

---

## Replay

A repeated mutation attempt.

Replays MUST NOT change committed state.

---

## Duplicate

A mutation that has already been observed.

Duplicates MUST be rejected.

---

## Network Chaos

Unreliable network behavior:

- duplication
- loss
- reordering
- delay

---

## Convergence

The property that independent runtimes produce the same result.

---

## Canonical Decision

The single accepted outcome among competing inputs.

---

## Determinism

The property that the same inputs produce the same result.

---

## Summary

VRP terminology is defined around execution, not transport.