# VRP Packet Binding Model

---

## 1. Principle

Packets are not trusted.

They are validated against execution context.

---

## 2. Binding Context

A packet is bound to:

- session_id
- epoch
- authority
- mutation_id
- sequence
- path_id (optional)

---

## 3. Validation

A packet MUST be rejected if:

- binding does not match session context
- epoch mismatch
- authority mismatch
- mutation_id conflict
- sequence violation

---

## 4. Cryptographic Layer

Packet integrity MUST be protected via AEAD.

Binding MUST be part of authenticated data.

---

## 5. Path Independence

Packet validity does not depend on path.

Different paths MUST NOT affect correctness.

---

## 6. Late Arrival

Late packets:

- MAY arrive
- MUST be validated
- MUST be rejected if invalid