# VRP API

This document defines the minimal API surface for integrating VRP as an execution correctness layer.

VRP does not decide whether a packet arrived.

VRP decides whether a mutation is allowed to commit.

---

## 1. Minimal API

Accept(input) → decision

Where decision is one of:

ACCEPTED  
REJECTED_DUPLICATE  
REJECTED_NON_AUTHORITY  
REJECTED_STALE_EPOCH  
REJECTED_REPLAY  
REJECTED_INVALID_BINDING  

---

## 2. Input Shape

A minimal VRP input contains:

session_id  
mutation_id  
authority  
epoch  
sequence  
payload_hash  

Optional:

path_id  
timestamp  
metadata  

---

## 3. Example Input

session_id=bank-session-001  
mutation_id=payment-001  
authority=node-b  
epoch=3  
sequence=12  
payload_hash=hash(payment:100)  

---

## 4. Core Rule

A mutation is not valid because it arrived.

A mutation is valid only if Accept(input) returns:

ACCEPTED  

---

## 5. Commit Boundary

Application MUST only mutate state after acceptance.

decision = vrp.Accept(input)

if decision == ACCEPTED:
    apply mutation
else:
    reject

---

## 6. Rejection Is Normal

Rejection is not failure.

Rejection prevents invalid execution.

Examples:

duplicate → REJECTED_DUPLICATE  
old epoch → REJECTED_STALE_EPOCH  
wrong authority → REJECTED_NON_AUTHORITY  
old sequence → REJECTED_REPLAY  

---

## 7. Determinism

Given the same state and input:

Accept(input)

MUST return the same decision.

---

## 8. Summary

VRP answers one question:

is this mutation allowed to change state?