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

type Result string

const (
	Winner   Result = "WINNER"
	Rejected Result = "REJECTED"
)

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

	fmt.Println("=== VRP MULTI-NODE CONVERGENCE DEMO ===")
	fmt.Println("Invariant: independent runtimes must select the same winner\n")

	// Same input set
	decisions := []Decision{
		{Mutation: "payment-001", Epoch: 2, Authority: "node-a", Priority: 90},
		{Mutation: "payment-001", Epoch: 3, Authority: "node-b", Priority: 50},
		{Mutation: "payment-001", Epoch: 3, Authority: "node-c", Priority: 10},
	}

	nodeA := NewRuntime("node-A-runtime")
	nodeB := NewRuntime("node-B-runtime")

	winnerA := nodeA.Resolve(decisions)
	winnerB := nodeB.Resolve(decisions)

	fmt.Println("=== INPUT ===")
	for _, d := range decisions {
		fmt.Printf("mutation=%s epoch=%d authority=%s priority=%d\n",
			d.Mutation, d.Epoch, d.Authority, d.Priority)
	}

	fmt.Println()
	fmt.Println("=== NODE A RESULT ===")
	fmt.Printf("winner=%s epoch=%d authority=%s priority=%d\n",
		winnerA.Mutation, winnerA.Epoch, winnerA.Authority, winnerA.Priority)

	fmt.Println()
	fmt.Println("=== NODE B RESULT ===")
	fmt.Printf("winner=%s epoch=%d authority=%s priority=%d\n",
		winnerB.Mutation, winnerB.Epoch, winnerB.Authority, winnerB.Priority)

	fmt.Println()
	fmt.Println("=== VERDICT ===")

	if winnerA == winnerB {
		fmt.Println("CONSISTENT")
		fmt.Println("Proof: independent runtimes converged to the same decision")
	} else {
		fmt.Println("DIVERGENCE DETECTED")
	}
}