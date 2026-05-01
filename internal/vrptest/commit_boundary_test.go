package vrptest

import (
	"sync"
	"testing"
)

// Minimal interface we expect from your logic
type Decision string

const (
	ACCEPTED            Decision = "ACCEPTED"
	REJECTED_DUPLICATE  Decision = "REJECTED_DUPLICATE"
)

// Replace this with your real function later
// For now this is a simple thread-safe mock to prove the pattern
type Engine struct {
	mu        sync.Mutex
	committed map[string]bool
}

func NewEngine() *Engine {
	return &Engine{
		committed: make(map[string]bool),
	}
}

func (e *Engine) Accept(mutationID string) Decision {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.committed[mutationID] {
		return REJECTED_DUPLICATE
	}

	e.committed[mutationID] = true
	return ACCEPTED
}

func TestConcurrentCommitBoundary(t *testing.T) {
	engine := NewEngine()

	const workers = 1000
	const mutationID = "payment-001"

	var wg sync.WaitGroup
	wg.Add(workers)

	results := make(chan Decision, workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			results <- engine.Accept(mutationID)
		}()
	}

	wg.Wait()
	close(results)

	var accepted int
	var rejected int

	for r := range results {
		if r == ACCEPTED {
			accepted++
		}
		if r == REJECTED_DUPLICATE {
			rejected++
		}
	}

	if accepted != 1 {
		t.Fatalf("expected 1 ACCEPTED, got %d", accepted)
	}

	if rejected != workers-1 {
		t.Fatalf("expected %d REJECTED_DUPLICATE, got %d", workers-1, rejected)
	}
}