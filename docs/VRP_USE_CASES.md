# VRP Use Cases

This document describes practical scenarios where VRP semantics apply.

VRP is designed for systems where network disorder must not corrupt execution.

---

## 1. Bank Transfer

### Problem

A payment request is sent.

The response is lost.

The client retries.

The system receives the same logical mutation twice.

Without a commit boundary, this may produce double execution.

---

### VRP Behavior

payment-001 → ACCEPTED  
payment-001 retry → REJECTED_DUPLICATE  

---

### Guarantee

one logical transfer → one committed state transition

---

## 2. Distributed Job Execution

### Problem

A job scheduler sends work to a node.

The network duplicates the request.

Two workers attempt to process the same job.

Without deterministic commit admission, the job may execute twice.

---

### VRP Behavior

job-778 → ACCEPTED  
job-778 duplicate → REJECTED_DUPLICATE  

---

### Guarantee

one logical job → one canonical execution

---

## 3. API Request Under Retry

### Problem

An API request times out.

The client retries.

The server receives multiple equivalent mutation attempts.

Transport success does not prove execution correctness.

---

### VRP Behavior

request-42 → ACCEPTED  
request-42 retry → REJECTED_DUPLICATE  

---

### Guarantee

retry cannot create a second valid mutation

---

## 4. Authority Race

### Problem

Two or more authorities attempt to decide the same mutation.

Network timing makes multiple decisions appear valid.

Without deterministic resolution, state may diverge.

---

### VRP Behavior

epoch=3 → ACCEPTED  
epoch=2 → REJECTED_LOWER_EPOCH  
same epoch lower priority → REJECTED_LOWER_PRIORITY  

---

### Guarantee

multiple decision sources → one canonical decision

---

## 5. Path Migration

### Problem

A session moves between network paths.

Packets from the old path may arrive late.

Without validation, stale inputs may mutate state.

---

### VRP Behavior

stale input → REJECTED  
valid input → evaluated  

---

### Guarantee

path change does not corrupt session state

---

## 6. Summary

VRP applies wherever the system must preserve correctness under:

- retries
- packet duplication
- packet loss
- out-of-order delivery
- authority races
- path migration
- stale inputs

VRP does not make the network reliable.

VRP makes execution independent of network reliability.