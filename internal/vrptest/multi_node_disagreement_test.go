package vrptest

import "testing"

type MultiNodeDecision string

const (
	MultiNodeAccepted          MultiNodeDecision = "ACCEPTED"
	MultiNodeRejectedDuplicate MultiNodeDecision = "REJECTED_DUPLICATE"
)

type MultiNodeRuntime struct {
	committed map[string]bool
}

func NewMultiNodeRuntime() *MultiNodeRuntime {
	return &MultiNodeRuntime{
		committed: make(map[string]bool),
	}
}

func (r *MultiNodeRuntime) Accept(mutationID string) MultiNodeDecision {
	if r.committed[mutationID] {
		return MultiNodeRejectedDuplicate
	}

	r.committed[mutationID] = true
	return MultiNodeAccepted
}

func TestMultiNodeSameMutationSameOutcome(t *testing.T) {
	nodeA := NewMultiNodeRuntime()
	nodeB := NewMultiNodeRuntime()

	inputsA := []string{
		"payment-001",
		"payment-001",
	}

	inputsB := []string{
		"payment-001",
		"payment-001",
	}

	var decisionsA []MultiNodeDecision
	var decisionsB []MultiNodeDecision

	for _, mutationID := range inputsA {
		decisionsA = append(decisionsA, nodeA.Accept(mutationID))
	}

	for _, mutationID := range inputsB {
		decisionsB = append(decisionsB, nodeB.Accept(mutationID))
	}

	if len(decisionsA) != len(decisionsB) {
		t.Fatalf("decision length mismatch: nodeA=%d nodeB=%d", len(decisionsA), len(decisionsB))
	}

	for i := range decisionsA {
		if decisionsA[i] != decisionsB[i] {
			t.Fatalf("decision mismatch at index %d: nodeA=%s nodeB=%s", i, decisionsA[i], decisionsB[i])
		}
	}
}

func TestMultiNodeDisagreementIsDetected(t *testing.T) {
	nodeA := NewMultiNodeRuntime()
	nodeB := NewMultiNodeRuntime()

	nodeA.Accept("payment-001")

	if nodeA.committed["payment-001"] == nodeB.committed["payment-001"] {
		t.Fatalf("expected disagreement before synchronization")
	}

	nodeB.Accept("payment-001")

	if nodeA.committed["payment-001"] != nodeB.committed["payment-001"] {
		t.Fatalf("expected both nodes to converge after same mutation is observed")
	}
}