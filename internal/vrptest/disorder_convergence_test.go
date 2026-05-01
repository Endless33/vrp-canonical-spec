package vrptest

import "testing"

type DisorderDecision string

const (
	DisorderAccepted          DisorderDecision = "ACCEPTED"
	DisorderRejectedDuplicate DisorderDecision = "REJECTED_DUPLICATE"
)

type DisorderInput struct {
	MutationID string
}

type DisorderEngine struct {
	committed map[string]bool
}

func NewDisorderEngine() *DisorderEngine {
	return &DisorderEngine{
		committed: make(map[string]bool),
	}
}

func (e *DisorderEngine) Accept(input DisorderInput) DisorderDecision {
	if e.committed[input.MutationID] {
		return DisorderRejectedDuplicate
	}

	e.committed[input.MutationID] = true
	return DisorderAccepted
}

func (e *DisorderEngine) FinalState() map[string]bool {
	out := make(map[string]bool)

	for mutationID, committed := range e.committed {
		out[mutationID] = committed
	}

	return out
}

func TestDisorderConvergence(t *testing.T) {
	nodeA := NewDisorderEngine()
	nodeB := NewDisorderEngine()

	inputsA := []DisorderInput{
		{MutationID: "payment-001"},
		{MutationID: "payment-002"},
		{MutationID: "payment-001"},
		{MutationID: "payment-003"},
		{MutationID: "payment-002"},
	}

	inputsB := []DisorderInput{
		{MutationID: "payment-002"},
		{MutationID: "payment-001"},
		{MutationID: "payment-003"},
		{MutationID: "payment-001"},
		{MutationID: "payment-002"},
	}

	for _, input := range inputsA {
		nodeA.Accept(input)
	}

	for _, input := range inputsB {
		nodeB.Accept(input)
	}

	stateA := nodeA.FinalState()
	stateB := nodeB.FinalState()

	if len(stateA) != len(stateB) {
		t.Fatalf("expected same committed set size, got nodeA=%d nodeB=%d", len(stateA), len(stateB))
	}

	for mutationID := range stateA {
		if !stateB[mutationID] {
			t.Fatalf("mutation %s committed on nodeA but not nodeB", mutationID)
		}
	}

	for mutationID := range stateB {
		if !stateA[mutationID] {
			t.Fatalf("mutation %s committed on nodeB but not nodeA", mutationID)
		}
	}
}