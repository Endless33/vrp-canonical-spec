package vrptest

import "testing"

type LostResponseDecision string

const (
	LostResponseAccepted          LostResponseDecision = "ACCEPTED"
	LostResponseReturnedCanonical LostResponseDecision = "RETURNED_CANONICAL_RESULT"
)

type LostResponseResult struct {
	MutationID string
	Decision   LostResponseDecision
	Result     string
	Balance    int
}

type LostResponseRuntime struct {
	committed map[string]LostResponseResult
	balance   int
}

func NewLostResponseRuntime() *LostResponseRuntime {
	return &LostResponseRuntime{
		committed: make(map[string]LostResponseResult),
	}
}

func (r *LostResponseRuntime) Transfer(mutationID string, amount int) LostResponseResult {
	if cached, ok := r.committed[mutationID]; ok {
		return LostResponseResult{
			MutationID: cached.MutationID,
			Decision:   LostResponseReturnedCanonical,
			Result:     cached.Result,
			Balance:    cached.Balance,
		}
	}

	r.balance += amount

	result := LostResponseResult{
		MutationID: mutationID,
		Decision:   LostResponseAccepted,
		Result:     "transfer_committed",
		Balance:    r.balance,
	}

	r.committed[mutationID] = result

	return result
}

func TestLostAcceptedResponseRetryReturnsCanonicalResult(t *testing.T) {
	runtime := NewLostResponseRuntime()

	first := runtime.Transfer("payment-001", 100)

	if first.Decision != LostResponseAccepted {
		t.Fatalf("expected first decision ACCEPTED, got %s", first.Decision)
	}

	if first.Balance != 100 {
		t.Fatalf("expected first balance 100, got %d", first.Balance)
	}

	// Simulate lost ACCEPTED response:
	// client does not receive first result and retries the same mutation.
	retry := runtime.Transfer("payment-001", 100)

	if retry.Decision != LostResponseReturnedCanonical {
		t.Fatalf("expected retry to return canonical result, got %s", retry.Decision)
	}

	if retry.Result != first.Result {
		t.Fatalf("expected same result, got first=%s retry=%s", first.Result, retry.Result)
	}

	if retry.Balance != first.Balance {
		t.Fatalf("expected same balance, got first=%d retry=%d", first.Balance, retry.Balance)
	}

	if runtime.balance != 100 {
		t.Fatalf("expected balance to mutate once, got %d", runtime.balance)
	}
}

func TestRepeatedRetriesDoNotMutateStateAgain(t *testing.T) {
	runtime := NewLostResponseRuntime()

	runtime.Transfer("payment-001", 100)
	runtime.Transfer("payment-001", 100)
	runtime.Transfer("payment-001", 100)
	runtime.Transfer("payment-001", 100)

	if runtime.balance != 100 {
		t.Fatalf("expected one mutation, got balance %d", runtime.balance)
	}

	if len(runtime.committed) != 1 {
		t.Fatalf("expected one committed mutation, got %d", len(runtime.committed))
	}
}