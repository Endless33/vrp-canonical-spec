package main

import (
	"fmt"
)

type Mutation struct {
	ID        string
	Epoch     int
	Authority string
}

type CommitResult string

const (
	Accepted        CommitResult = "ACCEPTED"
	RejectedDup     CommitResult = "REJECTED_DUPLICATE"
	RejectedStale   CommitResult = "REJECTED_STALE_EPOCH"
	RejectedAuth    CommitResult = "REJECTED_NON_AUTHORITY"
)

type Runtime struct {
	currentEpoch     int
	currentAuthority string
	committed        map[string]bool
}

func NewRuntime(epoch int, authority string) *Runtime {
	return &Runtime{
		currentEpoch:     epoch,
		currentAuthority: authority,
		committed:        make(map[string]bool),
	}
}

func (r *Runtime) Evaluate(m Mutation) CommitResult {

	// 1. Authority check
	if m.Authority != r.currentAuthority {
		return RejectedAuth
	}

	// 2. Epoch check
	if m.Epoch < r.currentEpoch {
		return RejectedStale
	}

	// 3. Duplicate check
	if r.committed[m.ID] {
		return RejectedDup
	}

	// 4. Commit
	r.committed[m.ID] = true
	return Accepted
}

func main() {

	fmt.Println("=== VRP CANONICAL COMMIT CONTRACT DEMO ===")
	fmt.Println("Invariant: one mutation commits at most once\n")

	r := NewRuntime(1, "A")

	events := []Mutation{
		{ID: "tx1", Epoch: 1, Authority: "A"}, // valid
		{ID: "tx1", Epoch: 1, Authority: "A"}, // duplicate
		{ID: "tx2", Epoch: 1, Authority: "B"}, // wrong authority
		{ID: "tx3", Epoch: 0, Authority: "A"}, // stale epoch
		{ID: "tx4", Epoch: 1, Authority: "A"}, // valid
	}

	results := make(map[CommitResult]int)

	for i, e := range events {
		result := r.Evaluate(e)
		results[result]++

		fmt.Printf("EVENT %d → %+v\n", i+1, e)
		fmt.Printf("  → %s\n\n", result)
	}

	fmt.Println("=== SUMMARY ===")
	for k, v := range results {
		fmt.Printf("%s = %d\n", k, v)
	}

	fmt.Println("\n=== VERDICT ===")

	if results[Accepted] == 2 &&
		results[RejectedDup] == 1 &&
		results[RejectedAuth] == 1 &&
		results[RejectedStale] == 1 {

		fmt.Println("CONSISTENT")
		fmt.Println("Proof: commit contract enforced correctly")
	} else {
		fmt.Println("VIOLATION DETECTED")
	}
}