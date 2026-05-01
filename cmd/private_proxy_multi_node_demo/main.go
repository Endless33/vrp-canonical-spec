package main

import (
	"fmt"
)

type Decision string

const (
	Accepted          Decision = "ACCEPTED"
	RejectedDuplicate Decision = "REJECTED_DUPLICATE"
)

type Input struct {
	MutationID string
}

type Runtime struct {
	committed map[string]bool
}

func NewRuntime() *Runtime {
	return &Runtime{
		committed: make(map[string]bool),
	}
}

func (r *Runtime) Accept(input Input) Decision {
	if r.committed[input.MutationID] {
		return RejectedDuplicate
	}
	r.committed[input.MutationID] = true
	return Accepted
}

func runScenario(name string, runtime *Runtime, inputs []Input) []Decision {
	fmt.Println("===", name, "===")

	var decisions []Decision

	for _, in := range inputs {
		d := runtime.Accept(in)
		decisions = append(decisions, d)

		fmt.Printf("mutation=%s → %s\n", in.MutationID, d)
	}

	fmt.Println()
	return decisions
}

func main() {
	fmt.Println("=== VRP MULTI-NODE DEMO ===")
	fmt.Println("Invariant: independent runtimes must produce same decisions")
	fmt.Println()

	inputsA := []Input{
		{MutationID: "payment-001"},
		{MutationID: "payment-001"},
	}

	inputsB := []Input{
		{MutationID: "payment-001"},
		{MutationID: "payment-001"},
	}

	nodeA := NewRuntime()
	nodeB := NewRuntime()

	decA := runScenario("NODE A", nodeA, inputsA)
	decB := runScenario("NODE B", nodeB, inputsB)

	fmt.Println("=== COMPARISON ===")

	same := true
	for i := range decA {
		if decA[i] != decB[i] {
			same = false
		}
	}

	if same {
		fmt.Println("CONSISTENT")
		fmt.Println("Proof: both nodes made identical decisions")
	} else {
		fmt.Println("VIOLATION")
	}
}