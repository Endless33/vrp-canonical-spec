# VRP Convergence Contract

This document defines how independent runtimes converge to the same result.

---

## 1. Problem

In distributed systems:

- nodes receive inputs in different order
- network conditions differ per node
- timing is not consistent

Without strict rules, this leads to divergence.

---

## 2. Core Guarantee

Given the same logical input set:

All runtimes MUST converge to the same canonical decision.

---

## 3. Input Independence

The following MUST NOT affect the result:

- packet delivery order
- timing differences
- transport path
- duplication or delay

---

## 4. Deterministic Resolution

Convergence is achieved through:

- epoch ordering
- deterministic priority rules
- stable tie-break logic

These rules MUST produce the same result on all nodes.

---

## 5. Canonical Winner

For any conflicting decision set:

- exactly one winner MUST exist
- all other decisions MUST be rejected

---

## 6. No Divergence Allowed

Two runtimes observing the same logical inputs:

MUST NOT produce different outcomes.

If divergence occurs, it is a violation.

---

## 7. Execution Rule

same logical input set → same canonical result

---

## 8. Summary

VRP guarantees:

- deterministic convergence
- independence from delivery order
- identical outcomes across runtimes