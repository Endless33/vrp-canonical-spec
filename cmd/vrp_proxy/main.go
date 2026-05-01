package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Decision string

const (
	Accepted          Decision = "ACCEPTED"
	RejectedDuplicate Decision = "REJECTED_DUPLICATE"
	RejectedMissingID Decision = "REJECTED_MISSING_MUTATION_ID"
)

type Runtime struct {
	mu        sync.Mutex
	committed map[string]bool
}

type Verdict struct {
	MutationID string   `json:"mutation_id"`
	Decision   Decision `json:"decision"`
	Reason     string   `json:"reason"`
	Time       string   `json:"time"`
}

func NewRuntime() *Runtime {
	return &Runtime{
		committed: make(map[string]bool),
	}
}

func (r *Runtime) Accept(mutationID string) Verdict {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now().UTC().Format(time.RFC3339)

	if mutationID == "" {
		return Verdict{
			MutationID: mutationID,
			Decision:   RejectedMissingID,
			Reason:     "missing X-Mutation-ID header",
			Time:       now,
		}
	}

	if r.committed[mutationID] {
		return Verdict{
			MutationID: mutationID,
			Decision:   RejectedDuplicate,
			Reason:     "mutation already committed",
			Time:       now,
		}
	}

	r.committed[mutationID] = true

	return Verdict{
		MutationID: mutationID,
		Decision:   Accepted,
		Reason:     "canonical mutation accepted",
		Time:       now,
	}
}

func main() {
	runtime := NewRuntime()

	http.HandleFunc("/transfer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		mutationID := r.Header.Get("X-Mutation-ID")
		verdict := runtime.Accept(mutationID)

		log.Printf(
			"mutation=%s decision=%s reason=%s",
			verdict.MutationID,
			verdict.Decision,
			verdict.Reason,
		)

		w.Header().Set("Content-Type", "application/json")

		switch verdict.Decision {
		case Accepted:
			w.WriteHeader(http.StatusOK)
		case RejectedDuplicate:
			w.WriteHeader(http.StatusConflict)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}

		_ = json.NewEncoder(w).Encode(verdict)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "vrp_proxy=ok")
	})

	fmt.Println("=== VRP PROXY ===")
	fmt.Println("Listening on http://127.0.0.1:8080")
	fmt.Println()
	fmt.Println("Test:")
	fmt.Println(`curl -X POST http://127.0.0.1:8080/transfer -H "X-Mutation-ID: payment-001"`)
	fmt.Println(`curl -X POST http://127.0.0.1:8080/transfer -H "X-Mutation-ID: payment-001"`)
	fmt.Println()

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}