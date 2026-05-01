# VRP Proxy

This document describes the minimal VRP proxy included in this repository.

The proxy is the first touchable entry point for testing VRP behavior through HTTP requests.

---

## Purpose

The proxy demonstrates one core rule:

delivery does not imply validity.

A request may arrive successfully and still be rejected by VRP.

---

## What It Shows

The proxy accepts HTTP requests and evaluates them through a minimal VRP commit boundary.

It shows:

- first mutation → accepted
- repeated mutation → rejected as duplicate
- state-changing operation → allowed only after VRP acceptance

---

## Run

go run ./cmd/vrp_proxy

---

## Test

curl -X POST http://127.0.0.1:8080/transfer -H "X-Mutation-ID: payment-001"  
curl -X POST http://127.0.0.1:8080/transfer -H "X-Mutation-ID: payment-001"  

---

## Expected Behavior

first request  → ACCEPTED  
second request → REJECTED_DUPLICATE  

---

## Meaning

The proxy turns VRP from a specification into a runnable boundary.

A developer can send real HTTP requests and observe deterministic commit behavior.

---

## Important Note

This proxy is intentionally minimal.

It is not a production proxy.

It exists to prove the commit boundary behavior through a simple external interface.

---

## Summary

VRP proxy demonstrates:

HTTP request → VRP decision → commit or reject

This is the smallest practical way to touch VRP behavior.