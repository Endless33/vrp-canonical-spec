package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Decision string

const (
	ACCEPTED                 Decision = "ACCEPTED"
	REJECTED_DUPLICATE       Decision = "REJECTED_DUPLICATE"
	RETURNED_CANONICAL       Decision = "RETURNED_CANONICAL_RESULT"
)

type Response struct {
	MutationID string    `json:"mutation_id"`
	Decision   Decision  `json:"decision"`
	Result     string    `json:"result"`
	Time       time.Time `json:"time"`
}

type Server struct {
	mu        sync.Mutex
	committed map[string]Response
}

func NewServer() *Server {
	return &Server{
		committed: make(map[string]Response),
	}
}

func (s *Server) handleTransfer(w http.ResponseWriter, r *http.Request) {
	mutationID := r.Header.Get("X-Mutation-ID")
	if mutationID == "" {
		http.Error(w, "missing mutation id", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if cached, ok := s.committed[mutationID]; ok {
		resp := cached
		resp.Decision = RETURNED_CANONICAL
		json.NewEncoder(w).Encode(resp)
		log.Printf("mutation=%s decision=%s (cached)", mutationID, resp.Decision)
		return
	}

	resp := Response{
		MutationID: mutationID,
		Decision:   ACCEPTED,
		Result:     "committed",
		Time:       time.Now(),
	}

	s.committed[mutationID] = resp

	json.NewEncoder(w).Encode(resp)
	log.Printf("mutation=%s decision=%s", mutationID, resp.Decision)
}

func main() {
	s := NewServer()

	http.HandleFunc("/transfer", s.handleTransfer)

	log.Println("VRP PROXY SERVER listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}