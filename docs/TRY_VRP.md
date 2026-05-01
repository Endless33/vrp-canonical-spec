# Try VRP in 5 Minutes

This is the fastest way to understand VRP.

You do not need to rebuild your system.

You wrap one mutation boundary.

---

## 1. Pick One Operation

Choose any state-changing action.

Examples:

POST /transfer  
POST /jobs/run  

---

## 2. Add Mutation ID

Every request must have a stable id.

Example:

mutation_id=payment-001

Retries MUST reuse the same id.

---

## 3. Call VRP Before Mutation

decision = vrp.Accept(input)

if decision == ACCEPTED:
    apply_state_change()
else:
    reject_or_ignore()

---

## 4. Observe Behavior

Send the same request twice:

payment-001 → ACCEPTED  
payment-001 retry → REJECTED_DUPLICATE  

---

## 5. Change Authority

wrong authority → REJECTED_NON_AUTHORITY  

---

## 6. Change Epoch

old epoch → REJECTED_STALE_EPOCH  

---

## 7. What You Just Saw

- retry does not re-execute
- duplicates do not corrupt state
- invalid authority cannot mutate state
- delivery does not define correctness

---

## 8. Why This Matters

Without VRP:

retry → possible double execution

With VRP:

retry → evaluated, but cannot commit twice

---

## 9. Summary

You did not change your system.

You added one boundary:

mutation → VRP → commit

That is the VRP model.