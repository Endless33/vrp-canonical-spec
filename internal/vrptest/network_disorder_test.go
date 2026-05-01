package vrptest

import "testing"

type NetworkPacket struct {
	ID         string
	MutationID string
	Drop       bool
	Duplicate  bool
}

type NetworkRuntime struct {
	committed map[string]bool
}

func NewNetworkRuntime() *NetworkRuntime {
	return &NetworkRuntime{
		committed: make(map[string]bool),
	}
}

func (r *NetworkRuntime) Accept(mutationID string) {
	if mutationID == "" {
		return
	}

	if r.committed[mutationID] {
		return
	}

	r.committed[mutationID] = true
}

func ApplyNetworkDisorder(runtime *NetworkRuntime, packets []NetworkPacket) {
	for _, packet := range packets {
		if packet.Drop {
			continue
		}

		runtime.Accept(packet.MutationID)

		if packet.Duplicate {
			runtime.Accept(packet.MutationID)
		}
	}
}

func TestNetworkDisorderDoesNotCorruptCommittedState(t *testing.T) {
	runtime := NewNetworkRuntime()

	packets := []NetworkPacket{
		{ID: "pkt-1", MutationID: "payment-001"},
		{ID: "pkt-2", MutationID: "payment-002", Drop: true},
		{ID: "pkt-3", MutationID: "payment-001", Duplicate: true},
		{ID: "pkt-4", MutationID: "payment-003"},
		{ID: "pkt-5", MutationID: "payment-003", Duplicate: true},
	}

	ApplyNetworkDisorder(runtime, packets)

	if len(runtime.committed) != 2 {
		t.Fatalf("expected 2 committed mutations, got %d", len(runtime.committed))
	}

	if !runtime.committed["payment-001"] {
		t.Fatalf("expected payment-001 to be committed")
	}

	if runtime.committed["payment-002"] {
		t.Fatalf("expected dropped payment-002 not to be committed")
	}

	if !runtime.committed["payment-003"] {
		t.Fatalf("expected payment-003 to be committed")
	}
}

func TestNetworkDisorderReorderedInputsConverge(t *testing.T) {
	nodeA := NewNetworkRuntime()
	nodeB := NewNetworkRuntime()

	packetsA := []NetworkPacket{
		{ID: "pkt-1", MutationID: "payment-001"},
		{ID: "pkt-2", MutationID: "payment-002"},
		{ID: "pkt-3", MutationID: "payment-001", Duplicate: true},
		{ID: "pkt-4", MutationID: "payment-003"},
	}

	packetsB := []NetworkPacket{
		{ID: "pkt-4", MutationID: "payment-003"},
		{ID: "pkt-3", MutationID: "payment-001", Duplicate: true},
		{ID: "pkt-2", MutationID: "payment-002"},
		{ID: "pkt-1", MutationID: "payment-001"},
	}

	ApplyNetworkDisorder(nodeA, packetsA)
	ApplyNetworkDisorder(nodeB, packetsB)

	if len(nodeA.committed) != len(nodeB.committed) {
		t.Fatalf("expected same committed set size, got nodeA=%d nodeB=%d", len(nodeA.committed), len(nodeB.committed))
	}

	for mutationID := range nodeA.committed {
		if !nodeB.committed[mutationID] {
			t.Fatalf("mutation %s committed on nodeA but not nodeB", mutationID)
		}
	}

	for mutationID := range nodeB.committed {
		if !nodeA.committed[mutationID] {
			t.Fatalf("mutation %s committed on nodeB but not nodeA", mutationID)
		}
	}
}