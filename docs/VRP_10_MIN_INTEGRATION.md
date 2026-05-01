# VRP 10 Minute Integration

This document shows how to integrate VRP as an execution correctness boundary in an existing system.

VRP does not replace the application.

It wraps the point where state mutations happen.

---

## 1. Choose One Operation

Start with one state-changing endpoint.

Examples:

POST /transfer  
POST /jobs/run  
POST /orders/create  

Pick something where retries can cause damage.

---

## 2. Add Mutation Identity

Every request must have a stable mutation_id.

Example:

mutation_id=payment-001

This must remain the same across retries.

---

## 3. Build VRP Input

Before applying state:

session_id  
mutation_id  
authority  
epoch  
sequence  
payload_hash  

---

## 4. Call VRP Before Mutation

decision = vrp.Accept(input)

if decision == ACCEPTED:
    apply_state_change()
else:
    reject_or_ignore()

---

## 5. Do Not Retry Commit

If VRP rejects:

do not retry blindly  
do not replay mutation  
do not rebuild state  

---

## 6. Log the Decision

At minimum:

session_id  
mutation_id  
authority  
epoch  
sequence  
decision  

Example:

mutation=payment-001 decision=ACCEPTED  
mutation=payment-001 decision=REJECTED_DUPLICATE  

---

## 7. Expected Behavior

valid mutation → ACCEPTED  
retry → REJECTED_DUPLICATE  
wrong authority → REJECTED_NON_AUTHORITY  
old epoch → REJECTED_STALE_EPOCH  

---

## 8. Minimal Flow

request arrives  
→ build VRP input  
→ VRP Accept(input)  
→ if ACCEPTED → mutate state  
→ else → reject  

---

## 9. What You Get

- no double execution  
- safe retries  
- deterministic commits  
- clean logs  
- audit-friendly behavior  

---

## 10. Summary

VRP can be added to one endpoint in minutes.

Start small.

Protect one mutation boundary.