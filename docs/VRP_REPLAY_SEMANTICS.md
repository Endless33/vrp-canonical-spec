# VRP Replay Semantics

---

## 1. Replay Definition

Replay is the re-delivery of a previously sent input.

---

## 2. Replay vs Retry

Replay:
- network-level duplication

Retry:
- application-level re-execution

VRP does not treat retry as recovery.

---

## 3. Replay Handling

A replayed mutation MUST NOT commit again.

---

## 4. Duplicate Delivery

Duplicate inputs:

- MAY be delivered
- MUST be evaluated
- MUST be rejected if already committed

---

## 5. Sequence Control

Replay window MUST enforce sequence validity.

Old sequence numbers MUST be rejected.

---

## 6. Commit Safety

Replay MUST NOT produce a second valid commit.

---

## 7. Re-evaluation

Inputs MAY be re-evaluated.

But:

State mutation MUST remain singular.