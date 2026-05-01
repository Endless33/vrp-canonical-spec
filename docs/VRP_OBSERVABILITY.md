# VRP Observability

This document defines the minimum observability required for a VRP-enabled system.

VRP must explain every decision.

---

## 1. Why Observability Matters

In a continuity-first model:

- rejection is expected
- duplicates are normal
- disorder is normal

The system must show why decisions were made.

---

## 2. Required Log Fields

Every mutation decision SHOULD include:

session_id  
mutation_id  
authority  
epoch  
sequence  
decision  
reason  

Optional:

path_id  
payload_hash  
timestamp  

---

## 3. Example Logs

session=session-001 mutation=payment-001 authority=node-b epoch=3 sequence=12 decision=ACCEPTED reason=canonical_commit  

session=session-001 mutation=payment-001 authority=node-b epoch=3 sequence=13 decision=REJECTED_DUPLICATE reason=already_committed  

session=session-001 mutation=payment-002 authority=node-a epoch=3 sequence=14 decision=REJECTED_NON_AUTHORITY reason=invalid_authority  

session=session-001 mutation=payment-003 authority=node-b epoch=2 sequence=15 decision=REJECTED_STALE_EPOCH reason=epoch_is_stale  

---

## 4. Minimum Metrics

A VRP system SHOULD expose:

vrp_accepted_total  
vrp_rejected_duplicate_total  
vrp_rejected_non_authority_total  
vrp_rejected_stale_epoch_total  
vrp_rejected_replay_total  
vrp_convergence_success_total  
vrp_convergence_violation_total  

---

## 5. Health Signals

A VRP runtime is healthy when:

- accepted mutations are deterministic  
- duplicates are rejected  
- stale inputs are rejected  
- authority is consistent  
- nodes converge to the same result  

---

## 6. Operator Questions

The system must answer:

why was this mutation accepted?  
why was this mutation rejected?  
which authority made the decision?  
which epoch was active?  
was this a duplicate or replay?  

---

## 7. Proof-Oriented Logs

VRP logs are not only diagnostics.

They are evidence.

Example:

mutation=payment-001 decision=ACCEPTED  
mutation=payment-001 decision=REJECTED_DUPLICATE  
verdict=CONSISTENT  

---

## 8. Summary

VRP observability turns runtime behavior into verifiable evidence.

Correctness must be visible.