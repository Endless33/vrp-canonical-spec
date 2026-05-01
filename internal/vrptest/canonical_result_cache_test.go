package vrptest

import "testing"

type CachedDecision string

const (
	CachedAccepted          CachedDecision = "ACCEPTED"
	CachedReturnedCanonical CachedDecision = "RETURNED_CANONICAL_RESULT"
)

type CachedResult struct {
	MutationID string
	Decision   CachedDecision
	Result     string
	Balance    int
}

type CachedRuntime struct {
	committed map[string]CachedResult
	balance   int
}

func NewCachedRuntime() *CachedRuntime {
	return &CachedRuntime{
		committed: make(map[string]CachedResult),
		balance:   0,
	}
}

func (r *CachedRuntime) Transfer(mutationID string, amount int) CachedResult {
	if cached, ok := r.committed[mutationID]; ok {
		return CachedResult{
			MutationID: cached.MutationID,
			Decision:   CachedReturnedCanonical,
			Result:     cached.Result,
			Balance:    cached.Balance,
		}
	}

	r.balance += amount

	result := CachedResult{
		MutationID: mutationID,
		Decision:   CachedAccepted,
		Result:     "transfer_committed",
		Balance:    r.balance,
	}

	r.committed[mutationID] = result

	return result
}

func TestCanonicalResultReturnedOnRetry(t *testing.T) {
	runtime := NewCachedRuntime()

	first := runtime.Transfer("payment-001", 100)

	if first.Decision != CachedAccepted {
		t.Fatalf("expected first decision ACCEPTED, got %s", first.Decision)
	}

	if first.Balance != 100 {
		t.Fatalf("expected balance 100 after first commit, got %d", first.Balance)
	}

	// Simulates lost ACCEPTED response.
	// Client retries the same logical mutation.
	retry := runtime.Transfer("payment-001", 100)

	if retry.Decision != CachedReturnedCanonical {
		t.Fatalf("expected retry to return canonical result, got %s", retry.Decision)
	}

	if retry.Result != first.Result {
		t.Fatalf("expected same canonical result, got first=%s retry=%s", first.Result, retry.Result)
	}

	if retry.Balance != first.Balance {
		t.Fatalf("expected same balance, got first=%d retry=%d", first.Balance, retry.Balance)
	}

	if runtime.balance != 100 {
		t.Fatalf("expected state to mutate once, got balance %d", runtime.balance)
	}
}

func TestCanonicalResultDoesNotMutateStateTwice(t *testing.T) {
	runtime := NewCachedRuntime()

	runtime.Transfer("payment-001", 100)
	runtime.Transfer("payment-001", 100)
	runtime.Transfer("payment-001", 100)

	if runtime.balance != 100 {
		t.Fatalf("expected one state mutation, got balance %d", runtime.balance)
	}

	if len(runtime.committed) != 1 {
		t.Fatalf("expected one committed mutation, got %d", len(runtime.committed))
	}
}