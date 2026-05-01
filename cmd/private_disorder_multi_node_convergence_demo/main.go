package main

import (
	"fmt"
	"sort"
)

type Decision struct {
	Mutation  string
	Epoch     int
	Authority string
	Priority  int
}

type Runtime struct {
	name string
}

func NewRuntime(name string) *Runtime {
	return &Runtime{name: name}
}

func (r *Runtime) Resolve(decisions []Decision) Decision {

	sorted := make([]Decision, len(decisions))
	copy(sorted, decisions)

	sort.SliceStable(sorted, func(i, j int) bool {
		if sorted[i].Epoch != sorted[j].Epoch {
			return sorted[i].Epoch > sorted[j].Epoch
		}
		if sorted[i].Priority != sorted[j].Priority {
			return sorted[i].Priority > sorted[j].Priority
		}
		return sorted[i].Authority < sorted[j].Authority
	})

	return sorted[0]
}

func main() {

	fmt.Println("=== VRP DISORDER + MULTI-NODE CONVERGENCE ===")
	fmt.Println("Invariant: different input order must not change decision\n")

	// Same logical decisions, different delivery order
	nodeAInput := []Decision{
		{Mutation: "payment-001", Epoch: 3, Authority: "node-c", Priority: 10},
		{Mutation: "payment-001", Epoch: 2, Authority: "node-a", Priority: 90},
		{Mutation: "payment-001", Epoch: 3, Authority: "node-b", Priority: 50},
	}

	nodeBInput := []Decision{
		{Mutation: "payment-001", Epoch: 2, Authority: "node-a", Priority: 90},
		{Mutation: "payment-001", Epoch: 3, Authority: "node-b", Priority: 50},
		{Mutation: "payment-001", Epoch: 3, Authority: "node-c", Priority: 10},
	}

	nodeA := NewRuntime("node-A-runtime")
	nodeB := NewRuntime("node-B-runtime")

	winnerA := nodeA.Resolve(nodeAInput)
	winnerB := nodeB.Resolve(nodeBInput)

	fmt.Println("=== NODE A INPUT (reordered) ===")
	for _, d := range nodeAInput {
		fmt.Printf("epoch=%d authority=%s priority=%d\n",
			d.Epoch, d.Authority, d.Priority)
	}

	fmt.Println()
	fmt.Println("=== NODE B INPUT (different order) ===")
	for _, d := range nodeBInput {
		fmt.Printf("epoch=%d authority=%s priority=%d\n",
			d.Epoch, d.Authority, d.Priority)
	}

	fmt.Println()
	fmt.Println("=== RESULTS ===")

	fmt.Printf("node A winner → epoch=%d authority=%s\n",
		winnerA.Epoch, winnerA.Authority)

	fmt.Printf("node B winner → epoch=%d authority=%s\n",
		winnerB.Epoch, winnerB.Authority)

	fmt.Println()
	fmt.Println("=== VERDICT ===")

	if winnerA == winnerB {
		fmt.Println("CONSISTENT")
		fmt.Println("Proof: disorder did not affect convergence")
	} else {
		fmt.Println("DIVERGENCE DETECTED")
	}
}