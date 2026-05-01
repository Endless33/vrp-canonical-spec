# VRP Authority Race Contract

This document defines how VRP resolves conflicting decisions
when multiple authorities attempt to commit simultaneously.

---

## 1. Problem Definition

In distributed systems, multiple nodes may attempt to:

- process the same mutation
- produce conflicting decisions
- act as authority at the same time

This condition is referred to as an authority race.

---

## 2. Core Guarantee

At any moment:

Only one canonical decision MUST be accepted for a mutation.

All other competing decisions MUST be rejected.

---

## 3. Authority Model

Authority is bound to:

- session
- epoch

At any given epoch:

Only one authority is considered valid.

---

## 4. Race Condition

A race occurs when:

- multiple authorities exist
- or multiple nodes believe they are authoritative
- or decisions are produced concurrently

---

## 5. Resolution Rules

A decision MAY be accepted if and only if:

- it belongs to the current epoch
- it is produced by the canonical authority
- it has not been superseded by a higher epoch decision

---

## 6. Epoch Priority

When two decisions conflict:

- the decision with the higher epoch MUST win
- lower epoch decisions MUST be rejected

---

## 7. Same Epoch Conflict

If two decisions exist within the same epoch:

A deterministic tie-break MUST be applied.

Possible tie-break inputs:

- authority identity
- deterministic ordering
- predefined priority rules

The tie-break MUST:

- produce a single canonical winner
- be consistent across all nodes

---

## 8. Non-Authority Rejection

Any decision from a non-canonical authority MUST be rejected.

This includes:

- outdated authorities
- parallel authorities
- misconfigured nodes

---

## 9. Late Arrival Handling

Late decisions:

- MAY arrive after a canonical decision is committed
- MUST NOT override committed state

---

## 10. Commit Safety

Once a mutation is committed:

- all competing decisions MUST be rejected
- no alternative outcome is allowed

---

## 11. Deterministic Outcome

All nodes observing the same set of inputs MUST converge to:

- the same canonical decision
- the same final state

---

## 12. Failure Handling

If authority cannot be resolved deterministically:

Execution MUST NOT proceed.

VRP prioritizes correctness over availability.

---

## 13. Summary

Authority races are expected.

They MUST NOT produce:

- multiple commits
- conflicting state
- non-deterministic outcomes

VRP enforces:

- single authority
- deterministic resolution
- canonical commit