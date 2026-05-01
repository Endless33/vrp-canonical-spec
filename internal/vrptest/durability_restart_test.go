package vrptest

import (
	"encoding/json"
	"os"
	"testing"
)

type DurableDecision string

const (
	DurableAccepted          DurableDecision = "ACCEPTED"
	DurableReturnedCanonical DurableDecision = "RETURNED_CANONICAL_RESULT"
)

type DurableRecord struct {
	MutationID string
	Result     string
	Balance    int
}

type DurableRuntime struct {
	filePath string
	records  map[string]DurableRecord
	balance  int
}

func NewDurableRuntime(filePath string) *DurableRuntime {
	r := &DurableRuntime{
		filePath: filePath,
		records:  make(map[string]DurableRecord),
	}

	r.load()

	return r
}

func (r *DurableRuntime) load() {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return
	}

	var stored map[string]DurableRecord
	if err := json.Unmarshal(data, &stored); err == nil {
		r.records = stored

		// rebuild balance deterministically
		total := 0
		for _, rec := range stored {
			total += rec.Balance
		}
		r.balance = total
	}
}

func (r *DurableRuntime) persist() {
	data, _ := json.Marshal(r.records)
	_ = os.WriteFile(r.filePath, data, 0644)
}

func (r *DurableRuntime) Transfer(mutationID string, amount int) (DurableDecision, DurableRecord) {
	if rec, ok := r.records[mutationID]; ok {
		return DurableReturnedCanonical, rec
	}

	r.balance += amount

	rec := DurableRecord{
		MutationID: mutationID,
		Result:     "committed",
		Balance:    amount,
	}

	r.records[mutationID] = rec
	r.persist()

	return DurableAccepted, rec
}

func TestDurabilityAcrossRestart(t *testing.T) {
	file := "test_store.json"
	defer os.Remove(file)

	r1 := NewDurableRuntime(file)

	decision1, rec1 := r1.Transfer("payment-001", 100)

	if decision1 != DurableAccepted {
		t.Fatalf("expected ACCEPTED, got %s", decision1)
	}

	// simulate restart
	r2 := NewDurableRuntime(file)

	decision2, rec2 := r2.Transfer("payment-001", 100)

	if decision2 != DurableReturnedCanonical {
		t.Fatalf("expected RETURNED_CANONICAL_RESULT after restart, got %s", decision2)
	}

	if rec1.Result != rec2.Result {
		t.Fatalf("expected same result after restart")
	}

	if len(r2.records) != 1 {
		t.Fatalf("expected 1 record after restart, got %d", len(r2.records))
	}
}

func TestDurabilityDoesNotDuplicateStateAfterRestart(t *testing.T) {
	file := "test_store.json"
	defer os.Remove(file)

	r1 := NewDurableRuntime(file)
	r1.Transfer("payment-001", 100)

	// restart
	r2 := NewDurableRuntime(file)
	r2.Transfer("payment-001", 100)
	r2.Transfer("payment-001", 100)

	if len(r2.records) != 1 {
		t.Fatalf("expected single record after retries, got %d", len(r2.records))
	}
}