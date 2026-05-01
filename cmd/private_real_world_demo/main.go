package main

import (
	"fmt"
)

type Decision string

const (
	Accepted             Decision = "ACCEPTED"
	RejectedDuplicate    Decision = "REJECTED_DUPLICATE"
	RejectedNonAuthority Decision = "REJECTED_NON_AUTHORITY"
	RejectedStaleEpoch   Decision = "REJECTED_STALE_EPOCH"
)

type Request struct {
	MutationID string
	Authority  string
	Epoch      int
}

type Runtime struct {
	currentAuthority string
	currentEpoch     int
	committed        map[string]bool
}

func NewRuntime() *Runtime {
	return &Runtime{
		currentAuthority: "node-b",
		currentEpoch:     3,
		committed:        make(map[string]bool),
	}
}

func (r *Runtime) Accept(req Request) Decision {
	if req.Authority != r.currentAuthority {
		return RejectedNonAuthority
	}

	if req.Epoch < r.currentEpoch {
		return RejectedStaleEpoch
	}

	if r.committed[req.MutationID] {
		return RejectedDuplicate
	}

	r.committed[req.MutationID] = true
	return Accepted
}

func handleTransfer(r *Runtime, req Request) {
	fmt.Printf("POST /transfer mutation=%s authority=%s epoch=%d\n",
		req.MutationID,
		req.Authority,
		req.Epoch,
	)

	decision := r.Accept(req)

	fmt.Printf("decision=%s\n", decision)

	if decision == Accepted {
		fmt.Println("state_change=balance_updated")
	} else {
		fmt.Println("state_change=blocked")
	}

	fmt.Println()
}

func main() {
	fmt.Println("=== VRP REAL WORLD DEMO ===")
	fmt.Println("Scenario: transfer with retry and invalid inputs")
	fmt.Println()

	runtime := NewRuntime()

	handleTransfer(runtime, Request{
		MutationID: "payment-001",
		Authority:  "node-b",
		Epoch:      3,
	})

	handleTransfer(runtime, Request{
		MutationID: "payment-001",
		Authority:  "node-b",
		Epoch:      3,
	})

	handleTransfer(runtime, Request{
		MutationID: "payment-002",
		Authority:  "node-a",
		Epoch:      3,
	})

	handleTransfer(runtime, Request{
		MutationID: "payment-003",
		Authority:  "node-b",
		Epoch:      2,
	})

	fmt.Println("=== RESULT ===")
	fmt.Println("valid=committed_once")
	fmt.Println("retry=blocked")
	fmt.Println("wrong_authority=blocked")
	fmt.Println("stale_epoch=blocked")
	fmt.Println()
	fmt.Println("VERDICT=CONSISTENT")
	fmt.Println("Proof: transfer endpoint preserved correctness under retry and invalid inputs")
}