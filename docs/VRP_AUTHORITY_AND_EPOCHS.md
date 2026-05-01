# VRP Authority and Epochs

---

## 1. Authority Model

At any given moment:

Only one authority is considered canonical for a session.

---

## 2. Authority Properties

Authority:

- defines valid decision source
- determines canonical commit
- enforces ordering

---

## 3. Authority Conflict

If multiple authorities exist:

- higher epoch wins
- same epoch requires deterministic resolution

Non-winning authority MUST be ignored.

---

## 4. Epoch Definition

Epoch represents a continuity segment of execution.

---

## 5. Epoch Properties

Epochs are:

- strictly monotonic
- non-reversible
- authoritative boundaries

---

## 6. Epoch Transition

Transition to a new epoch MUST:

- invalidate previous authority
- prevent stale mutations from committing

---

## 7. Stale Input Handling

Inputs from older epochs:

- MAY be observed
- MUST NOT mutate state

---

## 8. Authority + Epoch Binding

Authority is valid only within its epoch.

Authority outside its epoch MUST be rejected.