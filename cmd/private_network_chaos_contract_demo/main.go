package main

import (
	"fmt"
)

type Packet struct {
	ID        string
	Mutation  string
	Duplicate bool
	Dropped   bool
}

type Runtime struct {
	committed map[string]bool
}

func NewRuntime() *Runtime {
	return &Runtime{
		committed: make(map[string]bool),
	}
}

func (r *Runtime) Process(p Packet) string {

	if p.Dropped {
		return "IGNORED_DROPPED"
	}

	if r.committed[p.Mutation] {
		return "REJECTED_DUPLICATE"
	}

	r.committed[p.Mutation] = true
	return "ACCEPTED"
}

func main() {

	fmt.Println("=== VRP NETWORK CHAOS DEMO ===")
	fmt.Println("Invariant: duplicate or dropped packets must not corrupt state\n")

	r := NewRuntime()

	packets := []Packet{
		{ID: "pkt-1", Mutation: "payment-001"},
		{ID: "pkt-1-dup", Mutation: "payment-001", Duplicate: true},
		{ID: "pkt-2", Mutation: "payment-002", Dropped: true},
		{ID: "pkt-3", Mutation: "payment-003"},
	}

	results := make(map[string]int)

	for _, p := range packets {
		result := r.Process(p)
		results[result]++

		fmt.Printf("packet=%s mutation=%s → %s\n", p.ID, p.Mutation, result)
	}

	fmt.Println("\n=== SUMMARY ===")
	for k, v := range results {
		fmt.Printf("%s = %d\n", k, v)
	}

	fmt.Println("\n=== VERDICT ===")

	if results["ACCEPTED"] == 2 &&
		results["REJECTED_DUPLICATE"] == 1 &&
		results["IGNORED_DROPPED"] == 1 {

		fmt.Println("CONSISTENT")
		fmt.Println("Proof: network disorder did not corrupt state")
	} else {
		fmt.Println("VIOLATION DETECTED")
	}
}