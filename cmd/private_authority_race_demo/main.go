package main

import (
	"fmt"
	"sort"
)

type Decision struct {
	Mutation  string
	Epoch     int
	Authority string
	Priority  int
	Value     string
}

type RaceResult string

const (
	AcceptedWinner       RaceResult = "ACCEPTED_WINNER"
	RejectedLowerEpoch   RaceResult = "REJECTED_LOWER_EPOCH"
	RejectedLowerPriority RaceResult = "REJECTED_LOWER_PRIORITY"
	RejectedDuplicate    RaceResult = "REJECTED_DUPLICATE_DECISION"
)

type EvaluatedDecision struct {
	Decision Decision
	Result   RaceResult
	Reason   string
}

type Runtime struct {
	committed map[string]bool
}

func NewRuntime() *Runtime {
	return &Runtime{
		committed: make(map[string]bool),
	}
}

func (r *Runtime) Resolve(decisions []Decision) []EvaluatedDecision {
	if len(decisions) == 0 {
		return nil
	}

	sorted := make([]Decision, len(decisions))
	copy(sorted, decisions)

	sort.SliceStable(sorted, func(i, j int) bool {
		if sorted[i].Epoch != sorted[j].Epoch {
			return sorted[i].Epoch > sorted[j].Epoch
		}

		if sorted[i].Priority != sorted[j].Priority {
			return sorted[i].Priority > sorted[j].Priority
		}

		return sorted[i].Authority < sorted[j].Authority
	})

	winner := sorted[0]
	results := make([]EvaluatedDecision, 0, len(sorted))

	for _, d := range sorted {
		if d.Mutation != winner.Mutation {
			results = append(results, EvaluatedDecision{
				Decision: d,
				Result:   RejectedDuplicate,
				Reason:   "decision belongs to a different race group",
			})
			continue
		}

		if r.committed[d.Mutation] {
			results = append(results, EvaluatedDecision{
				Decision: d,
				Result:   RejectedDuplicate,
				Reason:   "mutation already has a canonical decision",
			})
			continue
		}

		if d.Epoch < winner.Epoch {
			results = append(results, EvaluatedDecision{
				Decision: d,
				Result:   RejectedLowerEpoch,
				Reason:   "lower epoch cannot override higher epoch",
			})
			continue
		}

		if d.Epoch == winner.Epoch && d.Priority < winner.Priority {
			results = append(results, EvaluatedDecision{
				Decision: d,
				Result:   RejectedLowerPriority,
				Reason:   "same epoch conflict resolved by deterministic priority",
			})
			continue
		}

		r.committed[d.Mutation] = true

		results = append(results, EvaluatedDecision{
			Decision: d,
			Result:   AcceptedWinner,
			Reason:   "canonical authority decision selected",
		})
	}

	return results
}

func main() {
	fmt.Println("=== VRP AUTHORITY RACE DEMO ===")
	fmt.Println("Invariant: competing authorities must produce one canonical decision")
	fmt.Println()

	runtime := NewRuntime()

	decisions := []Decision{
		{
			Mutation:  "payment-001",
			Epoch:     2,
			Authority: "node-a",
			Priority:  90,
			Value:     "commit:100",
		},
		{
			Mutation:  "payment-001",
			Epoch:     3,
			Authority: "node-b",
			Priority:  50,
			Value:     "commit:100",
		},
		{
			Mutation:  "payment-001",
			Epoch:     3,
			Authority: "node-c",
			Priority:  10,
			Value:     "commit:999",
		},
	}

	results := runtime.Resolve(decisions)

	summary := make(map[RaceResult]int)

	fmt.Println("=== RACE INPUT ===")
	for _, d := range decisions {
		fmt.Printf(
			"mutation=%s epoch=%d authority=%s priority=%d value=%s\n",
			d.Mutation,
			d.Epoch,
			d.Authority,
			d.Priority,
			d.Value,
		)
	}

	fmt.Println()
	fmt.Println("=== RACE RESULT ===")
	for _, result := range results {
		summary[result.Result]++

		fmt.Printf(
			"mutation=%s epoch=%d authority=%s priority=%d → %s (%s)\n",
			result.Decision.Mutation,
			result.Decision.Epoch,
			result.Decision.Authority,
			result.Decision.Priority,
			result.Result,
			result.Reason,
		)
	}

	fmt.Println()
	fmt.Println("=== SUMMARY ===")
	for k, v := range summary {
		fmt.Printf("%s = %d\n", k, v)
	}

	fmt.Println()
	fmt.Println("=== VERDICT ===")

	if summary[AcceptedWinner] == 1 &&
		summary[RejectedLowerEpoch] == 1 &&
		summary[RejectedLowerPriority] == 1 {

		fmt.Println("CONSISTENT")
		fmt.Println("Proof: authority race produced one canonical winner")
	} else {
		fmt.Println("VIOLATION DETECTED")
	}
}